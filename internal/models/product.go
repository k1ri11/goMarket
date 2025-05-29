package models

import (
	"time"
)

const TableNameProduct = "product"

// Product mapped from table <product>
type Product struct {
	ProductID      int32      `gorm:"column:product_id;type:integer;primaryKey;autoIncrement:true" json:"product_id"`
	Name           string     `gorm:"column:name;type:character varying(100);not null" json:"name"`
	Brand          *string    `gorm:"column:brand;type:character varying(50)" json:"brand"`
	Model          *string    `gorm:"column:model;type:character varying(50)" json:"model"`
	Price          float64    `gorm:"column:price;type:numeric(10,2);not null" json:"price"`
	Stock          *int32     `gorm:"column:stock;type:integer" json:"stock"`
	Description    *string    `gorm:"column:description;type:text" json:"description"`
	CreatedAt      *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	ReviewQuantity *int32     `gorm:"column:review_quantity" json:"review_quantity"`
	BasePrice      *float64   `gorm:"column:base_price" json:"base_price"`
	Rating         *float64   `gorm:"column:rating" json:"rating"`
	ImageURL       *string    `gorm:"column:image_url" json:"image_url"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
