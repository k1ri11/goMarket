package models

const TableNameOrderItem = "order_item"

// OrderItem mapped from table <order_item>
type OrderItem struct {
	OrderItemID int32   `gorm:"column:order_item_id;type:integer;primaryKey;autoIncrement:true" json:"order_item_id"`
	OrderID     *int32  `gorm:"column:order_id;type:integer" json:"order_id"`
	ProductID   *int32  `gorm:"column:product_id;type:integer" json:"product_id"`
	Quantity    int32   `gorm:"column:quantity;type:integer;not null" json:"quantity"`
	Price       float64 `gorm:"column:price;type:numeric(10,2);not null" json:"price"`
}

// TableName OrderItem's table name
func (*OrderItem) TableName() string {
	return TableNameOrderItem
}
