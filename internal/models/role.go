package models

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	RoleID      int32   `gorm:"column:role_id;type:integer;primaryKey;autoIncrement:true" json:"role_id"`
	Name        string  `gorm:"column:name;type:character varying(50);not null" json:"name"`
	Description *string `gorm:"column:description;type:text" json:"description"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
