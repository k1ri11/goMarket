package services

import (
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"gorm.io/gorm"
)

type OrderItemService struct {
	db *gorm.DB
}

func NewOrderItemService(db *gorm.DB) *OrderItemService {
	return &OrderItemService{db: db}
}

func (s *OrderItemService) GetAllOrderItems() ([]dto.OrderItemResponse, error) {
	var orderItems []models.OrderItem
	if err := s.db.Find(&orderItems).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.OrderItemResponse, len(orderItems))
	for i, item := range orderItems {
		responses[i] = dto.OrderItemResponse{
			OrderItemID: item.OrderItemID,
			OrderID:     item.OrderID,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			Price:       item.Price,
		}
	}

	return responses, nil
}

func (s *OrderItemService) GetOrderItemByID(id int) (*dto.OrderItemResponse, error) {
	var item models.OrderItem
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}

	response := &dto.OrderItemResponse{
		OrderItemID: item.OrderItemID,
		OrderID:     item.OrderID,
		ProductID:   item.ProductID,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}

	return response, nil
}

func (s *OrderItemService) CreateOrderItem(req dto.CreateOrderItemRequest) (*dto.OrderItemResponse, error) {
	item := models.OrderItem{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Price:     req.Price,
	}

	if err := s.db.Create(&item).Error; err != nil {
		return nil, err
	}

	return &dto.OrderItemResponse{
		OrderItemID: item.OrderItemID,
		OrderID:     item.OrderID,
		ProductID:   item.ProductID,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}, nil
}

func (s *OrderItemService) UpdateOrderItem(id int, req dto.UpdateOrderItemRequest) (*dto.OrderItemResponse, error) {
	var item models.OrderItem
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}

	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}
	if req.Price != nil {
		item.Price = *req.Price
	}

	if err := s.db.Save(&item).Error; err != nil {
		return nil, err
	}

	return &dto.OrderItemResponse{
		OrderItemID: item.OrderItemID,
		OrderID:     item.OrderID,
		ProductID:   item.ProductID,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}, nil
}

func (s *OrderItemService) DeleteOrderItem(id int) error {
	if err := s.db.Delete(&models.OrderItem{}, id).Error; err != nil {
		return err
	}

	return nil
}
