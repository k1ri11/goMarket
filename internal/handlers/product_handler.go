package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetProducts возвращает список всех продуктов с применением фильтрации.
// @Summary Получение списка всех продуктов
// @Description Возвращает список всех продуктов с возможностью фильтрации по параметрам.
// @Tags Продукт
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param filter query dto.ProductFilterRequest false "Фильтр для поиска продуктов"
// @Success 200 {array} dto.ProductResponsePagination "Список продуктов"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID продукта"
// @Failure 422 {object} dto.ErrorResponse  "Ошибка валидации"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var filter dto.ProductFilterRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		var errorMessages []string
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			for _, e := range errs {
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation, condition: %s", e.Field(), e.ActualTag()))
			}
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Validation failed",
			"details": errorMessages,
			"code":    http.StatusUnprocessableEntity,
		})
		return
	}

	products, err := h.service.GetFilteredProducts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID возвращает продукт по ID.
// @Summary Получение продукта по ID
// @Description Возвращает продукт, соответствующий переданному ID.
// @Tags Продукт
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} dto.ProductResponse "Продукт"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID продукта"
// @Failure 404 {object} dto.ErrorResponse "Продукт не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct добавляет новый продукт.
// @Summary Добавление нового продукта
// @Description Добавляет новый продукт в базу данных.
// @Tags Продукт
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Данные нового продукта"
// @Success 201 {object} dto.ProductResponse "Созданный продукт"
// @Failure 400 {object} dto.ErrorResponse "Некорректные данные запроса"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct обновляет существующий продукт.
// @Summary Обновление существующего продукта
// @Description Обновляет данные продукта по переданному ID.
// @Tags Продукт
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Param product body dto.UpdateProductRequest true "Данные для обновления продукта"
// @Success 200 {object} dto.ProductResponse "Обновленный продукт"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID продукта"
// @Failure 404 {object} dto.ErrorResponse "Продукт не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.service.UpdateProduct(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct удаляет продукт по ID.
// @Summary Удаление продукта
// @Description Удаляет продукт из базы данных по ID.
// @Tags Продукт
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]interface{} "Продукт удален успешно"
// @Failure 400 {object} dto.ErrorResponse "Неверный ID продукта"
// @Failure 404 {object} dto.ErrorResponse "Продукт не найден"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Router /v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
