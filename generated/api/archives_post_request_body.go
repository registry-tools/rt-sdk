package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ArchivesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The checksumMd5Base64 property
    checksumMd5Base64 *string
    // The contentType property
    contentType *string
    // The name property
    name *string
    // The sizeBytes property
    sizeBytes *int64
}
// NewArchivesPostRequestBody instantiates a new ArchivesPostRequestBody and sets the default values.
func NewArchivesPostRequestBody()(*ArchivesPostRequestBody) {
    m := &ArchivesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateArchivesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateArchivesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewArchivesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ArchivesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChecksumMd5Base64 gets the checksum-md5-base64 property value. The checksumMd5Base64 property
// returns a *string when successful
func (m *ArchivesPostRequestBody) GetChecksumMd5Base64()(*string) {
    return m.checksumMd5Base64
}
// GetContentType gets the content-type property value. The contentType property
// returns a *string when successful
func (m *ArchivesPostRequestBody) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ArchivesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ArchivesPostRequestBody) GetName()(*string) {
    return m.name
}
// GetSizeBytes gets the size-bytes property value. The sizeBytes property
// returns a *int64 when successful
func (m *ArchivesPostRequestBody) GetSizeBytes()(*int64) {
    return m.sizeBytes
}
// Serialize serializes information the current object
func (m *ArchivesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ArchivesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChecksumMd5Base64 sets the checksum-md5-base64 property value. The checksumMd5Base64 property
func (m *ArchivesPostRequestBody) SetChecksumMd5Base64(value *string)() {
    m.checksumMd5Base64 = value
}
// SetContentType sets the content-type property value. The contentType property
func (m *ArchivesPostRequestBody) SetContentType(value *string)() {
    m.contentType = value
}
// SetName sets the name property value. The name property
func (m *ArchivesPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetSizeBytes sets the size-bytes property value. The sizeBytes property
func (m *ArchivesPostRequestBody) SetSizeBytes(value *int64)() {
    m.sizeBytes = value
}
type ArchivesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecksumMd5Base64()(*string)
    GetContentType()(*string)
    GetName()(*string)
    GetSizeBytes()(*int64)
    SetChecksumMd5Base64(value *string)()
    SetContentType(value *string)()
    SetName(value *string)()
    SetSizeBytes(value *int64)()
}
