package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ModulesItemItemItemVersionsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The modules property
    modules []ModulesItemItemItemVersionsGetResponse_modulesable
}
// NewModulesItemItemItemVersionsGetResponse instantiates a new ModulesItemItemItemVersionsGetResponse and sets the default values.
func NewModulesItemItemItemVersionsGetResponse()(*ModulesItemItemItemVersionsGetResponse) {
    m := &ModulesItemItemItemVersionsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateModulesItemItemItemVersionsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModulesItemItemItemVersionsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModulesItemItemItemVersionsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ModulesItemItemItemVersionsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ModulesItemItemItemVersionsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["modules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateModulesItemItemItemVersionsGetResponse_modulesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ModulesItemItemItemVersionsGetResponse_modulesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ModulesItemItemItemVersionsGetResponse_modulesable)
                }
            }
            m.SetModules(res)
        }
        return nil
    }
    return res
}
// GetModules gets the modules property value. The modules property
// returns a []ModulesItemItemItemVersionsGetResponse_modulesable when successful
func (m *ModulesItemItemItemVersionsGetResponse) GetModules()([]ModulesItemItemItemVersionsGetResponse_modulesable) {
    return m.modules
}
// Serialize serializes information the current object
func (m *ModulesItemItemItemVersionsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetModules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModules()))
        for i, v := range m.GetModules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("modules", cast)
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
func (m *ModulesItemItemItemVersionsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetModules sets the modules property value. The modules property
func (m *ModulesItemItemItemVersionsGetResponse) SetModules(value []ModulesItemItemItemVersionsGetResponse_modulesable)() {
    m.modules = value
}
type ModulesItemItemItemVersionsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetModules()([]ModulesItemItemItemVersionsGetResponse_modulesable)
    SetModules(value []ModulesItemItemItemVersionsGetResponse_modulesable)()
}
