package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// TerraformModuleVersionsRequestBuilder builds and executes requests for operations under \api\terraform-module-versions
type TerraformModuleVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TerraformModuleVersionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TerraformModuleVersionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTerraformModuleVersionsRequestBuilderInternal instantiates a new TerraformModuleVersionsRequestBuilder and sets the default values.
func NewTerraformModuleVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TerraformModuleVersionsRequestBuilder) {
    m := &TerraformModuleVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/terraform-module-versions", pathParameters),
    }
    return m
}
// NewTerraformModuleVersionsRequestBuilder instantiates a new TerraformModuleVersionsRequestBuilder and sets the default values.
func NewTerraformModuleVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TerraformModuleVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTerraformModuleVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post publish a Terraform Module Version
// Deprecated: This method is obsolete. Use PostAsTerraformModuleVersionsPostResponse instead.
// returns a TerraformModuleVersionsResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *TerraformModuleVersionsRequestBuilder) Post(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.TerraformModuleVersionable, requestConfiguration *TerraformModuleVersionsRequestBuilderPostRequestConfiguration)(TerraformModuleVersionsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTerraformModuleVersionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TerraformModuleVersionsResponseable), nil
}
// PostAsTerraformModuleVersionsPostResponse publish a Terraform Module Version
// returns a TerraformModuleVersionsPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *TerraformModuleVersionsRequestBuilder) PostAsTerraformModuleVersionsPostResponse(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.TerraformModuleVersionable, requestConfiguration *TerraformModuleVersionsRequestBuilderPostRequestConfiguration)(TerraformModuleVersionsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTerraformModuleVersionsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TerraformModuleVersionsPostResponseable), nil
}
// ToPostRequestInformation publish a Terraform Module Version
// returns a *RequestInformation when successful
func (m *TerraformModuleVersionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.TerraformModuleVersionable, requestConfiguration *TerraformModuleVersionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TerraformModuleVersionsRequestBuilder when successful
func (m *TerraformModuleVersionsRequestBuilder) WithUrl(rawUrl string)(*TerraformModuleVersionsRequestBuilder) {
    return NewTerraformModuleVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
