package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebbeliaev/to-do-list/internal/config"
	"github.com/glebbeliaev/to-do-list/internal/domain/models"
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	GetTasks() ([]models.Task, error)
	GetTask(string) (models.Task, error)
	SaveTask(models.Task) error
	UpdateTask(models.Task) error
	DeleteTask(string) error
}

type ServerApi struct {
	server *http.Server
	valid  *validator.Validate
	Repo   Repository
}

func New(cfg config.Config, repo Repository) *ServerApi {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	return &ServerApi{
		server: &server,
		valid:  validator.New(),
		Repo:   repo,
	}
}

func (s *ServerApi) Start() error {
	s.configRoutes()
	return s.server.ListenAndServe()
}

func (s *ServerApi) configRoutes() {

	router := gin.Default()
	router.GET("/tasks", s.getTasks)
	router.POST("/tasks", s.createTask)
	task := router.Group("/task")
	{
		task.PUT("/:tid", func(c *gin.Context) {})
		task.DELETE("/:tid", func(c *gin.Context) {})
		task.GET("/:tid", func(c *gin.Context) {})
	}
	s.server.Handler = router
}
