package services

import (
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"gorm.io/gorm"
)

type ShoppingCartService struct {
	db *gorm.DB
}

func NewShoppingCartService(db *gorm.DB) *ShoppingCartService {
	return &ShoppingCartService{db: db}
}

func (s *ShoppingCartService) GetAllShoppingCarts() ([]dto.ShoppingCartResponse, error) {
	var carts []models.ShoppingCart
	if err := s.db.Find(&carts).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.ShoppingCartResponse, len(carts))
	for i, cart := range carts {
		responses[i] = dto.ShoppingCartResponse{
			CartID:     cart.CartID,
			CustomerID: cart.CustomerID,
			CreatedAt:  cart.CreatedAt,
		}
	}

	return responses, nil
}

func (s *ShoppingCartService) GetShoppingCartByID(id int) (*dto.ShoppingCartResponse, error) {
	var cart models.ShoppingCart
	if err := s.db.First(&cart, id).Error; err != nil {
		return nil, err
	}

	response := &dto.ShoppingCartResponse{
		CartID:     cart.CartID,
		CustomerID: cart.CustomerID,
		CreatedAt:  cart.CreatedAt,
	}

	return response, nil
}

func (s *ShoppingCartService) CreateShoppingCart(req dto.CreateShoppingCartRequest) (*dto.ShoppingCartResponse, error) {
	cart := models.ShoppingCart{
		CustomerID: req.CustomerID,
	}

	if err := s.db.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &dto.ShoppingCartResponse{
		CartID:     cart.CartID,
		CustomerID: cart.CustomerID,
		CreatedAt:  cart.CreatedAt,
	}, nil
}

func (s *ShoppingCartService) UpdateShoppingCart(id int, req dto.UpdateShoppingCartRequest) (*dto.ShoppingCartResponse, error) {
	var cart models.ShoppingCart
	if err := s.db.First(&cart, id).Error; err != nil {
		return nil, err
	}

	if req.CustomerID != nil {
		cart.CustomerID = req.CustomerID
	}

	if err := s.db.Save(&cart).Error; err != nil {
		return nil, err
	}

	return &dto.ShoppingCartResponse{
		CartID:     cart.CartID,
		CustomerID: cart.CustomerID,
		CreatedAt:  cart.CreatedAt,
	}, nil
}

func (s *ShoppingCartService) DeleteShoppingCart(id int) error {
	if err := s.db.Delete(&models.ShoppingCart{}, id).Error; err != nil {
		return err
	}

	return nil
}
