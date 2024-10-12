package sdk

import (
	"context"
	"fmt"
	u "net/url"
	"sync"
	"time"

	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/registry-tools/rt-sdk/generated"
	internalauth "github.com/registry-tools/rt-sdk/generated/auth"
)

type accessTokenProvider struct {
	allowedHosts          *auth.AllowedHostsValidator
	clientID              string
	clientSecret          string
	host                  string
	accessToken           *string
	accessTokenExpiration time.Time
	mu                    sync.Mutex
}

var _ auth.AccessTokenProvider = &accessTokenProvider{}

func (c *accessTokenProvider) GetAllowedHostsValidator() *auth.AllowedHostsValidator {
	return c.allowedHosts
}

func (c *accessTokenProvider) getClientForTokenRefresh() (*generated.ApiClient, error) {
	provider, err := auth.NewApiKeyAuthenticationProvider("see query", "X-RT-SDK-REAUTH", auth.HEADER_KEYLOCATION)
	if err != nil {
		return nil, fmt.Errorf("error creating no-op auth provider for token refresh: %w", err)
	}
	adapter, err := NewRequestAdapter(provider, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request adapter for token refresh: %w", err)
	}
	adapter.SetBaseUrl(fmt.Sprintf("https://%s", c.host))
	return generated.NewApiClient(adapter), nil
}

func (c *accessTokenProvider) GetAuthorizationToken(context context.Context, url *u.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if c.accessToken != nil && (time.Now().Before(c.accessTokenExpiration)) {
		return *c.accessToken, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.accessToken != nil && (time.Now().Before(c.accessTokenExpiration)) {
		return *c.accessToken, nil
	}

	tokenClient, err := c.getClientForTokenRefresh()
	if err != nil {
		return "", err
	}

	authorizationCode := "authorization_code"
	tokenRequestBody := internalauth.NewTokenPostRequestBody()
	tokenRequestBody.SetClientId(&c.clientID)
	tokenRequestBody.SetClientSecret(&c.clientSecret)
	tokenRequestBody.SetGrantType(&authorizationCode)

	response, err := tokenClient.Auth().Token().PostAsTokenPostResponse(context, tokenRequestBody, nil)
	if err != nil {
		return "", fmt.Errorf("failed refreshing token: %w", err)
	}

	c.accessToken = response.GetAccessToken()
	c.accessTokenExpiration = time.Now().Add(time.Duration(*response.GetExpiresIn()) * time.Second)

	return *c.accessToken, nil
}
