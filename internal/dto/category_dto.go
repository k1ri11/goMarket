package dto

// CreateCategoryRequest represents the input for creating a category.
type CreateCategoryRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
}

// UpdateCategoryRequest represents the input for updating a category.
type UpdateCategoryRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// CategoryResponse represents the output for a category.
type CategoryResponse struct {
	CategoryID  int32   `json:"category_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
