package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type NamespacesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The namespace property
    namespace ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable
}
// NewNamespacesPostRequestBody instantiates a new NamespacesPostRequestBody and sets the default values.
func NewNamespacesPostRequestBody()(*NamespacesPostRequestBody) {
    m := &NamespacesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["namespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateNamespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespace(val.(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable))
        }
        return nil
    }
    return res
}
// GetNamespace gets the namespace property value. The namespace property
// returns a Namespaceable when successful
func (m *NamespacesPostRequestBody) GetNamespace()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable) {
    return m.namespace
}
// Serialize serializes information the current object
func (m *NamespacesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("namespace", m.GetNamespace())
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
func (m *NamespacesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNamespace sets the namespace property value. The namespace property
func (m *NamespacesPostRequestBody) SetNamespace(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable)() {
    m.namespace = value
}
type NamespacesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNamespace()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable)
    SetNamespace(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable)()
}
