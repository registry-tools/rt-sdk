package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NamespacesItemTagPublishersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tagPublisher property
    tagPublisher NamespacesItemTagPublishersPostRequestBody_tagPublisherable
}
// NewNamespacesItemTagPublishersPostRequestBody instantiates a new NamespacesItemTagPublishersPostRequestBody and sets the default values.
func NewNamespacesItemTagPublishersPostRequestBody()(*NamespacesItemTagPublishersPostRequestBody) {
    m := &NamespacesItemTagPublishersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesItemTagPublishersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesItemTagPublishersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesItemTagPublishersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesItemTagPublishersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesItemTagPublishersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tag-publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNamespacesItemTagPublishersPostRequestBody_tagPublisherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagPublisher(val.(NamespacesItemTagPublishersPostRequestBody_tagPublisherable))
        }
        return nil
    }
    return res
}
// GetTagPublisher gets the tag-publisher property value. The tagPublisher property
// returns a NamespacesItemTagPublishersPostRequestBody_tagPublisherable when successful
func (m *NamespacesItemTagPublishersPostRequestBody) GetTagPublisher()(NamespacesItemTagPublishersPostRequestBody_tagPublisherable) {
    return m.tagPublisher
}
// Serialize serializes information the current object
func (m *NamespacesItemTagPublishersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("tag-publisher", m.GetTagPublisher())
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
func (m *NamespacesItemTagPublishersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTagPublisher sets the tag-publisher property value. The tagPublisher property
func (m *NamespacesItemTagPublishersPostRequestBody) SetTagPublisher(value NamespacesItemTagPublishersPostRequestBody_tagPublisherable)() {
    m.tagPublisher = value
}
type NamespacesItemTagPublishersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTagPublisher()(NamespacesItemTagPublishersPostRequestBody_tagPublisherable)
    SetTagPublisher(value NamespacesItemTagPublishersPostRequestBody_tagPublisherable)()
}
