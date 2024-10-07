package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// VcsConnectorsRequestBuilder builds and executes requests for operations under \api\vcs-connectors
type VcsConnectorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VcsConnectorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VcsConnectorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByVcsConnectorId gets an item from the github.com/registry-tools/rt-sdk/generated.api.vcsConnectors.item collection
// returns a *VcsConnectorsVcsConnectorItemRequestBuilder when successful
func (m *VcsConnectorsRequestBuilder) ByVcsConnectorId(vcsConnectorId string)(*VcsConnectorsVcsConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if vcsConnectorId != "" {
        urlTplParams["vcs%2Dconnector%2Did"] = vcsConnectorId
    }
    return NewVcsConnectorsVcsConnectorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVcsConnectorsRequestBuilderInternal instantiates a new VcsConnectorsRequestBuilder and sets the default values.
func NewVcsConnectorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VcsConnectorsRequestBuilder) {
    m := &VcsConnectorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/vcs-connectors", pathParameters),
    }
    return m
}
// NewVcsConnectorsRequestBuilder instantiates a new VcsConnectorsRequestBuilder and sets the default values.
func NewVcsConnectorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VcsConnectorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVcsConnectorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new VCS Connector
// Deprecated: This method is obsolete. Use PostAsVcsConnectorsPostResponse instead.
// returns a VcsConnectorsResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *VcsConnectorsRequestBuilder) Post(ctx context.Context, body VcsConnectorsPostRequestBodyable, requestConfiguration *VcsConnectorsRequestBuilderPostRequestConfiguration)(VcsConnectorsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVcsConnectorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VcsConnectorsResponseable), nil
}
// PostAsVcsConnectorsPostResponse create a new VCS Connector
// returns a VcsConnectorsPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *VcsConnectorsRequestBuilder) PostAsVcsConnectorsPostResponse(ctx context.Context, body VcsConnectorsPostRequestBodyable, requestConfiguration *VcsConnectorsRequestBuilderPostRequestConfiguration)(VcsConnectorsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVcsConnectorsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VcsConnectorsPostResponseable), nil
}
// ToPostRequestInformation create a new VCS Connector
// returns a *RequestInformation when successful
func (m *VcsConnectorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VcsConnectorsPostRequestBodyable, requestConfiguration *VcsConnectorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VcsConnectorsRequestBuilder when successful
func (m *VcsConnectorsRequestBuilder) WithUrl(rawUrl string)(*VcsConnectorsRequestBuilder) {
    return NewVcsConnectorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
