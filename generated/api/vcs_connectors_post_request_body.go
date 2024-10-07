package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VcsConnectorsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The githubConnector property
    githubConnector VcsConnectorsPostRequestBody_githubConnectorable
}
// NewVcsConnectorsPostRequestBody instantiates a new VcsConnectorsPostRequestBody and sets the default values.
func NewVcsConnectorsPostRequestBody()(*VcsConnectorsPostRequestBody) {
    m := &VcsConnectorsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVcsConnectorsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVcsConnectorsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVcsConnectorsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VcsConnectorsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VcsConnectorsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["github-connector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVcsConnectorsPostRequestBody_githubConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubConnector(val.(VcsConnectorsPostRequestBody_githubConnectorable))
        }
        return nil
    }
    return res
}
// GetGithubConnector gets the github-connector property value. The githubConnector property
// returns a VcsConnectorsPostRequestBody_githubConnectorable when successful
func (m *VcsConnectorsPostRequestBody) GetGithubConnector()(VcsConnectorsPostRequestBody_githubConnectorable) {
    return m.githubConnector
}
// Serialize serializes information the current object
func (m *VcsConnectorsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("github-connector", m.GetGithubConnector())
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
func (m *VcsConnectorsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGithubConnector sets the github-connector property value. The githubConnector property
func (m *VcsConnectorsPostRequestBody) SetGithubConnector(value VcsConnectorsPostRequestBody_githubConnectorable)() {
    m.githubConnector = value
}
type VcsConnectorsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGithubConnector()(VcsConnectorsPostRequestBody_githubConnectorable)
    SetGithubConnector(value VcsConnectorsPostRequestBody_githubConnectorable)()
}
