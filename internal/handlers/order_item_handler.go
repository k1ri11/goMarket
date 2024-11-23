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

func (h *OrderItemHandler) GetAllOrderItems(c *gin.Context) {
	items, err := h.service.GetAllOrderItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

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
