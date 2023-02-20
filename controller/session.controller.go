// Date: 2018-01-31 15:00
// sessoin.go handles user session such as adding user ID and email
package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shereifsrf/savespent-api/service"
)

type SessionController interface {
	AddSession(c *gin.Context)
}

type sessionController struct {
	sessionService service.SessionService
}

// get new session controller
func NewSessionController(ctx context.Context) SessionController {
	ss := service.NewSessionService(ctx)

	return &sessionController{
		sessionService: ss,
	}
}

// AddSession adds user session
func (s *sessionController) AddSession(c *gin.Context) {
	// validate if the body exist with values
	var bSession service.BindedSession
	if err := c.ShouldBindJSON(&bSession); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// add user session
	user, err := s.sessionService.AddSession(&bSession)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

// register session routes
func RegisterSessionRoutes(ctx context.Context, r *gin.RouterGroup) {
	session := NewSessionController(ctx)

	r.POST("/session", session.AddSession)
}
