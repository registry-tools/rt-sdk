package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagPublishersItemTagPublishersGetResponseable instead.
type TagPublishersItemTagPublishersResponse struct {
    TagPublishersItemTagPublishersGetResponse
}
// NewTagPublishersItemTagPublishersResponse instantiates a new TagPublishersItemTagPublishersResponse and sets the default values.
func NewTagPublishersItemTagPublishersResponse()(*TagPublishersItemTagPublishersResponse) {
    m := &TagPublishersItemTagPublishersResponse{
        TagPublishersItemTagPublishersGetResponse: *NewTagPublishersItemTagPublishersGetResponse(),
    }
    return m
}
// CreateTagPublishersItemTagPublishersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagPublishersItemTagPublishersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTagPublishersItemTagPublishersResponse(), nil
}
// Deprecated: This class is obsolete. Use TagPublishersItemTagPublishersGetResponseable instead.
type TagPublishersItemTagPublishersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TagPublishersItemTagPublishersGetResponseable
}
