package models

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	CategoryID  int32   `gorm:"column:category_id;type:integer;primaryKey;autoIncrement:true" json:"category_id"`
	Name        string  `gorm:"column:name;type:character varying(100);not null" json:"name"`
	Description *string `gorm:"column:description;type:text" json:"description"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
