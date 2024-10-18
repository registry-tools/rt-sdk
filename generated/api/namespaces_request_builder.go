package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// NamespacesRequestBuilder builds and executes requests for operations under \api\namespaces
type NamespacesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NamespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NamespacesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByNamespaceId gets an item from the github.com/registry-tools/rt-sdk/generated.api.namespaces.item collection
// returns a *NamespacesNamespaceItemRequestBuilder when successful
func (m *NamespacesRequestBuilder) ByNamespaceId(namespaceId string)(*NamespacesNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if namespaceId != "" {
        urlTplParams["namespace%2Did"] = namespaceId
    }
    return NewNamespacesNamespaceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewNamespacesRequestBuilderInternal instantiates a new NamespacesRequestBuilder and sets the default values.
func NewNamespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesRequestBuilder) {
    m := &NamespacesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/namespaces", pathParameters),
    }
    return m
}
// NewNamespacesRequestBuilder instantiates a new NamespacesRequestBuilder and sets the default values.
func NewNamespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNamespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all Namespaces
// Deprecated: This method is obsolete. Use GetAsNamespacesGetResponse instead.
// returns a NamespacesResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *NamespacesRequestBuilderGetRequestConfiguration)(NamespacesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesResponseable), nil
}
// GetAsNamespacesGetResponse list all Namespaces
// returns a NamespacesGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesRequestBuilder) GetAsNamespacesGetResponse(ctx context.Context, requestConfiguration *NamespacesRequestBuilderGetRequestConfiguration)(NamespacesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesGetResponseable), nil
}
// Post create a new Namespace
// Deprecated: This method is obsolete. Use PostAsNamespacesPostResponse instead.
// returns a NamespacesResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesRequestBuilder) Post(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable, requestConfiguration *NamespacesRequestBuilderPostRequestConfiguration)(NamespacesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesResponseable), nil
}
// PostAsNamespacesPostResponse create a new Namespace
// returns a NamespacesPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesRequestBuilder) PostAsNamespacesPostResponse(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable, requestConfiguration *NamespacesRequestBuilderPostRequestConfiguration)(NamespacesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesPostResponseable), nil
}
// ToGetRequestInformation list all Namespaces
// returns a *RequestInformation when successful
func (m *NamespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NamespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new Namespace
// returns a *RequestInformation when successful
func (m *NamespacesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable, requestConfiguration *NamespacesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NamespacesRequestBuilder when successful
func (m *NamespacesRequestBuilder) WithUrl(rawUrl string)(*NamespacesRequestBuilder) {
    return NewNamespacesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
