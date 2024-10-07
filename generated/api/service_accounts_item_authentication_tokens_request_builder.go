package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// ServiceAccountsItemAuthenticationTokensRequestBuilder builds and executes requests for operations under \api\service-accounts\{service-account-id}\authentication-tokens
type ServiceAccountsItemAuthenticationTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceAccountsItemAuthenticationTokensRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAccountsItemAuthenticationTokensRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceAccountsItemAuthenticationTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAccountsItemAuthenticationTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAccountsItemAuthenticationTokensRequestBuilderInternal instantiates a new ServiceAccountsItemAuthenticationTokensRequestBuilder and sets the default values.
func NewServiceAccountsItemAuthenticationTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAccountsItemAuthenticationTokensRequestBuilder) {
    m := &ServiceAccountsItemAuthenticationTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/service-accounts/{service%2Daccount%2Did}/authentication-tokens", pathParameters),
    }
    return m
}
// NewServiceAccountsItemAuthenticationTokensRequestBuilder instantiates a new ServiceAccountsItemAuthenticationTokensRequestBuilder and sets the default values.
func NewServiceAccountsItemAuthenticationTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAccountsItemAuthenticationTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAccountsItemAuthenticationTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all Authentication Tokens for a Service Account
// Deprecated: This method is obsolete. Use GetAsAuthenticationTokensGetResponse instead.
// returns a ServiceAccountsItemAuthenticationTokensResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAccountsItemAuthenticationTokensRequestBuilderGetRequestConfiguration)(ServiceAccountsItemAuthenticationTokensResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceAccountsItemAuthenticationTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAccountsItemAuthenticationTokensResponseable), nil
}
// GetAsAuthenticationTokensGetResponse list all Authentication Tokens for a Service Account
// returns a ServiceAccountsItemAuthenticationTokensGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) GetAsAuthenticationTokensGetResponse(ctx context.Context, requestConfiguration *ServiceAccountsItemAuthenticationTokensRequestBuilderGetRequestConfiguration)(ServiceAccountsItemAuthenticationTokensGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceAccountsItemAuthenticationTokensGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAccountsItemAuthenticationTokensGetResponseable), nil
}
// Post create a new Authentication Token for a Service Account
// Deprecated: This method is obsolete. Use PostAsAuthenticationTokensPostResponse instead.
// returns a ServiceAccountsItemAuthenticationTokensResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) Post(ctx context.Context, body ServiceAccountsItemAuthenticationTokensPostRequestBodyable, requestConfiguration *ServiceAccountsItemAuthenticationTokensRequestBuilderPostRequestConfiguration)(ServiceAccountsItemAuthenticationTokensResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceAccountsItemAuthenticationTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAccountsItemAuthenticationTokensResponseable), nil
}
// PostAsAuthenticationTokensPostResponse create a new Authentication Token for a Service Account
// returns a ServiceAccountsItemAuthenticationTokensPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) PostAsAuthenticationTokensPostResponse(ctx context.Context, body ServiceAccountsItemAuthenticationTokensPostRequestBodyable, requestConfiguration *ServiceAccountsItemAuthenticationTokensRequestBuilderPostRequestConfiguration)(ServiceAccountsItemAuthenticationTokensPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceAccountsItemAuthenticationTokensPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAccountsItemAuthenticationTokensPostResponseable), nil
}
// ToGetRequestInformation list all Authentication Tokens for a Service Account
// returns a *RequestInformation when successful
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAccountsItemAuthenticationTokensRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new Authentication Token for a Service Account
// returns a *RequestInformation when successful
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceAccountsItemAuthenticationTokensPostRequestBodyable, requestConfiguration *ServiceAccountsItemAuthenticationTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServiceAccountsItemAuthenticationTokensRequestBuilder when successful
func (m *ServiceAccountsItemAuthenticationTokensRequestBuilder) WithUrl(rawUrl string)(*ServiceAccountsItemAuthenticationTokensRequestBuilder) {
    return NewServiceAccountsItemAuthenticationTokensRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
