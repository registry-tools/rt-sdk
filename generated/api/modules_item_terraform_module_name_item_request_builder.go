package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ModulesItemTerraformModuleNameItemRequestBuilder builds and executes requests for operations under \api\modules\{namespace}\{terraform-module-name}
type ModulesItemTerraformModuleNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTerraformModuleSystem gets an item from the github.com/registry-tools/rt-sdk/generated.api.modules.item.item.item collection
// returns a *ModulesItemItemTerraformModuleSystemItemRequestBuilder when successful
func (m *ModulesItemTerraformModuleNameItemRequestBuilder) ByTerraformModuleSystem(terraformModuleSystem string)(*ModulesItemItemTerraformModuleSystemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if terraformModuleSystem != "" {
        urlTplParams["terraform%2Dmodule%2Dsystem"] = terraformModuleSystem
    }
    return NewModulesItemItemTerraformModuleSystemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewModulesItemTerraformModuleNameItemRequestBuilderInternal instantiates a new ModulesItemTerraformModuleNameItemRequestBuilder and sets the default values.
func NewModulesItemTerraformModuleNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesItemTerraformModuleNameItemRequestBuilder) {
    m := &ModulesItemTerraformModuleNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/modules/{namespace}/{terraform%2Dmodule%2Dname}", pathParameters),
    }
    return m
}
// NewModulesItemTerraformModuleNameItemRequestBuilder instantiates a new ModulesItemTerraformModuleNameItemRequestBuilder and sets the default values.
func NewModulesItemTerraformModuleNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesItemTerraformModuleNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModulesItemTerraformModuleNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
