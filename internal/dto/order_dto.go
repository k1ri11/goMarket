package dto

import "time"

// DTO для Order

type CreateOrderRequest struct {
	CustomerID *int32  `json:"customer_id" binding:"required"`
	TotalPrice float64 `json:"total_price" binding:"required,gt=0"`
	Status     string  `json:"status" binding:"required"`
}

type UpdateOrderRequest struct {
	Status    *string    `json:"status,omitempty"`
	ShippedAt *time.Time `json:"shipped_at,omitempty" example:"2024-12-12T12:00:00Z"`
}

type OrderResponse struct {
	OrderID    *int32     `json:"order_id"`
	CustomerID *int32     `json:"customer_id"`
	TotalPrice float64    `json:"total_price"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	ShippedAt  *time.Time `json:"shipped_at,omitempty"`
}
