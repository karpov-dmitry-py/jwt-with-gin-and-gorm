package api

import (
	"github.com/gin-gonic/gin"
	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/pkg/service"
)

type Server struct {
	Router      *gin.Engine
	UserService service.UserService
}

// NewServer inits a new server
func NewServer(router *gin.Engine, userService service.UserService) Server {
	server := Server{
		Router:      router,
		UserService: userService,
	}

	server.setRouting()

	return server
}

func (s *Server) setRouting() {
	s.Router.GET("/alive", s.healthCheck)
	s.Router.POST("/signup", s.signUp)
}
