package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shereifsrf/savespent-api/controller"
)

func main() {
	r := gin.Default()

	r.GET("/get-count", controller.GetCount)
	r.POST("/session", controller.SessionPersist)

	r.Run()
}
