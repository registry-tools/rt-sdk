package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// NamespacesNamespaceItemRequestBuilder builds and executes requests for operations under \api\namespaces\{namespace-id}
type NamespacesNamespaceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NamespacesNamespaceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesNamespaceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NamespacesNamespaceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesNamespaceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NamespacesNamespaceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesNamespaceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewNamespacesNamespaceItemRequestBuilderInternal instantiates a new NamespacesNamespaceItemRequestBuilder and sets the default values.
func NewNamespacesNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesNamespaceItemRequestBuilder) {
    m := &NamespacesNamespaceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/namespaces/{namespace%2Did}", pathParameters),
    }
    return m
}
// NewNamespacesNamespaceItemRequestBuilder instantiates a new NamespacesNamespaceItemRequestBuilder and sets the default values.
func NewNamespacesNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNamespacesNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Namespace
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesNamespaceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *NamespacesNamespaceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get show details of a Namespace
// Deprecated: This method is obsolete. Use GetAsNamespaceGetResponse instead.
// returns a NamespacesItemNamespaceResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesNamespaceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *NamespacesNamespaceItemRequestBuilderGetRequestConfiguration)(NamespacesItemNamespaceResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemNamespaceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemNamespaceResponseable), nil
}
// GetAsNamespaceGetResponse show details of a Namespace
// returns a NamespacesItemNamespaceGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesNamespaceItemRequestBuilder) GetAsNamespaceGetResponse(ctx context.Context, requestConfiguration *NamespacesNamespaceItemRequestBuilderGetRequestConfiguration)(NamespacesItemNamespaceGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemNamespaceGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemNamespaceGetResponseable), nil
}
// Patch update a Namespace
// Deprecated: This method is obsolete. Use PatchAsNamespacePatchResponse instead.
// returns a NamespacesItemNamespaceResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesNamespaceItemRequestBuilder) Patch(ctx context.Context, body NamespacesItemNamespacePatchRequestBodyable, requestConfiguration *NamespacesNamespaceItemRequestBuilderPatchRequestConfiguration)(NamespacesItemNamespaceResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemNamespaceResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemNamespaceResponseable), nil
}
// PatchAsNamespacePatchResponse update a Namespace
// returns a NamespacesItemNamespacePatchResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesNamespaceItemRequestBuilder) PatchAsNamespacePatchResponse(ctx context.Context, body NamespacesItemNamespacePatchRequestBodyable, requestConfiguration *NamespacesNamespaceItemRequestBuilderPatchRequestConfiguration)(NamespacesItemNamespacePatchResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemNamespacePatchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemNamespacePatchResponseable), nil
}
// ServiceAccounts the serviceAccounts property
// returns a *NamespacesItemServiceAccountsRequestBuilder when successful
func (m *NamespacesNamespaceItemRequestBuilder) ServiceAccounts()(*NamespacesItemServiceAccountsRequestBuilder) {
    return NewNamespacesItemServiceAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TagPublishers the tagPublishers property
// returns a *NamespacesItemTagPublishersRequestBuilder when successful
func (m *NamespacesNamespaceItemRequestBuilder) TagPublishers()(*NamespacesItemTagPublishersRequestBuilder) {
    return NewNamespacesItemTagPublishersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a Namespace
// returns a *RequestInformation when successful
func (m *NamespacesNamespaceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *NamespacesNamespaceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation show details of a Namespace
// returns a *RequestInformation when successful
func (m *NamespacesNamespaceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NamespacesNamespaceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update a Namespace
// returns a *RequestInformation when successful
func (m *NamespacesNamespaceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body NamespacesItemNamespacePatchRequestBodyable, requestConfiguration *NamespacesNamespaceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *NamespacesNamespaceItemRequestBuilder when successful
func (m *NamespacesNamespaceItemRequestBuilder) WithUrl(rawUrl string)(*NamespacesNamespaceItemRequestBuilder) {
    return NewNamespacesNamespaceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
