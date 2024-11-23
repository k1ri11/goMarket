package services

import (
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) GetAllOrders() ([]dto.OrderResponse, error) {
	var orders []models.Order
	if err := s.db.Find(&orders).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.OrderResponse, len(orders))
	for i, order := range orders {
		responses[i] = dto.OrderResponse{
			OrderID:    order.OrderID,
			CustomerID: order.CustomerID,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt,
			ShippedAt:  order.ShippedAt,
		}
	}

	return responses, nil
}

func (s *OrderService) GetOrderByID(id int) (*dto.OrderResponse, error) {
	var order models.Order
	if err := s.db.First(&order, id).Error; err != nil {
		return nil, err
	}

	response := &dto.OrderResponse{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		ShippedAt:  order.ShippedAt,
	}

	return response, nil
}

func (s *OrderService) CreateOrder(req dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	order := models.Order{
		CustomerID: req.CustomerID,
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}

	if err := s.db.Create(&order).Error; err != nil {
		return nil, err
	}

	return &dto.OrderResponse{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		ShippedAt:  order.ShippedAt,
	}, nil
}

func (s *OrderService) UpdateOrder(id int, req dto.UpdateOrderRequest) (*dto.OrderResponse, error) {
	var order models.Order
	if err := s.db.First(&order, id).Error; err != nil {
		return nil, err
	}

	if req.Status != nil {
		order.Status = *req.Status
	}
	if req.ShippedAt != nil {
		order.ShippedAt = req.ShippedAt
	}

	if err := s.db.Save(&order).Error; err != nil {
		return nil, err
	}

	return &dto.OrderResponse{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		ShippedAt:  order.ShippedAt,
	}, nil
}

func (s *OrderService) DeleteOrder(id int) error {
	if err := s.db.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}

	return nil
}
