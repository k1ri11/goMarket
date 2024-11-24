package models

import (
	"time"
)

const TableNameShippingInfo = "shipping_info"

// ShippingInfo mapped from table <shipping_info>
type ShippingInfo struct {
	ShippingID            int32      `gorm:"column:shipping_id;type:integer;primaryKey;autoIncrement:true" json:"shipping_id"`
	OrderID               *int32     `gorm:"column:order_id;type:integer" json:"order_id"`
	ShippingAddress       string     `gorm:"column:shipping_address;type:text;not null" json:"shipping_address"`
	ShippingMethod        string     `gorm:"column:shipping_method;type:character varying(50);not null" json:"shipping_method"`
	ShippingCost          float64    `gorm:"column:shipping_cost;type:numeric(10,2);not null" json:"shipping_cost"`
	EstimatedDeliveryDate *time.Time `gorm:"column:estimated_delivery_date;type:date" json:"estimated_delivery_date"`
	ActualDeliveryDate    *time.Time `gorm:"column:actual_delivery_date;type:date" json:"actual_delivery_date"`
}

// TableName ShippingInfo's table name
func (*ShippingInfo) TableName() string {
	return TableNameShippingInfo
}
