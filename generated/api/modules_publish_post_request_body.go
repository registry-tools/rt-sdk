package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModulesPublishPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The module property
    module ModulesPublishPostRequestBody_moduleable
}
// NewModulesPublishPostRequestBody instantiates a new ModulesPublishPostRequestBody and sets the default values.
func NewModulesPublishPostRequestBody()(*ModulesPublishPostRequestBody) {
    m := &ModulesPublishPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateModulesPublishPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModulesPublishPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModulesPublishPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ModulesPublishPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ModulesPublishPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["module"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateModulesPublishPostRequestBody_moduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModule(val.(ModulesPublishPostRequestBody_moduleable))
        }
        return nil
    }
    return res
}
// GetModule gets the module property value. The module property
// returns a ModulesPublishPostRequestBody_moduleable when successful
func (m *ModulesPublishPostRequestBody) GetModule()(ModulesPublishPostRequestBody_moduleable) {
    return m.module
}
// Serialize serializes information the current object
func (m *ModulesPublishPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ModulesPublishPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetModule sets the module property value. The module property
func (m *ModulesPublishPostRequestBody) SetModule(value ModulesPublishPostRequestBody_moduleable)() {
    m.module = value
}
type ModulesPublishPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetModule()(ModulesPublishPostRequestBody_moduleable)
    SetModule(value ModulesPublishPostRequestBody_moduleable)()
}
