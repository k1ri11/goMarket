package services

import (
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) GetAllProducts() ([]dto.ProductResponse, error) {
	var products []models.Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = dto.ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Brand:       product.Brand,
			Model:       product.Model,
			Price:       product.Price,
			Stock:       product.Stock,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
		}
	}

	return responses, nil
}

func (s *ProductService) GetProductByID(id int) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	response := &dto.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Brand:       product.Brand,
		Model:       product.Model,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
	}

	return response, nil
}

func (s *ProductService) CreateProduct(req dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		Name:        req.Name,
		Brand:       req.Brand,
		Model:       req.Model,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
	}

	if err := s.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Brand:       product.Brand,
		Model:       product.Model,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
	}, nil
}

func (s *ProductService) UpdateProduct(id int, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	if req.Model != nil {
		product.Model = req.Model
	}
	if req.Brand != nil {
		product.Brand = req.Brand
	}
	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Stock != nil {
		product.Stock = req.Stock
	}

	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Brand:       product.Brand,
		Model:       product.Model,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
	}, nil
}

func (s *ProductService) DeleteProduct(id int) error {
	if err := s.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
