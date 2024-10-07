package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceAccountsItemAuthenticationTokensGetResponseable instead.
type ServiceAccountsItemAuthenticationTokensResponse struct {
    ServiceAccountsItemAuthenticationTokensGetResponse
}
// NewServiceAccountsItemAuthenticationTokensResponse instantiates a new ServiceAccountsItemAuthenticationTokensResponse and sets the default values.
func NewServiceAccountsItemAuthenticationTokensResponse()(*ServiceAccountsItemAuthenticationTokensResponse) {
    m := &ServiceAccountsItemAuthenticationTokensResponse{
        ServiceAccountsItemAuthenticationTokensGetResponse: *NewServiceAccountsItemAuthenticationTokensGetResponse(),
    }
    return m
}
// CreateServiceAccountsItemAuthenticationTokensResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceAccountsItemAuthenticationTokensResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAccountsItemAuthenticationTokensResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceAccountsItemAuthenticationTokensGetResponseable instead.
type ServiceAccountsItemAuthenticationTokensResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAccountsItemAuthenticationTokensGetResponseable
}
