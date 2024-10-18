package api

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

// ArchivesRequestBuilder builds and executes requests for operations under \api\archives
type ArchivesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ArchivesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ArchivesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewArchivesRequestBuilderInternal instantiates a new ArchivesRequestBuilder and sets the default values.
func NewArchivesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ArchivesRequestBuilder) {
    m := &ArchivesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/archives", pathParameters),
    }
    return m
}
// NewArchivesRequestBuilder instantiates a new ArchivesRequestBuilder and sets the default values.
func NewArchivesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ArchivesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewArchivesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post register a new file archive to be uploaded to the URL indicated in the response
// Deprecated: This method is obsolete. Use PostAsArchivesPostResponse instead.
// returns a ArchivesResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ArchivesRequestBuilder) Post(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Archiveable, requestConfiguration *ArchivesRequestBuilderPostRequestConfiguration)(ArchivesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateArchivesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ArchivesResponseable), nil
}
// PostAsArchivesPostResponse register a new file archive to be uploaded to the URL indicated in the response
// returns a ArchivesPostResponseable when successful
// returns a Errors error when the service returns a 4XX or 5XX status code
func (m *ArchivesRequestBuilder) PostAsArchivesPostResponse(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Archiveable, requestConfiguration *ArchivesRequestBuilderPostRequestConfiguration)(ArchivesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateErrorsFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateArchivesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ArchivesPostResponseable), nil
}
// ToPostRequestInformation register a new file archive to be uploaded to the URL indicated in the response
// returns a *RequestInformation when successful
func (m *ArchivesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Archiveable, requestConfiguration *ArchivesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ArchivesRequestBuilder when successful
func (m *ArchivesRequestBuilder) WithUrl(rawUrl string)(*ArchivesRequestBuilder) {
    return NewArchivesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
