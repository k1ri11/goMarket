package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"sync"
	"time"
)

type TaskService struct {
	maxWorkers     int
	productService *ProductService
	taskQueue      chan *Task
	tasks          map[string]*Task
	mu             sync.Mutex
}

type TaskStatus string

const (
	StatusPending  TaskStatus = "Pending"
	StatusRunning  TaskStatus = "Running"
	StatusSuccess  TaskStatus = "Success"
	StatusFailed   TaskStatus = "Failed"
	StatusCanceled TaskStatus = "Canceled"
)

type Task struct {
	ID       string
	Status   TaskStatus
	FilePath string
	Error    error
	Cancel   context.CancelFunc
}

func NewTaskService(productService *ProductService) *TaskService {
	ts := &TaskService{
		maxWorkers:     5,
		productService: productService,
		taskQueue:      make(chan *Task, 5),
		tasks:          make(map[string]*Task),
	}

	for i := 0; i < ts.maxWorkers; i++ {
		go ts.runWorker()
	}

	return ts
}

func (ts *TaskService) runWorker() {
	for task := range ts.taskQueue {
		ts.updateTaskStatus(task.ID, StatusRunning, nil)
		ctx, cancel := context.WithCancel(context.Background())
		task.Cancel = cancel
		err := ts.exportProducts(ctx, task)
		if err != nil {
			task.Error = err
			if err.Error() != "task was canceled" {
				ts.updateTaskStatus(task.ID, StatusFailed, err)
			}
		} else {
			ts.updateTaskStatus(task.ID, StatusSuccess, nil)
		}
	}
}

func (ts *TaskService) updateTaskStatus(taskID string, status TaskStatus, err error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	fmt.Println(fmt.Sprintf("%s: status for %s changed %s -> %s",
		time.Now().Format("15:04:05"),
		taskID,
		ts.tasks[taskID].Status,
		status))

	if task, exists := ts.tasks[taskID]; exists {
		task.Status = status
		task.Error = err
	}
}

func (ts *TaskService) exportProducts(ctx context.Context, task *Task) error {
	products, err := ts.productService.GetAllProducts()
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("export/products_%s.json", task.ID)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	time.Sleep(10 * time.Second)

	select {
	case <-ctx.Done():
		return errors.New("task was canceled")
	default:
		if err := encoder.Encode(products); err != nil {
			return err
		}
	}

	task.FilePath = filePath
	return nil
}

func (ts *TaskService) CreateExportTask() string {
	taskID := uuid.New().String()
	task := &Task{
		ID:     taskID,
		Status: StatusPending,
	}

	ts.mu.Lock()
	ts.tasks[taskID] = task
	ts.mu.Unlock()

	ts.taskQueue <- task
	return taskID
}

func (ts *TaskService) GetTaskStatus(taskID string) (*Task, error) {

	if task, exists := ts.tasks[taskID]; exists {
		return task, nil
	}
	return nil, errors.New("task not found")
}

func (ts *TaskService) CancelTask(taskID string) error {

	if task, exists := ts.tasks[taskID]; exists {
		if task.Status == StatusRunning {
			task.Cancel()
			ts.updateTaskStatus(task.ID, StatusCanceled, nil)
			return nil
		}
		return errors.New("task not running")
	}

	return errors.New("task not found")
}
