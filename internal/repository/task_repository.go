package repository

import (
	"sync"
	"time"

	"github.com/999Marv/task-manager-go/internal/models"
)

type TaskRepository struct {
	tasks  map[int]models.Task
	mu     sync.RWMutex
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (r *TaskRepository) Create(task models.Task) models.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.nextID++
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = task
	return task
}

func (r *TaskRepository) GetAll() []models.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]models.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}
