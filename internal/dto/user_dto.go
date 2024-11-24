package dto

import "time"

type CreateUserRequest struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=6"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty" binding:"omitempty,email"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
}

type UserResponseDTO struct {
	CustomerID int        `json:"customer_id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Phone      *string    `json:"phone"`
	Address    *string    `json:"address"`
	City       *string    `json:"city"`
	Country    *string    `json:"country"`
	CreatedAt  *time.Time `json:"created_at"`
}
