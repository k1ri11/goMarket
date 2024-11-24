package models

import (
	"time"
)

const TableNameOrder = "order"

// Order mapped from table <order>
type Order struct {
	OrderID    *int32     `gorm:"column:order_id;type:integer;primaryKey;autoIncrement:true"`
	CustomerID *int32     `gorm:"column:customer_id;type:integer" json:"customer_id"`
	TotalPrice float64    `gorm:"column:total_price;type:numeric(10,2);not null" json:"total_price"`
	Status     string     `gorm:"column:status;type:character varying(50);not null" json:"status"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	ShippedAt  *time.Time `gorm:"column:shipped_at;type:timestamp without time zone" json:"shipped_at"`
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
