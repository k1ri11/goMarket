package dto

// DTO для CartItem

type CreateCartItemRequest struct {
	CartID    *int32 `json:"cart_id" binding:"required"`
	ProductID *int32 `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartItemRequest struct {
	Quantity *int32 `json:"quantity,omitempty" binding:"omitempty,min=1"`
}

type CartItemResponse struct {
	CartItemID int32  `json:"cart_item_id"`
	CartID     *int32 `json:"cart_id"`
	ProductID  *int32 `json:"product_id"`
	Quantity   int32  `json:"quantity"`
}
