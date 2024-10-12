package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ModulesRequestBuilder builds and executes requests for operations under \api\modules
type ModulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByNamespace gets an item from the github.com/registry-tools/rt-sdk/generated.api.modules.item collection
// returns a *ModulesWithNamespaceItemRequestBuilder when successful
func (m *ModulesRequestBuilder) ByNamespace(namespace string)(*ModulesWithNamespaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if namespace != "" {
        urlTplParams["namespace"] = namespace
    }
    return NewModulesWithNamespaceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewModulesRequestBuilderInternal instantiates a new ModulesRequestBuilder and sets the default values.
func NewModulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesRequestBuilder) {
    m := &ModulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/modules", pathParameters),
    }
    return m
}
// NewModulesRequestBuilder instantiates a new ModulesRequestBuilder and sets the default values.
func NewModulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModulesRequestBuilderInternal(urlParams, requestAdapter)
}
