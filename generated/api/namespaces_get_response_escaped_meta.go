package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8 "github.com/registry-tools/rt-sdk/generated/models"
)

type NamespacesGetResponse_meta struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The pagination property
    pagination ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.MetaPaginationable
}
// NewNamespacesGetResponse_meta instantiates a new NamespacesGetResponse_meta and sets the default values.
func NewNamespacesGetResponse_meta()(*NamespacesGetResponse_meta) {
    m := &NamespacesGetResponse_meta{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamespacesGetResponse_metaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamespacesGetResponse_metaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamespacesGetResponse_meta(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamespacesGetResponse_meta) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamespacesGetResponse_meta) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pagination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.CreateMetaPaginationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagination(val.(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.MetaPaginationable))
        }
        return nil
    }
    return res
}
// GetPagination gets the pagination property value. The pagination property
// returns a MetaPaginationable when successful
func (m *NamespacesGetResponse_meta) GetPagination()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.MetaPaginationable) {
    return m.pagination
}
// Serialize serializes information the current object
func (m *NamespacesGetResponse_meta) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pagination", m.GetPagination())
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
func (m *NamespacesGetResponse_meta) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPagination sets the pagination property value. The pagination property
func (m *NamespacesGetResponse_meta) SetPagination(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.MetaPaginationable)() {
    m.pagination = value
}
type NamespacesGetResponse_metaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPagination()(ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.MetaPaginationable)
    SetPagination(value ib0a1d83fbec960f9e17742ee01031c795b4f720ed854aa216c9d12a10c9701e8.MetaPaginationable)()
}
