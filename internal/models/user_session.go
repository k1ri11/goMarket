package models

import "time"

const TableNameUserSession = "user_session"

// UserSession mapped from table <user_session>
type UserSession struct {
	SessionID  int        `gorm:"column:session_id;type:integer;primaryKey;autoIncrement:true" json:"session_id"`
	CustomerID int        `gorm:"column:user_id;type:integer" json:"user_id"`
	StartTime  *time.Time `gorm:"column:start_time;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"start_time"`
	EndTime    *time.Time `gorm:"column:end_time;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"end_time"`
}

// TableName UserSession's table name
func (*UserSession) TableName() string {
	return TableNameUserSession
}
