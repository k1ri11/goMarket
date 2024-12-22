package handlers

import (
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

type ShoppingCartHandler struct {
	service *services.ShoppingCartService
}

func NewShoppingCartHandler(service *services.ShoppingCartService) *ShoppingCartHandler {
	return &ShoppingCartHandler{service: service}
}

// GetAllShoppingCarts возвращает список всех корзин покупок.
// @Summary Получение всех корзин покупок
// @Description Возвращает список всех корзин покупок.
// @Tags Корзина
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.ShoppingCartResponse "Список корзин покупок"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/shopping_carts [get]
func (h *ShoppingCartHandler) GetAllShoppingCarts(c *gin.Context) {
	carts, err := h.service.GetAllShoppingCarts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve shopping carts"})
		return
	}

	c.JSON(http.StatusOK, carts)
}

// GetShoppingCartByID возвращает корзину покупок по ID.
// @Summary Получение корзины покупок по ID
// @Description Возвращает корзину покупок по указанному ID.
// @Tags Корзина
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID корзины"
// @Success 200 {object} dto.ShoppingCartResponse "Корзина покупок"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID корзины"
// @Failure 404 {object} dto.ErrorResponse "Корзина покупок не найдена"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/shopping_carts/{id} [get]
func (h *ShoppingCartHandler) GetShoppingCartByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	cart, err := h.service.GetShoppingCartByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shopping cart not found"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// CreateShoppingCart добавляет новую корзину покупок.
// @Summary Создание корзины покупок
// @Description Создает новую корзину покупок.
// @Tags Корзина
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateShoppingCartRequest true "Данные для создания корзины"
// @Success 201 {object} dto.ShoppingCartResponse "Корзина покупок успешно создана"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/shopping_carts [post]
func (h *ShoppingCartHandler) CreateShoppingCart(c *gin.Context) {
	var req dto.CreateShoppingCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.service.CreateShoppingCart(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shopping cart"})
		return
	}

	c.JSON(http.StatusCreated, cart)
}

// UpdateShoppingCart обновляет корзину покупок по ID.
// @Summary Обновление корзины покупок
// @Description Обновляет данные корзины покупок по указанному ID.
// @Tags Корзина
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID корзины"
// @Param request body dto.UpdateShoppingCartRequest true "Данные для обновления корзины"
// @Success 200 {object} dto.ShoppingCartResponse "Корзина покупок успешно обновлена"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID корзины"
// @Failure 404 {object} dto.ErrorResponse "Корзина покупок не найдена"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/shopping_carts/{id} [put]
func (h *ShoppingCartHandler) UpdateShoppingCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	var req dto.UpdateShoppingCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.service.UpdateShoppingCart(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shopping cart"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// DeleteShoppingCart удаляет корзину покупок по ID.
// @Summary Удаление корзины покупок
// @Description Удаляет корзину покупок по указанному ID.
// @Tags Корзина
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID корзины"
// @Success 200 {object} map[string]interface{} "Корзина покупок удалена успешно"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID корзины"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/shopping_carts/{id} [delete]
func (h *ShoppingCartHandler) DeleteShoppingCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	if err := h.service.DeleteShoppingCart(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shopping cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shopping cart deleted successfully"})
}
