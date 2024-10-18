package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthenticationToken struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The expiresAfter property
    expiresAfter *string
    // The expiresAt property
    expiresAt *string
    // The id property
    id *string
    // The lastUsedAround property
    lastUsedAround *string
    // The token property
    token *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAuthenticationToken instantiates a new AuthenticationToken and sets the default values.
func NewAuthenticationToken()(*AuthenticationToken) {
    m := &AuthenticationToken{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthenticationTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationToken(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuthenticationToken) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created-at property value. The createdAt property
// returns a *Time when successful
func (m *AuthenticationToken) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *AuthenticationToken) GetDescription()(*string) {
    return m.description
}
// GetExpiresAfter gets the expires-after property value. The expiresAfter property
// returns a *string when successful
func (m *AuthenticationToken) GetExpiresAfter()(*string) {
    return m.expiresAfter
}
// GetExpiresAt gets the expires-at property value. The expiresAt property
// returns a *string when successful
func (m *AuthenticationToken) GetExpiresAt()(*string) {
    return m.expiresAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthenticationToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created-at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
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
    res["expires-at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["last-used-around"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsedAround(val)
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    res["updated-at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *AuthenticationToken) GetId()(*string) {
    return m.id
}
// GetLastUsedAround gets the last-used-around property value. The lastUsedAround property
// returns a *string when successful
func (m *AuthenticationToken) GetLastUsedAround()(*string) {
    return m.lastUsedAround
}
// GetToken gets the token property value. The token property
// returns a *string when successful
func (m *AuthenticationToken) GetToken()(*string) {
    return m.token
}
// GetUpdatedAt gets the updated-at property value. The updatedAt property
// returns a *Time when successful
func (m *AuthenticationToken) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *AuthenticationToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("token", m.GetToken())
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
func (m *AuthenticationToken) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created-at property value. The createdAt property
func (m *AuthenticationToken) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetDescription sets the description property value. The description property
func (m *AuthenticationToken) SetDescription(value *string)() {
    m.description = value
}
// SetExpiresAfter sets the expires-after property value. The expiresAfter property
func (m *AuthenticationToken) SetExpiresAfter(value *string)() {
    m.expiresAfter = value
}
// SetExpiresAt sets the expires-at property value. The expiresAt property
func (m *AuthenticationToken) SetExpiresAt(value *string)() {
    m.expiresAt = value
}
// SetId sets the id property value. The id property
func (m *AuthenticationToken) SetId(value *string)() {
    m.id = value
}
// SetLastUsedAround sets the last-used-around property value. The lastUsedAround property
func (m *AuthenticationToken) SetLastUsedAround(value *string)() {
    m.lastUsedAround = value
}
// SetToken sets the token property value. The token property
func (m *AuthenticationToken) SetToken(value *string)() {
    m.token = value
}
// SetUpdatedAt sets the updated-at property value. The updatedAt property
func (m *AuthenticationToken) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type AuthenticationTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetExpiresAfter()(*string)
    GetExpiresAt()(*string)
    GetId()(*string)
    GetLastUsedAround()(*string)
    GetToken()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetExpiresAfter(value *string)()
    SetExpiresAt(value *string)()
    SetId(value *string)()
    SetLastUsedAround(value *string)()
    SetToken(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
