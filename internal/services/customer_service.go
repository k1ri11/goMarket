package services

import (
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) GetAllCustomers() ([]dto.CustomerResponse, error) {
	var customers []models.Customer
	if err := s.db.Find(&customers).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.CustomerResponse, len(customers))
	for i, customer := range customers {
		responses[i] = dto.CustomerResponse{
			CustomerID: customer.CustomerID,
			FirstName:  customer.FirstName,
			LastName:   customer.LastName,
			Email:      customer.Email,
			Phone:      customer.Phone,
			Address:    customer.Address,
			City:       customer.City,
			Country:    customer.Country,
			CreatedAt:  customer.CreatedAt,
		}
	}

	return responses, nil
}

func (s *CustomerService) GetCustomerByID(id int) (*dto.CustomerResponse, error) {
	var customer models.Customer
	if err := s.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	response := &dto.CustomerResponse{
		CustomerID: customer.CustomerID,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		City:       customer.City,
		Country:    customer.Country,
		CreatedAt:  customer.CreatedAt,
	}

	return response, nil
}

func (s *CustomerService) CreateCustomer(req dto.CreateCustomerRequest) (*dto.CustomerResponse, error) {
	customer := models.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		City:      req.City,
		Country:   req.Country,
	}

	if err := s.db.Create(&customer).Error; err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		CustomerID: customer.CustomerID,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		City:       customer.City,
		Country:    customer.Country,
		CreatedAt:  customer.CreatedAt,
	}, nil
}

func (s *CustomerService) UpdateCustomer(id int, req dto.UpdateCustomerRequest) (*dto.CustomerResponse, error) {
	var customer models.Customer
	if err := s.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	if req.FirstName != nil {
		customer.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		customer.LastName = *req.LastName
	}
	if req.Email != nil {
		customer.Email = *req.Email
	}
	if req.Phone != nil {
		customer.Phone = req.Phone
	}
	if req.Address != nil {
		customer.Address = req.Address
	}
	if req.City != nil {
		customer.City = req.City
	}
	if req.Country != nil {
		customer.Country = req.Country
	}

	if err := s.db.Save(&customer).Error; err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		CustomerID: customer.CustomerID,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		City:       customer.City,
		Country:    customer.Country,
		CreatedAt:  customer.CreatedAt,
	}, nil
}

func (s *CustomerService) DeleteCustomer(id int) error {
	if err := s.db.Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}
