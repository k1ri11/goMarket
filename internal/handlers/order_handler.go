package handlers

import (
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// GetAllOrders возвращает список всех заказов.
// @Summary Получить все заказы
// @Description Возвращает список всех заказов
// @Tags Заказ
// @Produce json
// @Success 200 {array} dto.OrderResponse "Список заказов"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/orders [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrderByID возвращает заказ по ID.
// @Summary Получить заказ по ID
// @Description Возвращает данные заказа по его идентификатору
// @Tags Заказ
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} dto.OrderResponse "Данные заказа"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID заказа"
// @Failure 404 {object} dto.ErrorResponse "Заказ не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.service.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateOrder создает новый заказ.
// @Summary Создать заказ
// @Description Создает новый заказ
// @Tags Заказ
// @Accept json
// @Produce json
// @Param request body dto.CreateOrderRequest true "Данные для создания заказа"
// @Success 201 {object} dto.OrderResponse "Созданный заказ"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации запроса"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// UpdateOrder обновляет существующий заказ.
// @Summary Обновить заказ
// @Description Обновляет данные существующего заказа
// @Tags Заказ
// @Accept json
// @Produce json
// @Param id path int true "ID заказа"
// @Param request body dto.UpdateOrderRequest true "Данные для обновления заказа"
// @Success 200 {object} dto.OrderResponse "Обновленный заказ"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации запроса или ID заказа"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var req dto.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.UpdateOrder(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrder удаляет заказ по ID.
// @Summary Удалить заказ
// @Description Удаляет заказ по его идентификатору
// @Tags Заказ
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} map[string]string "Сообщение об успешном удалении"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID заказа"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := h.service.DeleteOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
