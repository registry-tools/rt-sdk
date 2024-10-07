package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type ServiceAccountsItemAuthenticationTokensGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data []ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable
    // The meta property
    meta ServiceAccountsItemAuthenticationTokensGetResponse_metaable
}
// NewServiceAccountsItemAuthenticationTokensGetResponse instantiates a new ServiceAccountsItemAuthenticationTokensGetResponse and sets the default values.
func NewServiceAccountsItemAuthenticationTokensGetResponse()(*ServiceAccountsItemAuthenticationTokensGetResponse) {
    m := &ServiceAccountsItemAuthenticationTokensGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceAccountsItemAuthenticationTokensGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceAccountsItemAuthenticationTokensGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAccountsItemAuthenticationTokensGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
// returns a []AuthenticationTokenable when successful
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) GetData()([]ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateAuthenticationTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)
                }
            }
            m.SetData(res)
        }
        return nil
    }
    res["meta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceAccountsItemAuthenticationTokensGetResponse_metaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeta(val.(ServiceAccountsItemAuthenticationTokensGetResponse_metaable))
        }
        return nil
    }
    return res
}
// GetMeta gets the meta property value. The meta property
// returns a ServiceAccountsItemAuthenticationTokensGetResponse_metaable when successful
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) GetMeta()(ServiceAccountsItemAuthenticationTokensGetResponse_metaable) {
    return m.meta
}
// Serialize serializes information the current object
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("meta", m.GetMeta())
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
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) SetData(value []ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)() {
    m.data = value
}
// SetMeta sets the meta property value. The meta property
func (m *ServiceAccountsItemAuthenticationTokensGetResponse) SetMeta(value ServiceAccountsItemAuthenticationTokensGetResponse_metaable)() {
    m.meta = value
}
type ServiceAccountsItemAuthenticationTokensGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()([]ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)
    GetMeta()(ServiceAccountsItemAuthenticationTokensGetResponse_metaable)
    SetData(value []ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)()
    SetMeta(value ServiceAccountsItemAuthenticationTokensGetResponse_metaable)()
}
