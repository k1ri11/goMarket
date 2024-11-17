package models

import (
	"time"
)

const TableNameReview = "review"

// Review mapped from table <review>
type Review struct {
	ReviewID   int32      `gorm:"column:review_id;type:integer;primaryKey;autoIncrement:true" json:"review_id"`
	ProductID  *int32     `gorm:"column:product_id;type:integer" json:"product_id"`
	CustomerID *int32     `gorm:"column:customer_id;type:integer" json:"customer_id"`
	Rating     *int32     `gorm:"column:rating;type:integer" json:"rating"`
	Comment    *string    `gorm:"column:comment;type:text" json:"comment"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Review's table name
func (*Review) TableName() string {
	return TableNameReview
}
