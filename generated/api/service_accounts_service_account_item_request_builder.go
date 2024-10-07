package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// ServiceAccountsServiceAccountItemRequestBuilder builds and executes requests for operations under \api\service-accounts\{service-account-id}
type ServiceAccountsServiceAccountItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceAccountsServiceAccountItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAccountsServiceAccountItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceAccountsServiceAccountItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAccountsServiceAccountItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationTokens the authenticationTokens property
// returns a *ServiceAccountsItemAuthenticationTokensRequestBuilder when successful
func (m *ServiceAccountsServiceAccountItemRequestBuilder) AuthenticationTokens()(*ServiceAccountsItemAuthenticationTokensRequestBuilder) {
    return NewServiceAccountsItemAuthenticationTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceAccountsServiceAccountItemRequestBuilderInternal instantiates a new ServiceAccountsServiceAccountItemRequestBuilder and sets the default values.
func NewServiceAccountsServiceAccountItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAccountsServiceAccountItemRequestBuilder) {
    m := &ServiceAccountsServiceAccountItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/service-accounts/{service%2Daccount%2Did}", pathParameters),
    }
    return m
}
// NewServiceAccountsServiceAccountItemRequestBuilder instantiates a new ServiceAccountsServiceAccountItemRequestBuilder and sets the default values.
func NewServiceAccountsServiceAccountItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAccountsServiceAccountItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAccountsServiceAccountItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Service Account
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsServiceAccountItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceAccountsServiceAccountItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get show details of a Service Account
// Deprecated: This method is obsolete. Use GetAsServiceAccountGetResponse instead.
// returns a ServiceAccountsItemServiceAccountResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsServiceAccountItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAccountsServiceAccountItemRequestBuilderGetRequestConfiguration)(ServiceAccountsItemServiceAccountResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceAccountsItemServiceAccountResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAccountsItemServiceAccountResponseable), nil
}
// GetAsServiceAccountGetResponse show details of a Service Account
// returns a ServiceAccountsItemServiceAccountGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ServiceAccountsServiceAccountItemRequestBuilder) GetAsServiceAccountGetResponse(ctx context.Context, requestConfiguration *ServiceAccountsServiceAccountItemRequestBuilderGetRequestConfiguration)(ServiceAccountsItemServiceAccountGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceAccountsItemServiceAccountGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAccountsItemServiceAccountGetResponseable), nil
}
// ToDeleteRequestInformation delete a Service Account
// returns a *RequestInformation when successful
func (m *ServiceAccountsServiceAccountItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceAccountsServiceAccountItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation show details of a Service Account
// returns a *RequestInformation when successful
func (m *ServiceAccountsServiceAccountItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAccountsServiceAccountItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServiceAccountsServiceAccountItemRequestBuilder when successful
func (m *ServiceAccountsServiceAccountItemRequestBuilder) WithUrl(rawUrl string)(*ServiceAccountsServiceAccountItemRequestBuilder) {
    return NewServiceAccountsServiceAccountItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
