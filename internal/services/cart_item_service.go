package services

import (
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"gorm.io/gorm"
)

type CartItemService struct {
	db *gorm.DB
}

func NewCartItemService(db *gorm.DB) *CartItemService {
	return &CartItemService{db: db}
}

func (s *CartItemService) GetAllCartItems() ([]dto.CartItemResponse, error) {
	var items []models.CartItem
	if err := s.db.Find(&items).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.CartItemResponse, len(items))
	for i, item := range items {
		responses[i] = dto.CartItemResponse{
			CartItemID: item.CartItemID,
			CartID:     item.CartID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
		}
	}

	return responses, nil
}

func (s *CartItemService) GetCartItemByID(id int) (*dto.CartItemResponse, error) {
	var item models.CartItem
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}

	response := &dto.CartItemResponse{
		CartItemID: item.CartItemID,
		CartID:     item.CartID,
		ProductID:  item.ProductID,
		Quantity:   item.Quantity,
	}

	return response, nil
}

func (s *CartItemService) CreateCartItem(req dto.CreateCartItemRequest) (*dto.CartItemResponse, error) {
	item := models.CartItem{
		CartID:    req.CartID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := s.db.Create(&item).Error; err != nil {
		return nil, err
	}

	return &dto.CartItemResponse{
		CartItemID: item.CartItemID,
		CartID:     item.CartID,
		ProductID:  item.ProductID,
		Quantity:   item.Quantity,
	}, nil
}

func (s *CartItemService) UpdateCartItem(id int, req dto.UpdateCartItemRequest) (*dto.CartItemResponse, error) {
	var item models.CartItem
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}

	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}

	if err := s.db.Save(&item).Error; err != nil {
		return nil, err
	}

	return &dto.CartItemResponse{
		CartItemID: item.CartItemID,
		CartID:     item.CartID,
		ProductID:  item.ProductID,
		Quantity:   item.Quantity,
	}, nil
}

func (s *CartItemService) DeleteCartItem(id int) error {
	if err := s.db.Delete(&models.CartItem{}, id).Error; err != nil {
		return err
	}

	return nil
}
