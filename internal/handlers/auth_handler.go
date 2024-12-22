package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

// Структура для работы с сервисом авторизации
type AuthHandler struct {
	authService *services.JWTAuthService
}

func NewAuthHandler(authService *services.JWTAuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func getUserIdFromContext(c *gin.Context) (int32, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("no userID found in context")
	}
	// Преобразуем userID в int32
	userIDInt, err := strconv.Atoi(userID.(string)) // преобразуем строку в int
	if err != nil {
		return 0, errors.New("invalid userID")
	}
	return int32(userIDInt), nil
}

// Login
// @Summary Вход пользователя
// @Description Авторизация пользователя с использованием email и пароля
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Данные для входа"
// @Success 200 {object} dto.LoginResponse "Успешная авторизация"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 401 {object} dto.ErrorResponse "Ошибка авторизации"
// @Router /v1/auth/jwt/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Вызов метода авторизации
	loginResponse, err := h.authService.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Возвращаем токен в ответе
	c.JSON(http.StatusOK, loginResponse)
}

// Logout
// @Summary Выход пользователя
// @Description Завершение сессии пользователя
// @Tags Авторизация
// @Produce json
// @Success 200 {object} map[string]interface{} "Успешный выход"
// @Failure 401 {object} dto.LoginResponse "Ошибка аутентификации"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/auth/jwt/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	userID, err := getUserIdFromContext(c) // Предполагаем, что userID хранится в контексте после авторизации
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	err = h.authService.Logout(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
