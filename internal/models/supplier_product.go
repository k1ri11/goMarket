package models

const TableNameSupplierProduct = "supplier_product"

// SupplierProduct mapped from table <supplier_product>
type SupplierProduct struct {
	SupplierProductID int32   `gorm:"column:supplier_product_id;type:integer;primaryKey;autoIncrement:true" json:"supplier_product_id"`
	SupplierID        *int32  `gorm:"column:supplier_id;type:integer" json:"supplier_id"`
	ProductID         *int32  `gorm:"column:product_id;type:integer" json:"product_id"`
	SupplyPrice       float64 `gorm:"column:supply_price;type:numeric(10,2);not null" json:"supply_price"`
	Quantity          int32   `gorm:"column:quantity;type:integer;not null" json:"quantity"`
}

// TableName SupplierProduct's table name
func (*SupplierProduct) TableName() string {
	return TableNameSupplierProduct
}
