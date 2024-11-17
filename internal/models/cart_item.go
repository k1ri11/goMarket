package models

const TableNameCartItem = "cart_item"

// CartItem mapped from table <cart_item>
type CartItem struct {
	CartItemID int32  `gorm:"column:cart_item_id;type:integer;primaryKey;autoIncrement:true" json:"cart_item_id"`
	CartID     *int32 `gorm:"column:cart_id;type:integer" json:"cart_id"`
	ProductID  *int32 `gorm:"column:product_id;type:integer" json:"product_id"`
	Quantity   int32  `gorm:"column:quantity;type:integer;not null" json:"quantity"`
}

// TableName CartItem's table name
func (*CartItem) TableName() string {
	return TableNameCartItem
}
