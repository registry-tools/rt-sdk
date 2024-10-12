package sdk

import (
	"crypto/tls"
	"fmt"
	nethttp "net/http"
	u "net/url"
	"time"

	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/registry-tools/rt-sdk/generated"
)

func NewInsecureSDKForTesting(host string) (SDK, error) {
	validator, err := auth.NewAllowedHostsValidatorErrorCheck([]string{
		host,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid host configuration: %w", err)
	}

	exampleToken := "insecure-testing-only"
	tokenProvider := &accessTokenProvider{
		allowedHosts:          validator,
		accessToken:           &exampleToken,
		accessTokenExpiration: time.Now().Add(2 * time.Hour),
	}

	authProvider := auth.NewBaseBearerTokenAuthenticationProvider(tokenProvider)

	adapter, err := NewRequestAdapter(authProvider, &nethttp.Client{
		Transport: &nethttp.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error creating request adapter: %w", err)
	}

	url, err := u.Parse(fmt.Sprintf("https://%s", host))
	if err != nil {
		return nil, fmt.Errorf("invalid host string: %w", err)
	}

	adapter.SetBaseUrl(url.String())

	client := generated.NewApiClient(adapter)

	return &sdk{
		client:   client,
		http:     adapter.Client,
		endpoint: url,
	}, nil
}
