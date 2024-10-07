package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// ModulesPublishRequestBuilder builds and executes requests for operations under \api\modules\publish
type ModulesPublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ModulesPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ModulesPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewModulesPublishRequestBuilderInternal instantiates a new ModulesPublishRequestBuilder and sets the default values.
func NewModulesPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesPublishRequestBuilder) {
    m := &ModulesPublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/modules/publish", pathParameters),
    }
    return m
}
// NewModulesPublishRequestBuilder instantiates a new ModulesPublishRequestBuilder and sets the default values.
func NewModulesPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModulesPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post publish a Terraform Module Version
// Deprecated: This method is obsolete. Use PostAsPublishPostResponse instead.
// returns a ModulesPublishResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ModulesPublishRequestBuilder) Post(ctx context.Context, body ModulesPublishPostRequestBodyable, requestConfiguration *ModulesPublishRequestBuilderPostRequestConfiguration)(ModulesPublishResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateModulesPublishResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ModulesPublishResponseable), nil
}
// PostAsPublishPostResponse publish a Terraform Module Version
// returns a ModulesPublishPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ModulesPublishRequestBuilder) PostAsPublishPostResponse(ctx context.Context, body ModulesPublishPostRequestBodyable, requestConfiguration *ModulesPublishRequestBuilderPostRequestConfiguration)(ModulesPublishPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateModulesPublishPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ModulesPublishPostResponseable), nil
}
// ToPostRequestInformation publish a Terraform Module Version
// returns a *RequestInformation when successful
func (m *ModulesPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, body ModulesPublishPostRequestBodyable, requestConfiguration *ModulesPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ModulesPublishRequestBuilder when successful
func (m *ModulesPublishRequestBuilder) WithUrl(rawUrl string)(*ModulesPublishRequestBuilder) {
    return NewModulesPublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
