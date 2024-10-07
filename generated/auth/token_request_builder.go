package auth

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TokenRequestBuilder builds and executes requests for operations under \auth\token
type TokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTokenRequestBuilderInternal instantiates a new TokenRequestBuilder and sets the default values.
func NewTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenRequestBuilder) {
    m := &TokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/auth/token", pathParameters),
    }
    return m
}
// NewTokenRequestBuilder instantiates a new TokenRequestBuilder and sets the default values.
func NewTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post exchange credentials for an access token
// Deprecated: This method is obsolete. Use PostAsTokenPostResponse instead.
// returns a TokenResponseable when successful
// returns a Token400Error error when the service returns a 400 status code
func (m *TokenRequestBuilder) Post(ctx context.Context, body TokenPostRequestBodyable, requestConfiguration *TokenRequestBuilderPostRequestConfiguration)(TokenResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateToken400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TokenResponseable), nil
}
// PostAsTokenPostResponse exchange credentials for an access token
// returns a TokenPostResponseable when successful
// returns a Token400Error error when the service returns a 400 status code
func (m *TokenRequestBuilder) PostAsTokenPostResponse(ctx context.Context, body TokenPostRequestBodyable, requestConfiguration *TokenRequestBuilderPostRequestConfiguration)(TokenPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": CreateToken400ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTokenPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TokenPostResponseable), nil
}
// ToPostRequestInformation exchange credentials for an access token
// returns a *RequestInformation when successful
func (m *TokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body TokenPostRequestBodyable, requestConfiguration *TokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TokenRequestBuilder when successful
func (m *TokenRequestBuilder) WithUrl(rawUrl string)(*TokenRequestBuilder) {
    return NewTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
