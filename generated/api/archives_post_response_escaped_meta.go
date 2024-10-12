package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ArchivesPostResponse_meta struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The headers property
    headers ArchivesPostResponse_meta_headersable
    // The uploadUrl property
    uploadUrl *string
}
// NewArchivesPostResponse_meta instantiates a new ArchivesPostResponse_meta and sets the default values.
func NewArchivesPostResponse_meta()(*ArchivesPostResponse_meta) {
    m := &ArchivesPostResponse_meta{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateArchivesPostResponse_metaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateArchivesPostResponse_metaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewArchivesPostResponse_meta(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ArchivesPostResponse_meta) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ArchivesPostResponse_meta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["headers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateArchivesPostResponse_meta_headersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaders(val.(ArchivesPostResponse_meta_headersable))
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
// returns a ArchivesPostResponse_meta_headersable when successful
func (m *ArchivesPostResponse_meta) GetHeaders()(ArchivesPostResponse_meta_headersable) {
    return m.headers
}
// GetUploadUrl gets the upload-url property value. The uploadUrl property
// returns a *string when successful
func (m *ArchivesPostResponse_meta) GetUploadUrl()(*string) {
    return m.uploadUrl
}
// Serialize serializes information the current object
func (m *ArchivesPostResponse_meta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ArchivesPostResponse_meta) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeaders sets the headers property value. The headers property
func (m *ArchivesPostResponse_meta) SetHeaders(value ArchivesPostResponse_meta_headersable)() {
    m.headers = value
}
// SetUploadUrl sets the upload-url property value. The uploadUrl property
func (m *ArchivesPostResponse_meta) SetUploadUrl(value *string)() {
    m.uploadUrl = value
}
type ArchivesPostResponse_metaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeaders()(ArchivesPostResponse_meta_headersable)
    GetUploadUrl()(*string)
    SetHeaders(value ArchivesPostResponse_meta_headersable)()
    SetUploadUrl(value *string)()
}
