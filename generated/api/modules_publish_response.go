package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ModulesPublishPostResponseable instead.
type ModulesPublishResponse struct {
    ModulesPublishPostResponse
}
// NewModulesPublishResponse instantiates a new ModulesPublishResponse and sets the default values.
func NewModulesPublishResponse()(*ModulesPublishResponse) {
    m := &ModulesPublishResponse{
        ModulesPublishPostResponse: *NewModulesPublishPostResponse(),
    }
    return m
}
// CreateModulesPublishResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModulesPublishResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModulesPublishResponse(), nil
}
// Deprecated: This class is obsolete. Use ModulesPublishPostResponseable instead.
type ModulesPublishResponseable interface {
    ModulesPublishPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
