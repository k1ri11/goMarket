package dto

import "time"

// DTO для ShoppingCart
type CreateShoppingCartRequest struct {
	CustomerID *int32 `json:"customer_id" binding:"required"`
}

type UpdateShoppingCartRequest struct {
	CustomerID *int32 `json:"customer_id,omitempty"`
}

type ShoppingCartResponse struct {
	CartID     int32      `json:"cart_id"`
	CustomerID *int32     `json:"customer_id"`
	CreatedAt  *time.Time `json:"created_at"`
}
