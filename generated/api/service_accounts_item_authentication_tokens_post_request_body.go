package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceAccountsItemAuthenticationTokensPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The authenticationToken property
    authenticationToken ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable
}
// NewServiceAccountsItemAuthenticationTokensPostRequestBody instantiates a new ServiceAccountsItemAuthenticationTokensPostRequestBody and sets the default values.
func NewServiceAccountsItemAuthenticationTokensPostRequestBody()(*ServiceAccountsItemAuthenticationTokensPostRequestBody) {
    m := &ServiceAccountsItemAuthenticationTokensPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceAccountsItemAuthenticationTokensPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceAccountsItemAuthenticationTokensPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAccountsItemAuthenticationTokensPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthenticationToken gets the authentication-token property value. The authenticationToken property
// returns a ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody) GetAuthenticationToken()(ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable) {
    return m.authenticationToken
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authentication-token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationToken(val.(ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authentication-token", m.GetAuthenticationToken())
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
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthenticationToken sets the authentication-token property value. The authenticationToken property
func (m *ServiceAccountsItemAuthenticationTokensPostRequestBody) SetAuthenticationToken(value ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable)() {
    m.authenticationToken = value
}
type ServiceAccountsItemAuthenticationTokensPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationToken()(ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable)
    SetAuthenticationToken(value ServiceAccountsItemAuthenticationTokensPostRequestBody_authenticationTokenable)()
}
