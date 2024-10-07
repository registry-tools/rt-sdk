package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type NamespacesItemServiceAccountsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data []ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable
}
// NewNamespacesItemServiceAccountsGetResponse instantiates a new NamespacesItemServiceAccountsGetResponse and sets the default values.
func NewNamespacesItemServiceAccountsGetResponse()(*NamespacesItemServiceAccountsGetResponse) {
    m := &NamespacesItemServiceAccountsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesItemServiceAccountsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemServiceAccountsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemServiceAccountsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesItemServiceAccountsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
// returns a []ServiceAccountable when successful
func (m *NamespacesItemServiceAccountsGetResponse) GetData()([]ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesItemServiceAccountsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateServiceAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)
                }
            }
            m.SetData(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *NamespacesItemServiceAccountsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetData()))
        for i, v := range m.GetData() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("data", cast)
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
func (m *NamespacesItemServiceAccountsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *NamespacesItemServiceAccountsGetResponse) SetData(value []ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)() {
    m.data = value
}
type NamespacesItemServiceAccountsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()([]ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)
    SetData(value []ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)()
}
