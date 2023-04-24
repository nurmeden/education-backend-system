package http

import (
	"fmt"
	handler "task-backend/microservice1/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

// Server представляет HTTP-сервер
type Server struct {
	router  *gin.Engine
	handler *handler.StudentHandler
	// logger  logger.Logger
}

// NewServer создает и возвращает новый экземпляр HTTP-сервера
// func NewServer(handler *handler.StudentHandler, logger logger.Logger) *Server {
func NewServer(handler *handler.StudentHandler) *Server {
	router := gin.Default()

	// Зарегистрируйте маршруты и соответствующие обработчики здесь
	router.POST("/students", handler.CreateStudent)

	return &Server{
		router:  router,
		handler: handler,
		// logger:  logger,
	}
}

// Start запускает HTTP-сервер и начинает прослушивание входящих HTTP-запросов
func (s *Server) Start(addr string) error {
	// s.logger.Infof("Starting HTTP server on %s...", addr)
	err := s.router.Run(addr)
	if err != nil {
		return fmt.Errorf("failed to start HTTP server: %v", err)
	}
	return nil
}
