package main

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/controllers"
	"goGinGormProject/pkg/initializers"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

// TODO: logging, hexagonal, error handling, uuid
func main() {
	r := gin.Default()

	r.GET("/api/v1/posts/:id", controllers.GetPostByID)
	r.GET("/api/v1/posts", controllers.GetPosts)
	r.POST("/api/v1/posts", controllers.CreatePost)
	r.PUT("/api/v1/posts/:id", controllers.UpdatePost)
	r.DELETE("/api/v1/posts/:id", controllers.DeletePost)

	err := r.Run()
	if err != nil {
		log.Fatalf("error while running server")
	}
}
