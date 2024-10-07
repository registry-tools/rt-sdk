package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AuthenticationTokensItemTokenGetResponseable instead.
type AuthenticationTokensItemTokenResponse struct {
    AuthenticationTokensItemTokenGetResponse
}
// NewAuthenticationTokensItemTokenResponse instantiates a new AuthenticationTokensItemTokenResponse and sets the default values.
func NewAuthenticationTokensItemTokenResponse()(*AuthenticationTokensItemTokenResponse) {
    m := &AuthenticationTokensItemTokenResponse{
        AuthenticationTokensItemTokenGetResponse: *NewAuthenticationTokensItemTokenGetResponse(),
    }
    return m
}
// CreateAuthenticationTokensItemTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationTokensItemTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationTokensItemTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use AuthenticationTokensItemTokenGetResponseable instead.
type AuthenticationTokensItemTokenResponseable interface {
    AuthenticationTokensItemTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
