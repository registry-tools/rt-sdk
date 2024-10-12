package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TerraformModuleVersionsPostRequestBody_module struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The namespace property
    namespace *string
    // The system property
    system *string
    // The version property
    version *string
}
// NewTerraformModuleVersionsPostRequestBody_module instantiates a new TerraformModuleVersionsPostRequestBody_module and sets the default values.
func NewTerraformModuleVersionsPostRequestBody_module()(*TerraformModuleVersionsPostRequestBody_module) {
    m := &TerraformModuleVersionsPostRequestBody_module{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTerraformModuleVersionsPostRequestBody_moduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTerraformModuleVersionsPostRequestBody_moduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTerraformModuleVersionsPostRequestBody_module(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TerraformModuleVersionsPostRequestBody_module) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TerraformModuleVersionsPostRequestBody_module) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["namespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespace(val)
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TerraformModuleVersionsPostRequestBody_module) GetName()(*string) {
    return m.name
}
// GetNamespace gets the namespace property value. The namespace property
// returns a *string when successful
func (m *TerraformModuleVersionsPostRequestBody_module) GetNamespace()(*string) {
    return m.namespace
}
// GetSystem gets the system property value. The system property
// returns a *string when successful
func (m *TerraformModuleVersionsPostRequestBody_module) GetSystem()(*string) {
    return m.system
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *TerraformModuleVersionsPostRequestBody_module) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *TerraformModuleVersionsPostRequestBody_module) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("namespace", m.GetNamespace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *TerraformModuleVersionsPostRequestBody_module) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *TerraformModuleVersionsPostRequestBody_module) SetName(value *string)() {
    m.name = value
}
// SetNamespace sets the namespace property value. The namespace property
func (m *TerraformModuleVersionsPostRequestBody_module) SetNamespace(value *string)() {
    m.namespace = value
}
// SetSystem sets the system property value. The system property
func (m *TerraformModuleVersionsPostRequestBody_module) SetSystem(value *string)() {
    m.system = value
}
// SetVersion sets the version property value. The version property
func (m *TerraformModuleVersionsPostRequestBody_module) SetVersion(value *string)() {
    m.version = value
}
type TerraformModuleVersionsPostRequestBody_moduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNamespace()(*string)
    GetSystem()(*string)
    GetVersion()(*string)
    SetName(value *string)()
    SetNamespace(value *string)()
    SetSystem(value *string)()
    SetVersion(value *string)()
}
