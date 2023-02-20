package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shereifsrf/savespent-api/controller"
	"github.com/shereifsrf/savespent-api/dao"
)

var Router *gin.Engine

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dao.Initialise()

	r := setupRouter()
	// set PORT
	r.Run(":" + os.Getenv("PORT"))
}

// setup router
func setupRouter() *gin.Engine {
	r := gin.Default()
	// get context
	ctx := context.Background()

	// v1 API
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// status
	controller.RegisterStatusRoutes(v1)
	// session
	controller.RegisterSessionRoutes(ctx, v1)

	return r
}
