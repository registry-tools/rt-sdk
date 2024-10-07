package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The expiresAfter property
    expiresAfter *string
}
// NewServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken instantiates a new ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken and sets the default values.
func NewServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken()(*ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) {
    m := &ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) GetDescription()(*string) {
    return m.description
}
// GetExpiresAfter gets the expires-after property value. The expiresAfter property
// returns a *string when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) GetExpiresAfter()(*string) {
    return m.expiresAfter
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["expires-after"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAfter(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("expires-after", m.GetExpiresAfter())
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
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) SetDescription(value *string)() {
    m.description = value
}
// SetExpiresAfter sets the expires-after property value. The expiresAfter property
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationToken) SetExpiresAfter(value *string)() {
    m.expiresAfter = value
}
type ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetExpiresAfter()(*string)
    SetDescription(value *string)()
    SetExpiresAfter(value *string)()
}
