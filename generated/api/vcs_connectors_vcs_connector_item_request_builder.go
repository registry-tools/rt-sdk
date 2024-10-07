package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// VcsConnectorsVcsConnectorItemRequestBuilder builds and executes requests for operations under \api\vcs-connectors\{vcs-connector-id}
type VcsConnectorsVcsConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VcsConnectorsVcsConnectorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VcsConnectorsVcsConnectorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVcsConnectorsVcsConnectorItemRequestBuilderInternal instantiates a new VcsConnectorsVcsConnectorItemRequestBuilder and sets the default values.
func NewVcsConnectorsVcsConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VcsConnectorsVcsConnectorItemRequestBuilder) {
    m := &VcsConnectorsVcsConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/vcs-connectors/{vcs%2Dconnector%2Did}", pathParameters),
    }
    return m
}
// NewVcsConnectorsVcsConnectorItemRequestBuilder instantiates a new VcsConnectorsVcsConnectorItemRequestBuilder and sets the default values.
func NewVcsConnectorsVcsConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VcsConnectorsVcsConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVcsConnectorsVcsConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a VCS Connector
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *VcsConnectorsVcsConnectorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VcsConnectorsVcsConnectorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete a VCS Connector
// returns a *RequestInformation when successful
func (m *VcsConnectorsVcsConnectorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VcsConnectorsVcsConnectorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VcsConnectorsVcsConnectorItemRequestBuilder when successful
func (m *VcsConnectorsVcsConnectorItemRequestBuilder) WithUrl(rawUrl string)(*VcsConnectorsVcsConnectorItemRequestBuilder) {
    return NewVcsConnectorsVcsConnectorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
