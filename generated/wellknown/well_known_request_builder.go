package wellknown

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WellKnownRequestBuilder builds and executes requests for operations under \.well-known
type WellKnownRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWellKnownRequestBuilderInternal instantiates a new WellKnownRequestBuilder and sets the default values.
func NewWellKnownRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WellKnownRequestBuilder) {
    m := &WellKnownRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/.well-known", pathParameters),
    }
    return m
}
// NewWellKnownRequestBuilder instantiates a new WellKnownRequestBuilder and sets the default values.
func NewWellKnownRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WellKnownRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWellKnownRequestBuilderInternal(urlParams, requestAdapter)
}
// TerraformJson the terraformJson property
// returns a *TerraformJsonRequestBuilder when successful
func (m *WellKnownRequestBuilder) TerraformJson()(*TerraformJsonRequestBuilder) {
    return NewTerraformJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
