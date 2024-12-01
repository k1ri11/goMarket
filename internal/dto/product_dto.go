package dto

import "time"

// CreateProductRequest represents the input for creating a product.
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Brand       *string `json:"brand,omitempty"`
	Model       *string `json:"model,omitempty"`
	Price       float64 `json:"price" binding:"required"`
	Stock       *int32  `json:"stock,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateProductRequest represents the input for updating a product.
type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty"`
	Brand       *string  `json:"brand,omitempty"`
	Model       *string  `json:"model,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int32   `json:"stock,omitempty"`
	Description *string  `json:"description,omitempty"`
}

type ProductFilterRequest struct {
	Name       *string  `form:"name,omitempty"`
	Brand      *string  `form:"brand,omitempty"`
	CategoryID *int32   `form:"category_id,omitempty"`
	MinPrice   *float64 `form:"min_price,omitempty"`
	MaxPrice   *float64 `form:"max_price,omitempty"`
	Page       int      `form:"page" binding:"gte=1"`
	PageSize   int      `form:"page_size" binding:"gte=1"`
}

// ProductResponse represents the output returned to the client.
type ProductResponse struct {
	ProductID   int32      `json:"product_id"`
	Name        string     `json:"name"`
	Brand       *string    `json:"brand"`
	Model       *string    `json:"model"`
	Price       float64    `json:"price"`
	Stock       *int32     `json:"stock"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
}

type ProductResponsePagination struct {
	Data       []ProductResponse `json:"data"`
	TotalCount int64             `json:"total_count"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
}
