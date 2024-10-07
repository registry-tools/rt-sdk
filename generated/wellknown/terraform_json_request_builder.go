package wellknown

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TerraformJsonRequestBuilder builds and executes requests for operations under \.well-known\terraform.json
type TerraformJsonRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TerraformJsonRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TerraformJsonRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTerraformJsonRequestBuilderInternal instantiates a new TerraformJsonRequestBuilder and sets the default values.
func NewTerraformJsonRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TerraformJsonRequestBuilder) {
    m := &TerraformJsonRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/.well-known/terraform.json", pathParameters),
    }
    return m
}
// NewTerraformJsonRequestBuilder instantiates a new TerraformJsonRequestBuilder and sets the default values.
func NewTerraformJsonRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TerraformJsonRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTerraformJsonRequestBuilderInternal(urlParams, requestAdapter)
}
// Get service discovery for terraform compatibility
// Deprecated: This method is obsolete. Use GetAsTerraformGetResponse instead.
// returns a TerraformJsonTerraformResponseable when successful
func (m *TerraformJsonRequestBuilder) Get(ctx context.Context, requestConfiguration *TerraformJsonRequestBuilderGetRequestConfiguration)(TerraformJsonTerraformResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTerraformJsonTerraformResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TerraformJsonTerraformResponseable), nil
}
// GetAsTerraformGetResponse service discovery for terraform compatibility
// returns a TerraformJsonTerraformGetResponseable when successful
func (m *TerraformJsonRequestBuilder) GetAsTerraformGetResponse(ctx context.Context, requestConfiguration *TerraformJsonRequestBuilderGetRequestConfiguration)(TerraformJsonTerraformGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTerraformJsonTerraformGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TerraformJsonTerraformGetResponseable), nil
}
// ToGetRequestInformation service discovery for terraform compatibility
// returns a *RequestInformation when successful
func (m *TerraformJsonRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TerraformJsonRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TerraformJsonRequestBuilder when successful
func (m *TerraformJsonRequestBuilder) WithUrl(rawUrl string)(*TerraformJsonRequestBuilder) {
    return NewTerraformJsonRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
