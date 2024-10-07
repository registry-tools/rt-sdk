package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// ModulesItemItemItemVersionsRequestBuilder builds and executes requests for operations under \api\modules\{namespace}\{terraform-module-name}\{terraform-module-system}\versions
type ModulesItemItemItemVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ModulesItemItemItemVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ModulesItemItemItemVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewModulesItemItemItemVersionsRequestBuilderInternal instantiates a new ModulesItemItemItemVersionsRequestBuilder and sets the default values.
func NewModulesItemItemItemVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesItemItemItemVersionsRequestBuilder) {
    m := &ModulesItemItemItemVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/modules/{namespace}/{terraform%2Dmodule%2Dname}/{terraform%2Dmodule%2Dsystem}/versions", pathParameters),
    }
    return m
}
// NewModulesItemItemItemVersionsRequestBuilder instantiates a new ModulesItemItemItemVersionsRequestBuilder and sets the default values.
func NewModulesItemItemItemVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesItemItemItemVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModulesItemItemItemVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list available versions of a terraform module
// Deprecated: This method is obsolete. Use GetAsVersionsGetResponse instead.
// returns a ModulesItemItemItemVersionsResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ModulesItemItemItemVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ModulesItemItemItemVersionsRequestBuilderGetRequestConfiguration)(ModulesItemItemItemVersionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateModulesItemItemItemVersionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ModulesItemItemItemVersionsResponseable), nil
}
// GetAsVersionsGetResponse list available versions of a terraform module
// returns a ModulesItemItemItemVersionsGetResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ModulesItemItemItemVersionsRequestBuilder) GetAsVersionsGetResponse(ctx context.Context, requestConfiguration *ModulesItemItemItemVersionsRequestBuilderGetRequestConfiguration)(ModulesItemItemItemVersionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateModulesItemItemItemVersionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ModulesItemItemItemVersionsGetResponseable), nil
}
// ToGetRequestInformation list available versions of a terraform module
// returns a *RequestInformation when successful
func (m *ModulesItemItemItemVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ModulesItemItemItemVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ModulesItemItemItemVersionsRequestBuilder when successful
func (m *ModulesItemItemItemVersionsRequestBuilder) WithUrl(rawUrl string)(*ModulesItemItemItemVersionsRequestBuilder) {
    return NewModulesItemItemItemVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
