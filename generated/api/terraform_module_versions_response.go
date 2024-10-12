package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TerraformModuleVersionsPostResponseable instead.
type TerraformModuleVersionsResponse struct {
    TerraformModuleVersionsPostResponse
}
// NewTerraformModuleVersionsResponse instantiates a new TerraformModuleVersionsResponse and sets the default values.
func NewTerraformModuleVersionsResponse()(*TerraformModuleVersionsResponse) {
    m := &TerraformModuleVersionsResponse{
        TerraformModuleVersionsPostResponse: *NewTerraformModuleVersionsPostResponse(),
    }
    return m
}
// CreateTerraformModuleVersionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTerraformModuleVersionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTerraformModuleVersionsResponse(), nil
}
// Deprecated: This class is obsolete. Use TerraformModuleVersionsPostResponseable instead.
type TerraformModuleVersionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TerraformModuleVersionsPostResponseable
}
