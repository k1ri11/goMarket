package models

import (
	"time"
)

const TableNamePayment = "payment"

// Payment mapped from table <payment>
type Payment struct {
	PaymentID     int32      `gorm:"column:payment_id;type:integer;primaryKey;autoIncrement:true" json:"payment_id"`
	OrderID       *int32     `gorm:"column:order_id;type:integer" json:"order_id"`
	Amount        float64    `gorm:"column:amount;type:numeric(10,2);not null" json:"amount"`
	PaymentMethod string     `gorm:"column:payment_method;type:character varying(50);not null" json:"payment_method"`
	PaymentStatus *string    `gorm:"column:payment_status;type:character varying(50);default:pending" json:"payment_status"`
	PaymentDate   *time.Time `gorm:"column:payment_date;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"payment_date"`
}

// TableName Payment's table name
func (*Payment) TableName() string {
	return TableNamePayment
}
