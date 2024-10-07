package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ServiceAccountsRequestBuilder builds and executes requests for operations under \api\service-accounts
type ServiceAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByServiceAccountId gets an item from the github.com/registry-tools/rt-sdk/generated.api.serviceAccounts.item collection
// returns a *ServiceAccountsServiceAccountItemRequestBuilder when successful
func (m *ServiceAccountsRequestBuilder) ByServiceAccountId(serviceAccountId string)(*ServiceAccountsServiceAccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if serviceAccountId != "" {
        urlTplParams["service%2Daccount%2Did"] = serviceAccountId
    }
    return NewServiceAccountsServiceAccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceAccountsRequestBuilderInternal instantiates a new ServiceAccountsRequestBuilder and sets the default values.
func NewServiceAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAccountsRequestBuilder) {
    m := &ServiceAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/service-accounts", pathParameters),
    }
    return m
}
// NewServiceAccountsRequestBuilder instantiates a new ServiceAccountsRequestBuilder and sets the default values.
func NewServiceAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
