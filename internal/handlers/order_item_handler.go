package handlers

import (
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

type OrderItemHandler struct {
	service *services.OrderItemService
}

func NewOrderItemHandler(service *services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{service: service}
}

// GetAllOrderItems
// @Summary      Получение всех товаров в заказах
// @Description  Возвращает список всех товаров, связанных с заказами
// @Tags         Элемент заказа
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Success      200  {array}  dto.OrderItemResponse
// @Failure      500  {object}  dto.ErrorResponse "Ошибка получения товаров в заказах"
// @Router       /v1/order_items [get]
func (h *OrderItemHandler) GetAllOrderItems(c *gin.Context) {
	items, err := h.service.GetAllOrderItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить товары в заказах"})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetOrderItemByID
// @Summary      Получение товара в заказе по ID
// @Description  Возвращает товар по ID из заказа
// @Tags         Элемент заказа
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID товара в заказе"
// @Success      200  {object}  dto.OrderItemResponse
// @Failure      400  {object}  dto.ErrorResponse "Неверный ID товара в заказе"
// @Failure      404  {object}  dto.ErrorResponse "Товар в заказе не найден"
// @Failure      500  {object}  dto.ErrorResponse "Ошибка получения товара"
// @Router       /v1/order_items/{id} [get]
func (h *OrderItemHandler) GetOrderItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order item ID"})
		return
	}

	item, err := h.service.GetOrderItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateOrderItem
// @Summary      Создание товара в заказе
// @Description  Создает новый товар в заказе
// @Tags         Элемент заказа
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOrderItemRequest  true  "Данные товара"
// @Success      201  {object}  dto.OrderItemResponse
// @Failure      400  {object}  dto.ErrorResponse "Ошибка в данных товара"
// @Failure      500  {object}  dto.ErrorResponse "Ошибка создания товара в заказе"
// @Router       /v1/order_items [post]
func (h *OrderItemHandler) CreateOrderItem(c *gin.Context) {
	var req dto.CreateOrderItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.CreateOrderItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// UpdateOrderItem
// @Summary      Обновление товара в заказе
// @Description  Обновляет данные товара в заказе по его ID
// @Tags         Элемент заказа
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id      path      int                          true  "ID товара в заказе"
// @Param        request body     dto.UpdateOrderItemRequest    true  "Обновленные данные товара"
// @Success      200     {object}  dto.OrderItemResponse
// @Failure      400     {object}  dto.ErrorResponse "Ошибка в данных товара"
// @Failure      404     {object}  dto.ErrorResponse "Товар в заказе не найден"
// @Failure      500     {object}  dto.ErrorResponse "Ошибка обновления товара"
// @Router       /v1/order_items/{id} [put]
func (h *OrderItemHandler) UpdateOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order item ID"})
		return
	}

	var req dto.UpdateOrderItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.UpdateOrderItem(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteOrderItem
// @Summary      Удаление товара из заказа
// @Description  Удаляет товар из заказа по ID
// @Tags         Элемент заказа
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID товара в заказе"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  dto.ErrorResponse "Неверный ID товара в заказе"
// @Failure      500  {object}  dto.ErrorResponse "Ошибка удаления товара из заказа"
// @Router       /v1/order_items/{id} [delete]
func (h *OrderItemHandler) DeleteOrderItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order item ID"})
		return
	}

	if err := h.service.DeleteOrderItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}
