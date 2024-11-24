package dto

// DTO для OrderItem

type CreateOrderItemRequest struct {
	OrderID   *int32  `json:"order_id" binding:"required"`
	ProductID *int32  `json:"product_id" binding:"required"`
	Quantity  int32   `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

type UpdateOrderItemRequest struct {
	Quantity *int32   `json:"quantity,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}

type OrderItemResponse struct {
	OrderItemID int32   `json:"order_item_id"`
	OrderID     *int32  `json:"order_id"`
	ProductID   *int32  `json:"product_id"`
	Quantity    int32   `json:"quantity"`
	Price       float64 `json:"price"`
}
