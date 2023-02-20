package controller

import "github.com/gin-gonic/gin"

type StatusController interface {
	GetStatus(c *gin.Context)
}

type statusController struct {
}

func NewStatusController() StatusController {
	return &statusController{}
}

func (s *statusController) GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

// register status routes
func RegisterStatusRoutes(r *gin.RouterGroup) {
	status := NewStatusController()
	r.GET("/status", status.GetStatus)
}
