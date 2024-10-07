package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NamespacesItemTagPublishersPostRequestBody_tagPublisher struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repo property
    repo *string
    // The vcsConnectorId property
    vcsConnectorId *string
}
// NewNamespacesItemTagPublishersPostRequestBody_tagPublisher instantiates a new NamespacesItemTagPublishersPostRequestBody_tagPublisher and sets the default values.
func NewNamespacesItemTagPublishersPostRequestBody_tagPublisher()(*NamespacesItemTagPublishersPostRequestBody_tagPublisher) {
    m := &NamespacesItemTagPublishersPostRequestBody_tagPublisher{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesItemTagPublishersPostRequestBody_tagPublisherFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemTagPublishersPostRequestBody_tagPublisherFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemTagPublishersPostRequestBody_tagPublisher(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepo(val)
        }
        return nil
    }
    res["vcs-connector-id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVcsConnectorId(val)
        }
        return nil
    }
    return res
}
// GetRepo gets the repo property value. The repo property
// returns a *string when successful
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) GetRepo()(*string) {
    return m.repo
}
// GetVcsConnectorId gets the vcs-connector-id property value. The vcsConnectorId property
// returns a *string when successful
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) GetVcsConnectorId()(*string) {
    return m.vcsConnectorId
}
// Serialize serializes information the current object
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("repo", m.GetRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vcs-connector-id", m.GetVcsConnectorId())
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
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepo sets the repo property value. The repo property
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) SetRepo(value *string)() {
    m.repo = value
}
// SetVcsConnectorId sets the vcs-connector-id property value. The vcsConnectorId property
func (m *NamespacesItemTagPublishersPostRequestBody_tagPublisher) SetVcsConnectorId(value *string)() {
    m.vcsConnectorId = value
}
type NamespacesItemTagPublishersPostRequestBody_tagPublisherable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepo()(*string)
    GetVcsConnectorId()(*string)
    SetRepo(value *string)()
    SetVcsConnectorId(value *string)()
}
