package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// AuthenticationTokensTokenItemRequestBuilder builds and executes requests for operations under \api\authentication-tokens\{token-id}
type AuthenticationTokensTokenItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationTokensTokenItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationTokensTokenItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationTokensTokenItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationTokensTokenItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationTokensTokenItemRequestBuilderInternal instantiates a new AuthenticationTokensTokenItemRequestBuilder and sets the default values.
func NewAuthenticationTokensTokenItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationTokensTokenItemRequestBuilder) {
    m := &AuthenticationTokensTokenItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/authentication-tokens/{token%2Did}", pathParameters),
    }
    return m
}
// NewAuthenticationTokensTokenItemRequestBuilder instantiates a new AuthenticationTokensTokenItemRequestBuilder and sets the default values.
func NewAuthenticationTokensTokenItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationTokensTokenItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationTokensTokenItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an Authentication Token
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *AuthenticationTokensTokenItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationTokensTokenItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get show details of an Authentication Token
// Deprecated: This method is obsolete. Use GetAsTokenGetResponse instead.
// returns a AuthenticationTokensItemTokenResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *AuthenticationTokensTokenItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationTokensTokenItemRequestBuilderGetRequestConfiguration)(AuthenticationTokensItemTokenResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationTokensItemTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationTokensItemTokenResponseable), nil
}
// GetAsTokenGetResponse show details of an Authentication Token
// returns a AuthenticationTokensItemTokenGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *AuthenticationTokensTokenItemRequestBuilder) GetAsTokenGetResponse(ctx context.Context, requestConfiguration *AuthenticationTokensTokenItemRequestBuilderGetRequestConfiguration)(AuthenticationTokensItemTokenGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationTokensItemTokenGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationTokensItemTokenGetResponseable), nil
}
// ToDeleteRequestInformation delete an Authentication Token
// returns a *RequestInformation when successful
func (m *AuthenticationTokensTokenItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationTokensTokenItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation show details of an Authentication Token
// returns a *RequestInformation when successful
func (m *AuthenticationTokensTokenItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationTokensTokenItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthenticationTokensTokenItemRequestBuilder when successful
func (m *AuthenticationTokensTokenItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationTokensTokenItemRequestBuilder) {
    return NewAuthenticationTokensTokenItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
