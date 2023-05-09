package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) healthCheck(c *gin.Context) {
	// nolint:gomnd
	c.JSON(200, gin.H{"status": "ok"})
}
