package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NamespacesItemServiceAccountsGetResponseable instead.
type NamespacesItemServiceAccountsResponse struct {
    NamespacesItemServiceAccountsGetResponse
}
// NewNamespacesItemServiceAccountsResponse instantiates a new NamespacesItemServiceAccountsResponse and sets the default values.
func NewNamespacesItemServiceAccountsResponse()(*NamespacesItemServiceAccountsResponse) {
    m := &NamespacesItemServiceAccountsResponse{
        NamespacesItemServiceAccountsGetResponse: *NewNamespacesItemServiceAccountsGetResponse(),
    }
    return m
}
// CreateNamespacesItemServiceAccountsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemServiceAccountsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemServiceAccountsResponse(), nil
}
// Deprecated: This class is obsolete. Use NamespacesItemServiceAccountsGetResponseable instead.
type NamespacesItemServiceAccountsResponseable interface {
    NamespacesItemServiceAccountsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
