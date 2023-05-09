package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/pkg/service"
)

func (s *Server) signUp(c *gin.Context) {
	var (
		user service.SignUpUser
		err  error
	)

	if err = c.Bind(&user); err != nil || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid req body"})
		return
	}

	userID, err := s.UserService.SignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": userID})
}
