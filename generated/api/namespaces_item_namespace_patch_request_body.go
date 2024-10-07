package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type NamespacesItemNamespacePatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The namespace property
    namespace ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable
}
// NewNamespacesItemNamespacePatchRequestBody instantiates a new NamespacesItemNamespacePatchRequestBody and sets the default values.
func NewNamespacesItemNamespacePatchRequestBody()(*NamespacesItemNamespacePatchRequestBody) {
    m := &NamespacesItemNamespacePatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesItemNamespacePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemNamespacePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemNamespacePatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesItemNamespacePatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesItemNamespacePatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *NamespacesItemNamespacePatchRequestBody) GetNamespace()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable) {
    return m.namespace
}
// Serialize serializes information the current object
func (m *NamespacesItemNamespacePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *NamespacesItemNamespacePatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNamespace sets the namespace property value. The namespace property
func (m *NamespacesItemNamespacePatchRequestBody) SetNamespace(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable)() {
    m.namespace = value
}
type NamespacesItemNamespacePatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNamespace()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable)
    SetNamespace(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.Namespaceable)()
}
