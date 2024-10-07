package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MetaPagination struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The currentPage property
    currentPage *int32
    // The nextUrl property
    nextUrl *string
    // The pageSize property
    pageSize *int32
    // The prevUrl property
    prevUrl *string
}
// NewMetaPagination instantiates a new MetaPagination and sets the default values.
func NewMetaPagination()(*MetaPagination) {
    m := &MetaPagination{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMetaPaginationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMetaPaginationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMetaPagination(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MetaPagination) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentPage gets the current-page property value. The currentPage property
// returns a *int32 when successful
func (m *MetaPagination) GetCurrentPage()(*int32) {
    return m.currentPage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MetaPagination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["current-page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPage(val)
        }
        return nil
    }
    res["next-url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextUrl(val)
        }
        return nil
    }
    res["page-size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageSize(val)
        }
        return nil
    }
    res["prev-url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrevUrl(val)
        }
        return nil
    }
    return res
}
// GetNextUrl gets the next-url property value. The nextUrl property
// returns a *string when successful
func (m *MetaPagination) GetNextUrl()(*string) {
    return m.nextUrl
}
// GetPageSize gets the page-size property value. The pageSize property
// returns a *int32 when successful
func (m *MetaPagination) GetPageSize()(*int32) {
    return m.pageSize
}
// GetPrevUrl gets the prev-url property value. The prevUrl property
// returns a *string when successful
func (m *MetaPagination) GetPrevUrl()(*string) {
    return m.prevUrl
}
// Serialize serializes information the current object
func (m *MetaPagination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("current-page", m.GetCurrentPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("next-url", m.GetNextUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("page-size", m.GetPageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prev-url", m.GetPrevUrl())
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
func (m *MetaPagination) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentPage sets the current-page property value. The currentPage property
func (m *MetaPagination) SetCurrentPage(value *int32)() {
    m.currentPage = value
}
// SetNextUrl sets the next-url property value. The nextUrl property
func (m *MetaPagination) SetNextUrl(value *string)() {
    m.nextUrl = value
}
// SetPageSize sets the page-size property value. The pageSize property
func (m *MetaPagination) SetPageSize(value *int32)() {
    m.pageSize = value
}
// SetPrevUrl sets the prev-url property value. The prevUrl property
func (m *MetaPagination) SetPrevUrl(value *string)() {
    m.prevUrl = value
}
type MetaPaginationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPage()(*int32)
    GetNextUrl()(*string)
    GetPageSize()(*int32)
    GetPrevUrl()(*string)
    SetCurrentPage(value *int32)()
    SetNextUrl(value *string)()
    SetPageSize(value *int32)()
    SetPrevUrl(value *string)()
}
