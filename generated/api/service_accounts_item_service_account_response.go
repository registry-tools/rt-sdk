package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceAccountsItemServiceAccountGetResponseable instead.
type ServiceAccountsItemServiceAccountResponse struct {
    ServiceAccountsItemServiceAccountGetResponse
}
// NewServiceAccountsItemServiceAccountResponse instantiates a new ServiceAccountsItemServiceAccountResponse and sets the default values.
func NewServiceAccountsItemServiceAccountResponse()(*ServiceAccountsItemServiceAccountResponse) {
    m := &ServiceAccountsItemServiceAccountResponse{
        ServiceAccountsItemServiceAccountGetResponse: *NewServiceAccountsItemServiceAccountGetResponse(),
    }
    return m
}
// CreateServiceAccountsItemServiceAccountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceAccountsItemServiceAccountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAccountsItemServiceAccountResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceAccountsItemServiceAccountGetResponseable instead.
type ServiceAccountsItemServiceAccountResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAccountsItemServiceAccountGetResponseable
}
