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

func (h *TaskHandler) CreateTask(c *gin.Context) {
	taskID := h.taskService.CreateExportTask()
	c.JSON(http.StatusAccepted, gin.H{"task_id": taskID})
}

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

func (h *TaskHandler) CancelTask(c *gin.Context) {
	taskID := c.Param("id")
	err := h.taskService.CancelTask(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task canceled"})
}
