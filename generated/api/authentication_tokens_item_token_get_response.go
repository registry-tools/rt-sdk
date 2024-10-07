package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type AuthenticationTokensItemTokenGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable
}
// NewAuthenticationTokensItemTokenGetResponse instantiates a new AuthenticationTokensItemTokenGetResponse and sets the default values.
func NewAuthenticationTokensItemTokenGetResponse()(*AuthenticationTokensItemTokenGetResponse) {
    m := &AuthenticationTokensItemTokenGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationTokensItemTokenGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationTokensItemTokenGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationTokensItemTokenGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationTokensItemTokenGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
// returns a AuthenticationTokenable when successful
func (m *AuthenticationTokensItemTokenGetResponse) GetData()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationTokensItemTokenGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateAuthenticationTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationTokensItemTokenGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
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
func (m *AuthenticationTokensItemTokenGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *AuthenticationTokensItemTokenGetResponse) SetData(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)() {
    m.data = value
}
type AuthenticationTokensItemTokenGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)
    SetData(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.AuthenticationTokenable)()
}
