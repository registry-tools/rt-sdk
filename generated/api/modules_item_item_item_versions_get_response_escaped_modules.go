package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModulesItemItemItemVersionsGetResponse_modules struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The versions property
    versions []ModulesItemItemItemVersionsGetResponse_modules_versionsable
}
// NewModulesItemItemItemVersionsGetResponse_modules instantiates a new ModulesItemItemItemVersionsGetResponse_modules and sets the default values.
func NewModulesItemItemItemVersionsGetResponse_modules()(*ModulesItemItemItemVersionsGetResponse_modules) {
    m := &ModulesItemItemItemVersionsGetResponse_modules{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateModulesItemItemItemVersionsGetResponse_modulesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModulesItemItemItemVersionsGetResponse_modulesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModulesItemItemItemVersionsGetResponse_modules(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ModulesItemItemItemVersionsGetResponse_modules) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ModulesItemItemItemVersionsGetResponse_modules) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateModulesItemItemItemVersionsGetResponse_modules_versionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ModulesItemItemItemVersionsGetResponse_modules_versionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ModulesItemItemItemVersionsGetResponse_modules_versionsable)
                }
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetVersions gets the versions property value. The versions property
// returns a []ModulesItemItemItemVersionsGetResponse_modules_versionsable when successful
func (m *ModulesItemItemItemVersionsGetResponse_modules) GetVersions()([]ModulesItemItemItemVersionsGetResponse_modules_versionsable) {
    return m.versions
}
// Serialize serializes information the current object
func (m *ModulesItemItemItemVersionsGetResponse_modules) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("versions", cast)
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
func (m *ModulesItemItemItemVersionsGetResponse_modules) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVersions sets the versions property value. The versions property
func (m *ModulesItemItemItemVersionsGetResponse_modules) SetVersions(value []ModulesItemItemItemVersionsGetResponse_modules_versionsable)() {
    m.versions = value
}
type ModulesItemItemItemVersionsGetResponse_modulesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVersions()([]ModulesItemItemItemVersionsGetResponse_modules_versionsable)
    SetVersions(value []ModulesItemItemItemVersionsGetResponse_modules_versionsable)()
}
