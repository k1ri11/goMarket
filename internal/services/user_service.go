package services

import (
	"errors"
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetAllUsers() ([]dto.UserResponseDTO, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponseDTO, len(users))
	for i, customer := range users {
		responses[i] = dto.UserResponseDTO{
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

func (s *UserService) GetUserByID(id int) (*dto.UserResponseDTO, error) {
	var customer models.User
	if err := s.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	response := &dto.UserResponseDTO{
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

func (s *UserService) CreateUser(user models.User) error {
	if err := s.db.Create(&user).Error; err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)
	// Добавляем пользователя в базу
	s.db.Create(user)
	return nil
}

func (s *UserService) UpdateCustomer(id int, req dto.UpdateUserRequest) (*dto.UserResponseDTO, error) {
	var customer models.User
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

	return &dto.UserResponseDTO{
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

func (s *UserService) DeleteCustomer(id int) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	// Проверяем пароль
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}
