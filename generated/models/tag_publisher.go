package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagPublisher struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The backfillPattern property
    backfillPattern *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *string
    // The repo property
    repo *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The vcsConnectorId property
    vcsConnectorId *string
}
// NewTagPublisher instantiates a new TagPublisher and sets the default values.
func NewTagPublisher()(*TagPublisher) {
    m := &TagPublisher{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTagPublisherFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagPublisherFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTagPublisher(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagPublisher) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBackfillPattern gets the backfill-pattern property value. The backfillPattern property
// returns a *string when successful
func (m *TagPublisher) GetBackfillPattern()(*string) {
    return m.backfillPattern
}
// GetCreatedAt gets the created-at property value. The createdAt property
// returns a *Time when successful
func (m *TagPublisher) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagPublisher) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["backfill-pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackfillPattern(val)
        }
        return nil
    }
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
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *TagPublisher) GetId()(*string) {
    return m.id
}
// GetRepo gets the repo property value. The repo property
// returns a *string when successful
func (m *TagPublisher) GetRepo()(*string) {
    return m.repo
}
// GetUpdatedAt gets the updated-at property value. The updatedAt property
// returns a *Time when successful
func (m *TagPublisher) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVcsConnectorId gets the vcs-connector-id property value. The vcsConnectorId property
// returns a *string when successful
func (m *TagPublisher) GetVcsConnectorId()(*string) {
    return m.vcsConnectorId
}
// Serialize serializes information the current object
func (m *TagPublisher) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("backfill-pattern", m.GetBackfillPattern())
        if err != nil {
            return err
        }
    }
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
func (m *TagPublisher) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBackfillPattern sets the backfill-pattern property value. The backfillPattern property
func (m *TagPublisher) SetBackfillPattern(value *string)() {
    m.backfillPattern = value
}
// SetCreatedAt sets the created-at property value. The createdAt property
func (m *TagPublisher) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetId sets the id property value. The id property
func (m *TagPublisher) SetId(value *string)() {
    m.id = value
}
// SetRepo sets the repo property value. The repo property
func (m *TagPublisher) SetRepo(value *string)() {
    m.repo = value
}
// SetUpdatedAt sets the updated-at property value. The updatedAt property
func (m *TagPublisher) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVcsConnectorId sets the vcs-connector-id property value. The vcsConnectorId property
func (m *TagPublisher) SetVcsConnectorId(value *string)() {
    m.vcsConnectorId = value
}
type TagPublisherable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackfillPattern()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetRepo()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVcsConnectorId()(*string)
    SetBackfillPattern(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetRepo(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVcsConnectorId(value *string)()
}
