package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goMarket/internal/services"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// CreateTask
// @Summary      Создание новой задачи экспорта
// @Description  Создает новую задачу для обработки экспорта и возвращает ID задачи
// @Tags         Задачи
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Success      202  {object}  map[string]interface{}
// @Failure      500  {object}  dto.ErrorResponse "Ошибка создания задачи"
// @Router       /v1/tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	taskID := h.taskService.CreateExportTask()
	c.JSON(http.StatusAccepted, gin.H{"task_id": taskID})
}

// GetTaskStatus
// @Summary      Получение статуса задачи по ID
// @Description  Получает статус, путь к файлу и возможные ошибки для задачи по ее ID
// @Tags         Задачи
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID задачи"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  dto.ErrorResponse "Неверный ID задачи"
// @Failure      404  {object}  dto.ErrorResponse "Задача не найдена"
// @Failure      500  {object}  dto.ErrorResponse "Ошибка получения статуса задачи"
// @Router       /v1/tasks/{id} [get]
func (h *TaskHandler) GetTaskStatus(c *gin.Context) {
	taskID := c.Param("id")
	task, err := h.taskService.GetTaskStatus(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task_id": task.ID,
		"status":  task.Status,
		"path":    task.FilePath,
		"error":   task.Error,
	})
}

// CancelTask
// @Summary      Отмена задачи по ID
// @Description  Отменяет задачу по ID и прекращает ее обработку
// @Tags         Задачи
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID задачи"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  dto.ErrorResponse "Неверный ID задачи"
// @Failure      404  {object}  dto.ErrorResponse "Задача не найдена"
// @Failure      500  {object}  dto.ErrorResponse "Ошибка отмены задачи"
// @Router       /v1/tasks/{id} [delete]
func (h *TaskHandler) CancelTask(c *gin.Context) {
	taskID := c.Param("id")
	err := h.taskService.CancelTask(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "задача отменена"})
}
