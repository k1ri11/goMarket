package models

const TableNameProductCategory = "product_category"

// ProductCategory mapped from table <product_category>
type ProductCategory struct {
	ProductCategoryID int32  `gorm:"column:product_category_id;type:integer;primaryKey;autoIncrement:true" json:"product_category_id"`
	ProductID         *int32 `gorm:"column:product_id;type:integer" json:"product_id"`
	CategoryID        *int32 `gorm:"column:category_id;type:integer" json:"category_id"`
}

// TableName ProductCategory's table name
func (*ProductCategory) TableName() string {
	return TableNameProductCategory
}
