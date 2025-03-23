package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebbeliaev/to-do-list/internal/domain/models"
	"github.com/google/uuid"
)

func (s *ServerApi) getTasks(c *gin.Context) {
	tasks, err := s.Repo.GetTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)

}

func (s *ServerApi) createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.TID = uuid.New().String()
	task.Status = "New"

	if err := s.Repo.SaveTask(task); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
