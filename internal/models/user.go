package models

import (
	"time"
)

const TableNameCustomer = "customer"

// User mapped from table <customer>
type User struct {
	CustomerID   int        `gorm:"column:customer_id;type:integer;primaryKey;autoIncrement:true" json:"customer_id"`
	FirstName    string     `gorm:"column:first_name;type:character varying(50);not null" json:"first_name"`
	LastName     string     `gorm:"column:last_name;type:character varying(50);not null" json:"last_name"`
	Email        string     `gorm:"column:email;type:character varying(100);not null" json:"email"`
	PasswordHash string     `gorm:"column:password_hash;type: varying(255);" json:"password_hash"`
	Phone        *string    `gorm:"column:phone;type:character varying(20)" json:"phone"`
	Address      *string    `gorm:"column:address;type:text" json:"address"`
	City         *string    `gorm:"column:city;type:character varying(50)" json:"city"`
	Country      *string    `gorm:"column:country;type:character varying(50)" json:"country"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameCustomer
}
