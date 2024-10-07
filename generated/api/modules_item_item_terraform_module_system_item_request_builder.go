package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ModulesItemItemTerraformModuleSystemItemRequestBuilder builds and executes requests for operations under \api\modules\{namespace}\{terraform-module-name}\{terraform-module-system}
type ModulesItemItemTerraformModuleSystemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewModulesItemItemTerraformModuleSystemItemRequestBuilderInternal instantiates a new ModulesItemItemTerraformModuleSystemItemRequestBuilder and sets the default values.
func NewModulesItemItemTerraformModuleSystemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesItemItemTerraformModuleSystemItemRequestBuilder) {
    m := &ModulesItemItemTerraformModuleSystemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/modules/{namespace}/{terraform%2Dmodule%2Dname}/{terraform%2Dmodule%2Dsystem}", pathParameters),
    }
    return m
}
// NewModulesItemItemTerraformModuleSystemItemRequestBuilder instantiates a new ModulesItemItemTerraformModuleSystemItemRequestBuilder and sets the default values.
func NewModulesItemItemTerraformModuleSystemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModulesItemItemTerraformModuleSystemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModulesItemItemTerraformModuleSystemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Versions the versions property
// returns a *ModulesItemItemItemVersionsRequestBuilder when successful
func (m *ModulesItemItemTerraformModuleSystemItemRequestBuilder) Versions()(*ModulesItemItemItemVersionsRequestBuilder) {
    return NewModulesItemItemItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
