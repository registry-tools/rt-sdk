package sdk

import (
	"context"
	"crypto/md5"
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	nethttp "net/http"
	u "net/url"
	"strings"
	"time"

	"github.com/registry-tools/rt-sdk/generated"
	"github.com/registry-tools/rt-sdk/generated/api"
	"github.com/registry-tools/rt-sdk/generated/wellknown"

	svchost "github.com/hashicorp/terraform-svchost"
	svcdisco "github.com/hashicorp/terraform-svchost/disco"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
)

type SDK interface {
	Api() *api.ApiRequestBuilder
	WellKnown() *wellknown.WellKnownRequestBuilder
	Client() *nethttp.Client
	Endpoint() *u.URL
	UploadFileArchive(ctx context.Context, key string, archive io.ReadSeeker) (*string, error)
}

type sdk struct {
	client   *generated.ApiClient
	http     *nethttp.Client
	endpoint *u.URL
}

func (s *sdk) Api() *api.ApiRequestBuilder {
	return s.client.Api()
}

func (s *sdk) WellKnown() *wellknown.WellKnownRequestBuilder {
	return s.client.WellKnown()
}

func (s *sdk) Client() *nethttp.Client {
	return s.http
}

func (s *sdk) Endpoint() *u.URL {
	return s.endpoint
}

// UploadFileArchive uploads a gzip file archive to the service and returns a signed ID
// that can be used in conjunction with creating other types of resources, such
// as terraform module versions.
func (s *sdk) UploadFileArchive(ctx context.Context, key string, archive io.ReadSeeker) (*string, error) {
	h := md5.New()
	size, err := io.Copy(h, archive)
	if err != nil {
		return nil, fmt.Errorf("failed to checksum archive: %w", err)
	}

	checksum := b64.StdEncoding.EncodeToString(h.Sum(nil))
	contentType := "application/gzip"

	body := api.NewArchivesPostRequestBody()
	body.SetName(&key)
	body.SetChecksumMd5Base64(&checksum)
	body.SetSizeBytes(&size)
	body.SetContentType(&contentType)

	response, err := s.Api().Archives().PostAsArchivesPostResponse(ctx, body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare file upload: %w", err)
	}

	uploadURL := response.GetMeta().GetUploadUrl()
	_, err = archive.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failed to rewind archive: %w", err)
	}

	putRequest, err := http.NewRequest(http.MethodPut, *uploadURL, archive)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize upload request: %w", err)
	}

	for key, value := range response.GetMeta().GetHeaders().GetAdditionalData() {
		val := value.(*string)
		if val != nil {
			putRequest.Header.Set(key, *val)
		}
	}
	putRequest.ContentLength = size
	uploadResponse, err := s.Client().Do(putRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to upload module: %w", err)
	}

	if uploadResponse.StatusCode >= 300 {
		body, _ := io.ReadAll(uploadResponse.Body)
		return nil, fmt.Errorf("unexpected status when uploading module: %s, %s", uploadResponse.Status, body)
	}

	return response.GetData().GetSignedId(), nil
}

func discover(host string) (*u.URL, error) {
	svchost, err := svchost.ForComparison(host)
	if err != nil {
		return nil, fmt.Errorf("invalid hostname %q: %w", host, err)
	}

	disco := svcdisco.New()
	url, err := disco.DiscoverServiceURL(svchost, "rt.v1")
	if err != nil {
		return nil, fmt.Errorf("error discovering service URL: %w", err)
	}

	return url, nil
}

func sdkWithAuthProvider(host string, atp *accessTokenProvider) (SDK, error) {
	authProvider := auth.NewBaseBearerTokenAuthenticationProvider(atp)

	adapter, err := NewRequestAdapter(authProvider, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request adapter: %w", err)
	}

	serviceURL, err := discover(host)
	if err != nil {
		return nil, err
	}

	// The generated SDK already knows about the "/api" part of the URL, so
	// it needs to be removed.
	baseURL, _ := strings.CutSuffix(serviceURL.String(), "/api")
	adapter.SetBaseUrl(baseURL)

	client := generated.NewApiClient(adapter)

	return &sdk{
		client:   client,
		http:     adapter.Client,
		endpoint: serviceURL,
	}, nil
}

func NewSDKWithAccessToken(host, accessToken string) (SDK, error) {
	validator, err := auth.NewAllowedHostsValidatorErrorCheck([]string{
		host,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid host configuration: %w", err)
	}

	tokenProvider := &accessTokenProvider{
		allowedHosts:          validator,
		accessToken:           &accessToken,
		accessTokenExpiration: time.Now().Add(876000 * time.Hour), // 100 years
		host:                  host,
	}

	return sdkWithAuthProvider(host, tokenProvider)
}

func NewSDK(host, clientID, clientSecret string) (SDK, error) {
	validator, err := auth.NewAllowedHostsValidatorErrorCheck([]string{
		host,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid host configuration: %w", err)
	}

	tokenProvider := &accessTokenProvider{
		allowedHosts: validator,
		clientID:     clientID,
		clientSecret: clientSecret,
		host:         host,
	}

	return sdkWithAuthProvider(host, tokenProvider)
}
