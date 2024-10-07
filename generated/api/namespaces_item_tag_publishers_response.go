package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NamespacesItemTagPublishersPostResponseable instead.
type NamespacesItemTagPublishersResponse struct {
    NamespacesItemTagPublishersPostResponse
}
// NewNamespacesItemTagPublishersResponse instantiates a new NamespacesItemTagPublishersResponse and sets the default values.
func NewNamespacesItemTagPublishersResponse()(*NamespacesItemTagPublishersResponse) {
    m := &NamespacesItemTagPublishersResponse{
        NamespacesItemTagPublishersPostResponse: *NewNamespacesItemTagPublishersPostResponse(),
    }
    return m
}
// CreateNamespacesItemTagPublishersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemTagPublishersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemTagPublishersResponse(), nil
}
// Deprecated: This class is obsolete. Use NamespacesItemTagPublishersPostResponseable instead.
type NamespacesItemTagPublishersResponseable interface {
    NamespacesItemTagPublishersPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
