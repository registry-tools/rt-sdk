package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ModulesWithNamespaceItemRequestBuilder builds and executes requests for operations under \api\modules\{namespace}
type ModulesWithNamespaceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTerraformModuleName gets an item from the github.com/registry-tools/rt-sdk/generated.api.modules.item.item collection
// returns a *ModulesItemTerraformModuleNameItemRequestBuilder when successful
func (m *ModulesWithNamespaceItemRequestBuilder) ByTerraformModuleName(terraformModuleName string)(*ModulesItemTerraformModuleNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if terraformModuleName != "" {
        urlTplParams["terraform%2Dmodule%2Dname"] = terraformModuleName
    }
    return NewModulesItemTerraformModuleNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewModulesWithNamespaceItemRequestBuilderInternal instantiates a new ModulesWithNamespaceItemRequestBuilder and sets the default values.
func NewModulesWithNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesWithNamespaceItemRequestBuilder) {
    m := &ModulesWithNamespaceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/modules/{namespace}", pathParameters),
    }
    return m
}
// NewModulesWithNamespaceItemRequestBuilder instantiates a new ModulesWithNamespaceItemRequestBuilder and sets the default values.
func NewModulesWithNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesWithNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModulesWithNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
