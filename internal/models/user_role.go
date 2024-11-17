package models

const TableNameUserRole = "user_role"

// UserRole mapped from table <user_role>
type UserRole struct {
	UserRoleID int32  `gorm:"column:user_role_id;type:integer;primaryKey;autoIncrement:true" json:"user_role_id"`
	CustomerID *int32 `gorm:"column:customer_id;type:integer" json:"customer_id"`
	RoleID     *int32 `gorm:"column:role_id;type:integer" json:"role_id"`
}

// TableName UserRole's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}
