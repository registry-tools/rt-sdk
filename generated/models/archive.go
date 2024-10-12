package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Archive struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The filename property
    filename *string
    // The signedId property
    signedId *string
}
// NewArchive instantiates a new Archive and sets the default values.
func NewArchive()(*Archive) {
    m := &Archive{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateArchiveFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateArchiveFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewArchive(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Archive) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Archive) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["filename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilename(val)
        }
        return nil
    }
    res["signed-id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedId(val)
        }
        return nil
    }
    return res
}
// GetFilename gets the filename property value. The filename property
// returns a *string when successful
func (m *Archive) GetFilename()(*string) {
    return m.filename
}
// GetSignedId gets the signed-id property value. The signedId property
// returns a *string when successful
func (m *Archive) GetSignedId()(*string) {
    return m.signedId
}
// Serialize serializes information the current object
func (m *Archive) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filename", m.GetFilename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signed-id", m.GetSignedId())
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
func (m *Archive) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFilename sets the filename property value. The filename property
func (m *Archive) SetFilename(value *string)() {
    m.filename = value
}
// SetSignedId sets the signed-id property value. The signedId property
func (m *Archive) SetSignedId(value *string)() {
    m.signedId = value
}
type Archiveable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilename()(*string)
    GetSignedId()(*string)
    SetFilename(value *string)()
    SetSignedId(value *string)()
}
