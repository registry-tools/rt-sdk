package auth

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TokenPostResponse struct {
    // The access_token property
    access_token *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Time in seconds until the token expires
    expires_in *int32
}
// NewTokenPostResponse instantiates a new TokenPostResponse and sets the default values.
func NewTokenPostResponse()(*TokenPostResponse) {
    m := &TokenPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTokenPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTokenPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTokenPostResponse(), nil
}
// GetAccessToken gets the access_token property value. The access_token property
// returns a *string when successful
func (m *TokenPostResponse) GetAccessToken()(*string) {
    return m.access_token
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TokenPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresIn gets the expires_in property value. Time in seconds until the token expires
// returns a *int32 when successful
func (m *TokenPostResponse) GetExpiresIn()(*int32) {
    return m.expires_in
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TokenPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessToken(val)
        }
        return nil
    }
    res["expires_in"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresIn(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TokenPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("access_token", m.GetAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("expires_in", m.GetExpiresIn())
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
// SetAccessToken sets the access_token property value. The access_token property
func (m *TokenPostResponse) SetAccessToken(value *string)() {
    m.access_token = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TokenPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresIn sets the expires_in property value. Time in seconds until the token expires
func (m *TokenPostResponse) SetExpiresIn(value *int32)() {
    m.expires_in = value
}
type TokenPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessToken()(*string)
    GetExpiresIn()(*int32)
    SetAccessToken(value *string)()
    SetExpiresIn(value *int32)()
}
