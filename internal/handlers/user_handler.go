package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/models"
	"goMarket/internal/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) getUserIdFromContext(c *gin.Context) (int, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("no userID found in context")
	}

	// Преобразуем userID в int32
	userIDInt, err := strconv.Atoi(userID.(string)) // преобразуем строку в int
	if err != nil {
		return 0, errors.New("invalid userID")
	}
	return userIDInt, nil
}

// GetAllUsers возвращает список всех пользователей.
// @Summary Получение всех пользователей
// @Description Возвращает список всех зарегистрированных пользователей.
// @Tags Пользователь
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.UserResponseDTO "Список пользователей"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID возвращает пользователя по ID.
// @Summary Получение пользователя по ID
// @Description Возвращает данные пользователя по указанному ID.
// @Tags Пользователь
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID пользователя"
// @Success 200 {object} dto.UserResponseDTO "Пользователь"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID пользователя"
// @Failure 404 {object} dto.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Register регистрирует нового пользователя.
// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя и регистрирует его в системе.
// @Tags Авторизация
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "Данные для регистрации пользователя"
// @Success 201 {object} map[string]interface{} "Пользователь успешно зарегистрирован"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Phone:        req.Phone,
		Address:      req.Address,
		City:         req.City,
		Country:      req.Country,
	}

	err = h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// GetCurrentUser возвращает информацию о текущем пользователе.
// @Summary Получение информации о текущем пользователе
// @Description Возвращает данные текущего пользователя, основываясь на информации из контекста.
// @Tags Авторизация
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserResponseDTO "Информация о текущем пользователе"
// @Failure 401 {object} dto.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/users/me [get]
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, err := h.getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	response := dto.UserResponseDTO{
		CustomerID: user.CustomerID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Phone:      user.Phone,
		Address:    user.Address,
		City:       user.City,
		Country:    user.Country,
		CreatedAt:  user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}

// UpdateCurrentUser обновляет информацию о текущем пользователе.
// @Summary Обновление данных текущего пользователя
// @Description Обновляет данные текущего пользователя на основе запроса.
// @Tags Пользователь
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateUserRequest true "Данные для обновления пользователя"
// @Success 200 {object} dto.UserResponseDTO "Пользователь успешно обновлен"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 404 {object} dto.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/users/current [put]
func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
	userID, err := h.getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateCustomer(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser удаляет пользователя по ID.
// @Summary Удаление пользователя
// @Description Удаляет пользователя из системы по указанному ID.
// @Tags Пользователь
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]interface{} "Пользователь успешно удален"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID пользователя"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.DeleteCustomer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
