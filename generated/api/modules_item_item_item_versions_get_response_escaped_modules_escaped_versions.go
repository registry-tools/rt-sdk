package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModulesItemItemItemVersionsGetResponse_modules_versions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The version property
    version *string
}
// NewModulesItemItemItemVersionsGetResponse_modules_versions instantiates a new ModulesItemItemItemVersionsGetResponse_modules_versions and sets the default values.
func NewModulesItemItemItemVersionsGetResponse_modules_versions()(*ModulesItemItemItemVersionsGetResponse_modules_versions) {
    m := &ModulesItemItemItemVersionsGetResponse_modules_versions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateModulesItemItemItemVersionsGetResponse_modules_versionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModulesItemItemItemVersionsGetResponse_modules_versionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModulesItemItemItemVersionsGetResponse_modules_versions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ModulesItemItemItemVersionsGetResponse_modules_versions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ModulesItemItemItemVersionsGetResponse_modules_versions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *ModulesItemItemItemVersionsGetResponse_modules_versions) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ModulesItemItemItemVersionsGetResponse_modules_versions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ModulesItemItemItemVersionsGetResponse_modules_versions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVersion sets the version property value. The version property
func (m *ModulesItemItemItemVersionsGetResponse_modules_versions) SetVersion(value *string)() {
    m.version = value
}
type ModulesItemItemItemVersionsGetResponse_modules_versionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVersion()(*string)
    SetVersion(value *string)()
}
