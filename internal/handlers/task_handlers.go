package handlers

import (
	"net/http"

	"github.com/999Marv/task-manager-go/internal/models"
	"github.com/999Marv/task-manager-go/internal/repository"
	"github.com/gin-gonic/gin"
)

var taskRepo = repository.NewTaskRepository()

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindBodyWithJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask := taskRepo.Create(newTask)
	c.JSON(http.StatusCreated, createdTask)
}

func GetAllTasks(c *gin.Context) {
	tasks := taskRepo.GetAll()
	c.JSON(http.StatusOK, tasks)
}
