package models

import (
	"time"
)

const TableNameShoppingCart = "shopping_cart"

// ShoppingCart mapped from table <shopping_cart>
type ShoppingCart struct {
	CartID     int32      `gorm:"column:cart_id;type:integer;primaryKey;autoIncrement:true" json:"cart_id"`
	CustomerID *int32     `gorm:"column:customer_id;type:integer" json:"customer_id"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName ShoppingCart's table name
func (*ShoppingCart) TableName() string {
	return TableNameShoppingCart
}
