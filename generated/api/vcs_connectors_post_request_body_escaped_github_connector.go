package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VcsConnectorsPostRequestBody_githubConnector struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The token property
    token *string
}
// NewVcsConnectorsPostRequestBody_githubConnector instantiates a new VcsConnectorsPostRequestBody_githubConnector and sets the default values.
func NewVcsConnectorsPostRequestBody_githubConnector()(*VcsConnectorsPostRequestBody_githubConnector) {
    m := &VcsConnectorsPostRequestBody_githubConnector{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVcsConnectorsPostRequestBody_githubConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVcsConnectorsPostRequestBody_githubConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVcsConnectorsPostRequestBody_githubConnector(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VcsConnectorsPostRequestBody_githubConnector) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *VcsConnectorsPostRequestBody_githubConnector) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VcsConnectorsPostRequestBody_githubConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetToken gets the token property value. The token property
// returns a *string when successful
func (m *VcsConnectorsPostRequestBody_githubConnector) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *VcsConnectorsPostRequestBody_githubConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
func (m *VcsConnectorsPostRequestBody_githubConnector) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *VcsConnectorsPostRequestBody_githubConnector) SetDescription(value *string)() {
    m.description = value
}
// SetToken sets the token property value. The token property
func (m *VcsConnectorsPostRequestBody_githubConnector) SetToken(value *string)() {
    m.token = value
}
type VcsConnectorsPostRequestBody_githubConnectorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetToken()(*string)
    SetDescription(value *string)()
    SetToken(value *string)()
}
