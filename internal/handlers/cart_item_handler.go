package handlers

import (
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

type CartItemHandler struct {
	service *services.CartItemService
}

func NewCartItemHandler(service *services.CartItemService) *CartItemHandler {
	return &CartItemHandler{service: service}
}

// GetAllCartItems
// @Summary Получение всех элементов корзины
// @Description Возвращает список всех элементов корзины
// @Tags Элемент корзины
// @Produce json
// @Success 200 {array} dto.CartItemResponse "Список элементов корзины"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/cart_items [get]
func (h *CartItemHandler) GetAllCartItems(c *gin.Context) {
	items, err := h.service.GetAllCartItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetCartItemByID
// @Summary Получение элемента корзины по ID
// @Description Возвращает информацию об элементе корзины по его ID
// @Tags Элемент корзины
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID элемента корзины"
// @Success 200 {object} dto.CartItemResponse "Информация об элементе корзины"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID"
// @Failure 404 {object} dto.ErrorResponse "Элемент корзины не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/cart_items/{id} [get]
func (h *CartItemHandler) GetCartItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	item, err := h.service.GetCartItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateCartItem
// @Summary Добавление нового элемента в корзину
// @Description Создает новый элемент корзины
// @Tags Элемент корзины
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateCartItemRequest true "Данные для нового элемента корзины"
// @Success 201 {object} dto.CartItemResponse "Созданный элемент корзины"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/cart_items [post]
func (h *CartItemHandler) CreateCartItem(c *gin.Context) {
	var req dto.CreateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.CreateCartItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateCartItem
// @Summary Обновление элемента корзины
// @Description Обновляет данные элемента корзины по его ID
// @Tags Элемент корзины
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID элемента корзины"
// @Param request body dto.UpdateCartItemRequest true "Данные для обновления элемента корзины"
// @Success 200 {object} dto.CartItemResponse "Обновленный элемент корзины"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации данных или некорректный ID"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/cart_items/{id} [put]
func (h *CartItemHandler) UpdateCartItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	var req dto.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.UpdateCartItem(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteCartItem
// @Summary Удаление элемента корзины
// @Description Удаляет элемент корзины по его ID
// @Tags Элемент корзины
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID элемента корзины"
// @Success 200 {object} map[string]interface{} "Успешное удаление"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/cart_items/{id} [delete]
func (h *CartItemHandler) DeleteCartItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	if err := h.service.DeleteCartItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item deleted successfully"})
}
