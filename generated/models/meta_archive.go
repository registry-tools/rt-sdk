package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MetaArchive struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The headers property
    headers MetaArchive_headersable
    // The uploadUrl property
    uploadUrl *string
}
// NewMetaArchive instantiates a new MetaArchive and sets the default values.
func NewMetaArchive()(*MetaArchive) {
    m := &MetaArchive{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMetaArchiveFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMetaArchiveFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMetaArchive(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MetaArchive) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MetaArchive) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["headers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetaArchive_headersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaders(val.(MetaArchive_headersable))
        }
        return nil
    }
    res["upload-url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadUrl(val)
        }
        return nil
    }
    return res
}
// GetHeaders gets the headers property value. The headers property
// returns a MetaArchive_headersable when successful
func (m *MetaArchive) GetHeaders()(MetaArchive_headersable) {
    return m.headers
}
// GetUploadUrl gets the upload-url property value. The uploadUrl property
// returns a *string when successful
func (m *MetaArchive) GetUploadUrl()(*string) {
    return m.uploadUrl
}
// Serialize serializes information the current object
func (m *MetaArchive) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("headers", m.GetHeaders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("upload-url", m.GetUploadUrl())
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
func (m *MetaArchive) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeaders sets the headers property value. The headers property
func (m *MetaArchive) SetHeaders(value MetaArchive_headersable)() {
    m.headers = value
}
// SetUploadUrl sets the upload-url property value. The uploadUrl property
func (m *MetaArchive) SetUploadUrl(value *string)() {
    m.uploadUrl = value
}
type MetaArchiveable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeaders()(MetaArchive_headersable)
    GetUploadUrl()(*string)
    SetHeaders(value MetaArchive_headersable)()
    SetUploadUrl(value *string)()
}
