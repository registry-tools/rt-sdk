package auth

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuthRequestBuilder builds and executes requests for operations under \auth
type AuthRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewAuthRequestBuilderInternal instantiates a new AuthRequestBuilder and sets the default values.
func NewAuthRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthRequestBuilder) {
    m := &AuthRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/auth", pathParameters),
    }
    return m
}
// NewAuthRequestBuilder instantiates a new AuthRequestBuilder and sets the default values.
func NewAuthRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthRequestBuilderInternal(urlParams, requestAdapter)
}
// Token the token property
// returns a *TokenRequestBuilder when successful
func (m *AuthRequestBuilder) Token()(*TokenRequestBuilder) {
    return NewTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
