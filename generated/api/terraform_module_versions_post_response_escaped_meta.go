package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TerraformModuleVersionsPostResponse_meta struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The downloadUrl property
    downloadUrl *string
}
// NewTerraformModuleVersionsPostResponse_meta instantiates a new TerraformModuleVersionsPostResponse_meta and sets the default values.
func NewTerraformModuleVersionsPostResponse_meta()(*TerraformModuleVersionsPostResponse_meta) {
    m := &TerraformModuleVersionsPostResponse_meta{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTerraformModuleVersionsPostResponse_metaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTerraformModuleVersionsPostResponse_metaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTerraformModuleVersionsPostResponse_meta(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TerraformModuleVersionsPostResponse_meta) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDownloadUrl gets the download-url property value. The downloadUrl property
// returns a *string when successful
func (m *TerraformModuleVersionsPostResponse_meta) GetDownloadUrl()(*string) {
    return m.downloadUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TerraformModuleVersionsPostResponse_meta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["download-url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TerraformModuleVersionsPostResponse_meta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("download-url", m.GetDownloadUrl())
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
func (m *TerraformModuleVersionsPostResponse_meta) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDownloadUrl sets the download-url property value. The downloadUrl property
func (m *TerraformModuleVersionsPostResponse_meta) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
type TerraformModuleVersionsPostResponse_metaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDownloadUrl()(*string)
    SetDownloadUrl(value *string)()
}
