package models

import (
	"time"
)

const TableNameUserSession = "user_session"

// UserSession mapped from table <user_session>
type UserSession struct {
	SessionID    int32      `gorm:"column:session_id;type:integer;primaryKey;autoIncrement:true" json:"session_id"`
	CustomerID   *int32     `gorm:"column:customer_id;type:integer" json:"customer_id"`
	SessionToken string     `gorm:"column:session_token;type:text;not null" json:"session_token"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	ExpiresAt    *time.Time `gorm:"column:expires_at;type:timestamp without time zone" json:"expires_at"`
}

// TableName UserSession's table name
func (*UserSession) TableName() string {
	return TableNameUserSession
}
