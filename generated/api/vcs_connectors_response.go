package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use VcsConnectorsPostResponseable instead.
type VcsConnectorsResponse struct {
    VcsConnectorsPostResponse
}
// NewVcsConnectorsResponse instantiates a new VcsConnectorsResponse and sets the default values.
func NewVcsConnectorsResponse()(*VcsConnectorsResponse) {
    m := &VcsConnectorsResponse{
        VcsConnectorsPostResponse: *NewVcsConnectorsPostResponse(),
    }
    return m
}
// CreateVcsConnectorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVcsConnectorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVcsConnectorsResponse(), nil
}
// Deprecated: This class is obsolete. Use VcsConnectorsPostResponseable instead.
type VcsConnectorsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VcsConnectorsPostResponseable
}
