package auth

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TokenPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The client ID associated with the organization or service account
    client_id *string
    // The client secret associated with the organization or service account
    client_secret *string
    // Should always be "authorization_code"
    grant_type *string
}
// NewTokenPostRequestBody instantiates a new TokenPostRequestBody and sets the default values.
func NewTokenPostRequestBody()(*TokenPostRequestBody) {
    m := &TokenPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTokenPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TokenPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the client_id property value. The client ID associated with the organization or service account
// returns a *string when successful
func (m *TokenPostRequestBody) GetClientId()(*string) {
    return m.client_id
}
// GetClientSecret gets the client_secret property value. The client secret associated with the organization or service account
// returns a *string when successful
func (m *TokenPostRequestBody) GetClientSecret()(*string) {
    return m.client_secret
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TokenPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["client_secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["grant_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantType(val)
        }
        return nil
    }
    return res
}
// GetGrantType gets the grant_type property value. Should always be "authorization_code"
// returns a *string when successful
func (m *TokenPostRequestBody) GetGrantType()(*string) {
    return m.grant_type
}
// Serialize serializes information the current object
func (m *TokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client_id", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("client_secret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("grant_type", m.GetGrantType())
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
func (m *TokenPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the client_id property value. The client ID associated with the organization or service account
func (m *TokenPostRequestBody) SetClientId(value *string)() {
    m.client_id = value
}
// SetClientSecret sets the client_secret property value. The client secret associated with the organization or service account
func (m *TokenPostRequestBody) SetClientSecret(value *string)() {
    m.client_secret = value
}
// SetGrantType sets the grant_type property value. Should always be "authorization_code"
func (m *TokenPostRequestBody) SetGrantType(value *string)() {
    m.grant_type = value
}
type TokenPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetGrantType()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetGrantType(value *string)()
}
