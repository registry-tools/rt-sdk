package api

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagPublishersRequestBuilder builds and executes requests for operations under \api\tag-publishers
type TagPublishersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ById gets an item from the github.com/registry-tools/rt-sdk/generated.api.tagPublishers.item collection
// returns a *TagPublishersTagPublishersItemRequestBuilder when successful
func (m *TagPublishersRequestBuilder) ById(id string)(*TagPublishersTagPublishersItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewTagPublishersTagPublishersItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTagPublishersRequestBuilderInternal instantiates a new TagPublishersRequestBuilder and sets the default values.
func NewTagPublishersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TagPublishersRequestBuilder) {
    m := &TagPublishersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/tag-publishers", pathParameters),
    }
    return m
}
// NewTagPublishersRequestBuilder instantiates a new TagPublishersRequestBuilder and sets the default values.
func NewTagPublishersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TagPublishersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTagPublishersRequestBuilderInternal(urlParams, requestAdapter)
}
