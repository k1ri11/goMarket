package models

const TableNameInventory = "inventory"

// Inventory mapped from table <inventory>
type Inventory struct {
	InventoryID       int32   `gorm:"column:inventory_id;type:integer;primaryKey;autoIncrement:true" json:"inventory_id"`
	ProductID         *int32  `gorm:"column:product_id;type:integer" json:"product_id"`
	Quantity          int32   `gorm:"column:quantity;type:integer;not null" json:"quantity"`
	WarehouseLocation *string `gorm:"column:warehouse_location;type:character varying(100)" json:"warehouse_location"`
}

// TableName Inventory's table name
func (*Inventory) TableName() string {
	return TableNameInventory
}
