package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// NamespacesItemTagPublishersRequestBuilder builds and executes requests for operations under \api\namespaces\{namespace-id}\tag-publishers
type NamespacesItemTagPublishersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NamespacesItemTagPublishersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NamespacesItemTagPublishersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewNamespacesItemTagPublishersRequestBuilderInternal instantiates a new NamespacesItemTagPublishersRequestBuilder and sets the default values.
func NewNamespacesItemTagPublishersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesItemTagPublishersRequestBuilder) {
    m := &NamespacesItemTagPublishersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/namespaces/{namespace%2Did}/tag-publishers", pathParameters),
    }
    return m
}
// NewNamespacesItemTagPublishersRequestBuilder instantiates a new NamespacesItemTagPublishersRequestBuilder and sets the default values.
func NewNamespacesItemTagPublishersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NamespacesItemTagPublishersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNamespacesItemTagPublishersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a Tag Publisher
// Deprecated: This method is obsolete. Use PostAsTagPublishersPostResponse instead.
// returns a NamespacesItemTagPublishersResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesItemTagPublishersRequestBuilder) Post(ctx context.Context, body NamespacesItemTagPublishersPostRequestBodyable, requestConfiguration *NamespacesItemTagPublishersRequestBuilderPostRequestConfiguration)(NamespacesItemTagPublishersResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemTagPublishersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemTagPublishersResponseable), nil
}
// PostAsTagPublishersPostResponse create a Tag Publisher
// returns a NamespacesItemTagPublishersPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *NamespacesItemTagPublishersRequestBuilder) PostAsTagPublishersPostResponse(ctx context.Context, body NamespacesItemTagPublishersPostRequestBodyable, requestConfiguration *NamespacesItemTagPublishersRequestBuilderPostRequestConfiguration)(NamespacesItemTagPublishersPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNamespacesItemTagPublishersPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(NamespacesItemTagPublishersPostResponseable), nil
}
// ToPostRequestInformation create a Tag Publisher
// returns a *RequestInformation when successful
func (m *NamespacesItemTagPublishersRequestBuilder) ToPostRequestInformation(ctx context.Context, body NamespacesItemTagPublishersPostRequestBodyable, requestConfiguration *NamespacesItemTagPublishersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NamespacesItemTagPublishersRequestBuilder when successful
func (m *NamespacesItemTagPublishersRequestBuilder) WithUrl(rawUrl string)(*NamespacesItemTagPublishersRequestBuilder) {
    return NewNamespacesItemTagPublishersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
