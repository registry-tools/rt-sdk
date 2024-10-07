package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type NamespacesItemServiceAccountsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The serviceAccount property
    serviceAccount ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable
}
// NewNamespacesItemServiceAccountsPostRequestBody instantiates a new NamespacesItemServiceAccountsPostRequestBody and sets the default values.
func NewNamespacesItemServiceAccountsPostRequestBody()(*NamespacesItemServiceAccountsPostRequestBody) {
    m := &NamespacesItemServiceAccountsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesItemServiceAccountsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemServiceAccountsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemServiceAccountsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesItemServiceAccountsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesItemServiceAccountsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["service-account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateServiceAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccount(val.(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable))
        }
        return nil
    }
    return res
}
// GetServiceAccount gets the service-account property value. The serviceAccount property
// returns a ServiceAccountable when successful
func (m *NamespacesItemServiceAccountsPostRequestBody) GetServiceAccount()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable) {
    return m.serviceAccount
}
// Serialize serializes information the current object
func (m *NamespacesItemServiceAccountsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("service-account", m.GetServiceAccount())
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
func (m *NamespacesItemServiceAccountsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetServiceAccount sets the service-account property value. The serviceAccount property
func (m *NamespacesItemServiceAccountsPostRequestBody) SetServiceAccount(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)() {
    m.serviceAccount = value
}
type NamespacesItemServiceAccountsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetServiceAccount()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)
    SetServiceAccount(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.ServiceAccountable)()
}
