package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuthenticationTokensRequestBuilder builds and executes requests for operations under \api\authentication-tokens
type AuthenticationTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTokenId gets an item from the github.com/registry-tools/rt-sdk/generated.api.authenticationTokens.item collection
// returns a *AuthenticationTokensTokenItemRequestBuilder when successful
func (m *AuthenticationTokensRequestBuilder) ByTokenId(tokenId string)(*AuthenticationTokensTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tokenId != "" {
        urlTplParams["token%2Did"] = tokenId
    }
    return NewAuthenticationTokensTokenItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationTokensRequestBuilderInternal instantiates a new AuthenticationTokensRequestBuilder and sets the default values.
func NewAuthenticationTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationTokensRequestBuilder) {
    m := &AuthenticationTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/authentication-tokens", pathParameters),
    }
    return m
}
// NewAuthenticationTokensRequestBuilder instantiates a new AuthenticationTokensRequestBuilder and sets the default values.
func NewAuthenticationTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationTokensRequestBuilderInternal(urlParams, requestAdapter)
}
