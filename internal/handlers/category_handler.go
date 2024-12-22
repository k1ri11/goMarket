package handlers

import (
	"github.com/gin-gonic/gin"
	"goMarket/internal/dto"
	"goMarket/internal/services"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetAllCategories возвращает список всех категорий.
// @Summary Get all categories
// @Description Получение списка всех категорий
// @Tags Категория
// @Produce json
// @Success 200 {array} dto.CategoryResponse "Список категорий"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID возвращает категорию по ID.
// @Summary Get category by ID
// @Description Получение категории по ID
// @Tags Категория
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryResponse "Категория"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID"
// @Failure 404 {object} dto.ErrorResponse "Категория не найдена"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// CreateCategory создает новую категорию.
// @Summary Create a new category
// @Description Создание новой категории
// @Tags Категория
// @Accept json
// @Produce json
// @Param input body dto.CreateCategoryRequest true "Данные для создания категории"
// @Success 201 {object} dto.CategoryResponse "Созданная категория"
// @Failure 400 {object} dto.ErrorResponse "Некорректные данные"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCategory обновляет существующую категорию.
// @Summary Update an existing category
// @Description Обновление категории по ID
// @Tags Категория
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param input body dto.UpdateCategoryRequest true "Данные для обновления категории"
// @Success 200 {object} dto.CategoryResponse "Обновленная категория"
// @Failure 400 {object} dto.ErrorResponse "Некорректные данные или ID"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.service.UpdateCategory(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory удаляет категорию по ID.
// @Summary Delete a category
// @Description Удаление категории по ID
// @Tags Категория
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
