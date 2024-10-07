package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApiRequestBuilder builds and executes requests for operations under \api
type ApiRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationTokens the authenticationTokens property
// returns a *AuthenticationTokensRequestBuilder when successful
func (m *ApiRequestBuilder) AuthenticationTokens()(*AuthenticationTokensRequestBuilder) {
    return NewAuthenticationTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiRequestBuilderInternal instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    m := &ApiRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api", pathParameters),
    }
    return m
}
// NewApiRequestBuilder instantiates a new ApiRequestBuilder and sets the default values.
func NewApiRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApiRequestBuilderInternal(urlParams, requestAdapter)
}
// Modules the modules property
// returns a *ModulesRequestBuilder when successful
func (m *ApiRequestBuilder) Modules()(*ModulesRequestBuilder) {
    return NewModulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Namespaces the namespaces property
// returns a *NamespacesRequestBuilder when successful
func (m *ApiRequestBuilder) Namespaces()(*NamespacesRequestBuilder) {
    return NewNamespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceAccounts the serviceAccounts property
// returns a *ServiceAccountsRequestBuilder when successful
func (m *ApiRequestBuilder) ServiceAccounts()(*ServiceAccountsRequestBuilder) {
    return NewServiceAccountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TagPublishers the tagPublishers property
// returns a *TagPublishersRequestBuilder when successful
func (m *ApiRequestBuilder) TagPublishers()(*TagPublishersRequestBuilder) {
    return NewTagPublishersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VcsConnectors the vcsConnectors property
// returns a *VcsConnectorsRequestBuilder when successful
func (m *ApiRequestBuilder) VcsConnectors()(*VcsConnectorsRequestBuilder) {
    return NewVcsConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
