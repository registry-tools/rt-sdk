package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TerraformModuleVersionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The meta property
    meta TerraformModuleVersionsPostRequestBody_metaable
    // The module property
    module TerraformModuleVersionsPostRequestBody_moduleable
}
// NewTerraformModuleVersionsPostRequestBody instantiates a new TerraformModuleVersionsPostRequestBody and sets the default values.
func NewTerraformModuleVersionsPostRequestBody()(*TerraformModuleVersionsPostRequestBody) {
    m := &TerraformModuleVersionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTerraformModuleVersionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTerraformModuleVersionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTerraformModuleVersionsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TerraformModuleVersionsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TerraformModuleVersionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["meta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTerraformModuleVersionsPostRequestBody_metaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(TerraformModuleVersionsPostRequestBody_metaable))
        }
        return nil
    }
    res["module"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTerraformModuleVersionsPostRequestBody_moduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModule(val.(TerraformModuleVersionsPostRequestBody_moduleable))
        }
        return nil
    }
    return res
}
// GetMeta gets the meta property value. The meta property
// returns a TerraformModuleVersionsPostRequestBody_metaable when successful
func (m *TerraformModuleVersionsPostRequestBody) GetMeta()(TerraformModuleVersionsPostRequestBody_metaable) {
    return m.meta
}
// GetModule gets the module property value. The module property
// returns a TerraformModuleVersionsPostRequestBody_moduleable when successful
func (m *TerraformModuleVersionsPostRequestBody) GetModule()(TerraformModuleVersionsPostRequestBody_moduleable) {
    return m.module
}
// Serialize serializes information the current object
func (m *TerraformModuleVersionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("meta", m.GetMeta())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("module", m.GetModule())
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
func (m *TerraformModuleVersionsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMeta sets the meta property value. The meta property
func (m *TerraformModuleVersionsPostRequestBody) SetMeta(value TerraformModuleVersionsPostRequestBody_metaable)() {
    m.meta = value
}
// SetModule sets the module property value. The module property
func (m *TerraformModuleVersionsPostRequestBody) SetModule(value TerraformModuleVersionsPostRequestBody_moduleable)() {
    m.module = value
}
type TerraformModuleVersionsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMeta()(TerraformModuleVersionsPostRequestBody_metaable)
    GetModule()(TerraformModuleVersionsPostRequestBody_moduleable)
    SetMeta(value TerraformModuleVersionsPostRequestBody_metaable)()
    SetModule(value TerraformModuleVersionsPostRequestBody_moduleable)()
}
