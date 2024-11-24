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

func (h *ShoppingCartHandler) GetAllShoppingCarts(c *gin.Context) {
	carts, err := h.service.GetAllShoppingCarts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve shopping carts"})
		return
	}

	c.JSON(http.StatusOK, carts)
}

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
