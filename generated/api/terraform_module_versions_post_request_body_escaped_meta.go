package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TerraformModuleVersionsPostRequestBody_meta struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The archiveId property
    archiveId *string
}
// NewTerraformModuleVersionsPostRequestBody_meta instantiates a new TerraformModuleVersionsPostRequestBody_meta and sets the default values.
func NewTerraformModuleVersionsPostRequestBody_meta()(*TerraformModuleVersionsPostRequestBody_meta) {
    m := &TerraformModuleVersionsPostRequestBody_meta{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTerraformModuleVersionsPostRequestBody_metaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTerraformModuleVersionsPostRequestBody_metaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTerraformModuleVersionsPostRequestBody_meta(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TerraformModuleVersionsPostRequestBody_meta) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchiveId gets the archive-id property value. The archiveId property
// returns a *string when successful
func (m *TerraformModuleVersionsPostRequestBody_meta) GetArchiveId()(*string) {
    return m.archiveId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TerraformModuleVersionsPostRequestBody_meta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archive-id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TerraformModuleVersionsPostRequestBody_meta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archive-id", m.GetArchiveId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TerraformModuleVersionsPostRequestBody_meta) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchiveId sets the archive-id property value. The archiveId property
func (m *TerraformModuleVersionsPostRequestBody_meta) SetArchiveId(value *string)() {
    m.archiveId = value
}
type TerraformModuleVersionsPostRequestBody_metaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveId()(*string)
    SetArchiveId(value *string)()
}
