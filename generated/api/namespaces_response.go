package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NamespacesGetResponseable instead.
type NamespacesResponse struct {
    NamespacesGetResponse
}
// NewNamespacesResponse instantiates a new NamespacesResponse and sets the default values.
func NewNamespacesResponse()(*NamespacesResponse) {
    m := &NamespacesResponse{
        NamespacesGetResponse: *NewNamespacesGetResponse(),
    }
    return m
}
// CreateNamespacesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesResponse(), nil
}
// Deprecated: This class is obsolete. Use NamespacesGetResponseable instead.
type NamespacesResponseable interface {
    NamespacesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
