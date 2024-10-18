package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// NamespacesItemServiceAccountsRequestBuilder builds and executes requests for operations under \api\namespaces\{namespace-id}\service-accounts
type NamespacesItemServiceAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NamespacesItemServiceAccountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesItemServiceAccountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NamespacesItemServiceAccountsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesItemServiceAccountsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewNamespacesItemServiceAccountsRequestBuilderInternal instantiates a new NamespacesItemServiceAccountsRequestBuilder and sets the default values.
func NewNamespacesItemServiceAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesItemServiceAccountsRequestBuilder) {
    m := &NamespacesItemServiceAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/namespaces/{namespace%2Did}/service-accounts", pathParameters),
    }
    return m
}
// NewNamespacesItemServiceAccountsRequestBuilder instantiates a new NamespacesItemServiceAccountsRequestBuilder and sets the default values.
func NewNamespacesItemServiceAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesItemServiceAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNamespacesItemServiceAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all Service Accounts
// Deprecated: This method is obsolete. Use GetAsServiceAccountsGetResponse instead.
// returns a NamespacesItemServiceAccountsResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesItemServiceAccountsRequestBuilder) Get(ctx context.Context, requestConfiguration *NamespacesItemServiceAccountsRequestBuilderGetRequestConfiguration)(NamespacesItemServiceAccountsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemServiceAccountsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemServiceAccountsResponseable), nil
}
// GetAsServiceAccountsGetResponse list all Service Accounts
// returns a NamespacesItemServiceAccountsGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesItemServiceAccountsRequestBuilder) GetAsServiceAccountsGetResponse(ctx context.Context, requestConfiguration *NamespacesItemServiceAccountsRequestBuilderGetRequestConfiguration)(NamespacesItemServiceAccountsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemServiceAccountsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemServiceAccountsGetResponseable), nil
}
// Post create a new Service Account
// Deprecated: This method is obsolete. Use PostAsServiceAccountsPostResponse instead.
// returns a NamespacesItemServiceAccountsResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesItemServiceAccountsRequestBuilder) Post(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable, requestConfiguration *NamespacesItemServiceAccountsRequestBuilderPostRequestConfiguration)(NamespacesItemServiceAccountsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemServiceAccountsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemServiceAccountsResponseable), nil
}
// PostAsServiceAccountsPostResponse create a new Service Account
// returns a NamespacesItemServiceAccountsPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesItemServiceAccountsRequestBuilder) PostAsServiceAccountsPostResponse(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable, requestConfiguration *NamespacesItemServiceAccountsRequestBuilderPostRequestConfiguration)(NamespacesItemServiceAccountsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemServiceAccountsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemServiceAccountsPostResponseable), nil
}
// ToGetRequestInformation list all Service Accounts
// returns a *RequestInformation when successful
func (m *NamespacesItemServiceAccountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NamespacesItemServiceAccountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new Service Account
// returns a *RequestInformation when successful
func (m *NamespacesItemServiceAccountsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable, requestConfiguration *NamespacesItemServiceAccountsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NamespacesItemServiceAccountsRequestBuilder when successful
func (m *NamespacesItemServiceAccountsRequestBuilder) WithUrl(rawUrl string)(*NamespacesItemServiceAccountsRequestBuilder) {
    return NewNamespacesItemServiceAccountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
