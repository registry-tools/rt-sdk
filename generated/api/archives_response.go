package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ArchivesPostResponseable instead.
type ArchivesResponse struct {
    ArchivesPostResponse
}
// NewArchivesResponse instantiates a new ArchivesResponse and sets the default values.
func NewArchivesResponse()(*ArchivesResponse) {
    m := &ArchivesResponse{
        ArchivesPostResponse: *NewArchivesPostResponse(),
    }
    return m
}
// CreateArchivesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateArchivesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewArchivesResponse(), nil
}
// Deprecated: This class is obsolete. Use ArchivesPostResponseable instead.
type ArchivesResponseable interface {
    ArchivesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
