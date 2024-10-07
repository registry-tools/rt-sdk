package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// TagPublishersTagPublishersItemRequestBuilder builds and executes requests for operations under \api\tag-publishers\{id}
type TagPublishersTagPublishersItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TagPublishersTagPublishersItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagPublishersTagPublishersItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TagPublishersTagPublishersItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagPublishersTagPublishersItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTagPublishersTagPublishersItemRequestBuilderInternal instantiates a new TagPublishersTagPublishersItemRequestBuilder and sets the default values.
func NewTagPublishersTagPublishersItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TagPublishersTagPublishersItemRequestBuilder) {
    m := &TagPublishersTagPublishersItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/tag-publishers/{id}", pathParameters),
    }
    return m
}
// NewTagPublishersTagPublishersItemRequestBuilder instantiates a new TagPublishersTagPublishersItemRequestBuilder and sets the default values.
func NewTagPublishersTagPublishersItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TagPublishersTagPublishersItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTagPublishersTagPublishersItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Tag Publisher
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *TagPublishersTagPublishersItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TagPublishersTagPublishersItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get show details of a Tag Publisher
// Deprecated: This method is obsolete. Use GetAsTagPublishersGetResponse instead.
// returns a TagPublishersItemTagPublishersResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *TagPublishersTagPublishersItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TagPublishersTagPublishersItemRequestBuilderGetRequestConfiguration)(TagPublishersItemTagPublishersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagPublishersItemTagPublishersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TagPublishersItemTagPublishersResponseable), nil
}
// GetAsTagPublishersGetResponse show details of a Tag Publisher
// returns a TagPublishersItemTagPublishersGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *TagPublishersTagPublishersItemRequestBuilder) GetAsTagPublishersGetResponse(ctx context.Context, requestConfiguration *TagPublishersTagPublishersItemRequestBuilderGetRequestConfiguration)(TagPublishersItemTagPublishersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagPublishersItemTagPublishersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TagPublishersItemTagPublishersGetResponseable), nil
}
// ToDeleteRequestInformation delete a Tag Publisher
// returns a *RequestInformation when successful
func (m *TagPublishersTagPublishersItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TagPublishersTagPublishersItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation show details of a Tag Publisher
// returns a *RequestInformation when successful
func (m *TagPublishersTagPublishersItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagPublishersTagPublishersItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TagPublishersTagPublishersItemRequestBuilder when successful
func (m *TagPublishersTagPublishersItemRequestBuilder) WithUrl(rawUrl string)(*TagPublishersTagPublishersItemRequestBuilder) {
    return NewTagPublishersTagPublishersItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
