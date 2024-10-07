package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModulesPublishPostRequestBody_module struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The archive property
    archive *string
    // The name property
    name *string
    // The namespace property
    namespace *string
    // The system property
    system *string
    // The version property
    version *string
}
// NewModulesPublishPostRequestBody_module instantiates a new ModulesPublishPostRequestBody_module and sets the default values.
func NewModulesPublishPostRequestBody_module()(*ModulesPublishPostRequestBody_module) {
    m := &ModulesPublishPostRequestBody_module{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateModulesPublishPostRequestBody_moduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModulesPublishPostRequestBody_moduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModulesPublishPostRequestBody_module(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ModulesPublishPostRequestBody_module) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchive gets the archive property value. The archive property
// returns a *string when successful
func (m *ModulesPublishPostRequestBody_module) GetArchive()(*string) {
    return m.archive
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ModulesPublishPostRequestBody_module) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchive(val)
        }
        return nil
    }
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
func (m *ModulesPublishPostRequestBody_module) GetName()(*string) {
    return m.name
}
// GetNamespace gets the namespace property value. The namespace property
// returns a *string when successful
func (m *ModulesPublishPostRequestBody_module) GetNamespace()(*string) {
    return m.namespace
}
// GetSystem gets the system property value. The system property
// returns a *string when successful
func (m *ModulesPublishPostRequestBody_module) GetSystem()(*string) {
    return m.system
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *ModulesPublishPostRequestBody_module) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ModulesPublishPostRequestBody_module) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archive", m.GetArchive())
        if err != nil {
            return err
        }
    }
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
func (m *ModulesPublishPostRequestBody_module) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchive sets the archive property value. The archive property
func (m *ModulesPublishPostRequestBody_module) SetArchive(value *string)() {
    m.archive = value
}
// SetName sets the name property value. The name property
func (m *ModulesPublishPostRequestBody_module) SetName(value *string)() {
    m.name = value
}
// SetNamespace sets the namespace property value. The namespace property
func (m *ModulesPublishPostRequestBody_module) SetNamespace(value *string)() {
    m.namespace = value
}
// SetSystem sets the system property value. The system property
func (m *ModulesPublishPostRequestBody_module) SetSystem(value *string)() {
    m.system = value
}
// SetVersion sets the version property value. The version property
func (m *ModulesPublishPostRequestBody_module) SetVersion(value *string)() {
    m.version = value
}
type ModulesPublishPostRequestBody_moduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchive()(*string)
    GetName()(*string)
    GetNamespace()(*string)
    GetSystem()(*string)
    GetVersion()(*string)
    SetArchive(value *string)()
    SetName(value *string)()
    SetNamespace(value *string)()
    SetSystem(value *string)()
    SetVersion(value *string)()
}
