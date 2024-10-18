package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Archive struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The checksumMd5Base64 property
    checksumMd5Base64 *string
    // The contentType property
    contentType *string
    // The filename property
    filename *string
    // The name property
    name *string
    // The signedId property
    signedId *string
    // The sizeBytes property
    sizeBytes *int64
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
// GetChecksumMd5Base64 gets the checksum-md5-base64 property value. The checksumMd5Base64 property
// returns a *string when successful
func (m *Archive) GetChecksumMd5Base64()(*string) {
    return m.checksumMd5Base64
}
// GetContentType gets the content-type property value. The contentType property
// returns a *string when successful
func (m *Archive) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Archive) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checksum-md5-base64"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecksumMd5Base64(val)
        }
        return nil
    }
    res["content-type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["size-bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeBytes(val)
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
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Archive) GetName()(*string) {
    return m.name
}
// GetSignedId gets the signed-id property value. The signedId property
// returns a *string when successful
func (m *Archive) GetSignedId()(*string) {
    return m.signedId
}
// GetSizeBytes gets the size-bytes property value. The sizeBytes property
// returns a *int64 when successful
func (m *Archive) GetSizeBytes()(*int64) {
    return m.sizeBytes
}
// Serialize serializes information the current object
func (m *Archive) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("checksum-md5-base64", m.GetChecksumMd5Base64())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content-type", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("size-bytes", m.GetSizeBytes())
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
// SetChecksumMd5Base64 sets the checksum-md5-base64 property value. The checksumMd5Base64 property
func (m *Archive) SetChecksumMd5Base64(value *string)() {
    m.checksumMd5Base64 = value
}
// SetContentType sets the content-type property value. The contentType property
func (m *Archive) SetContentType(value *string)() {
    m.contentType = value
}
// SetFilename sets the filename property value. The filename property
func (m *Archive) SetFilename(value *string)() {
    m.filename = value
}
// SetName sets the name property value. The name property
func (m *Archive) SetName(value *string)() {
    m.name = value
}
// SetSignedId sets the signed-id property value. The signedId property
func (m *Archive) SetSignedId(value *string)() {
    m.signedId = value
}
// SetSizeBytes sets the size-bytes property value. The sizeBytes property
func (m *Archive) SetSizeBytes(value *int64)() {
    m.sizeBytes = value
}
type Archiveable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecksumMd5Base64()(*string)
    GetContentType()(*string)
    GetFilename()(*string)
    GetName()(*string)
    GetSignedId()(*string)
    GetSizeBytes()(*int64)
    SetChecksumMd5Base64(value *string)()
    SetContentType(value *string)()
    SetFilename(value *string)()
    SetName(value *string)()
    SetSignedId(value *string)()
    SetSizeBytes(value *int64)()
}
