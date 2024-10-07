package auth

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TokenPostResponseable instead.
type TokenResponse struct {
    TokenPostResponse
}
// NewTokenResponse instantiates a new TokenResponse and sets the default values.
func NewTokenResponse()(*TokenResponse) {
    m := &TokenResponse{
        TokenPostResponse: *NewTokenPostResponse(),
    }
    return m
}
// CreateTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use TokenPostResponseable instead.
type TokenResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TokenPostResponseable
}
