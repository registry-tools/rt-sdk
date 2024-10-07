package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NamespacesItemNamespaceGetResponseable instead.
type NamespacesItemNamespaceResponse struct {
    NamespacesItemNamespaceGetResponse
}
// NewNamespacesItemNamespaceResponse instantiates a new NamespacesItemNamespaceResponse and sets the default values.
func NewNamespacesItemNamespaceResponse()(*NamespacesItemNamespaceResponse) {
    m := &NamespacesItemNamespaceResponse{
        NamespacesItemNamespaceGetResponse: *NewNamespacesItemNamespaceGetResponse(),
    }
    return m
}
// CreateNamespacesItemNamespaceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemNamespaceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemNamespaceResponse(), nil
}
// Deprecated: This class is obsolete. Use NamespacesItemNamespaceGetResponseable instead.
type NamespacesItemNamespaceResponseable interface {
    NamespacesItemNamespaceGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
