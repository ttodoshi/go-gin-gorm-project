package main

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/adapters/handler"
	"goGinGormProject/internal/adapters/repository/postgres"
	"goGinGormProject/internal/core/ports"
	"goGinGormProject/internal/core/servises"
	"goGinGormProject/pkg/env"
	"goGinGormProject/pkg/logging"
)

var (
	postService ports.PostService
	log         logging.Logger
)

func init() {
	env.LoadEnvVariables()
	log = logging.GetLogger()
}

// TODO: tests
func main() {
	postRepository := postgres.NewPostRepository()
	postService = servises.NewPostService(postRepository, log)
	initRoutes()
}

func initRoutes() {
	r := gin.Default()

	r.Use(handler.ErrorHandlerMiddleware())

	apiGroup := r.Group("/api")

	v1ApiGroup := apiGroup.Group("/v1")

	v1PostsGroup := v1ApiGroup.Group("/posts")
	{
		postHandler := handler.NewPostHandler(postService, log)
		v1PostsGroup.GET("/:uuid", postHandler.GetPostByUUID)
		v1PostsGroup.GET("/", postHandler.GetPosts)
		v1PostsGroup.POST("/", postHandler.CreatePost)
		v1PostsGroup.PUT("/:uuid", postHandler.UpdatePostByUUID)
		v1PostsGroup.DELETE("/:uuid", postHandler.DeletePostByUUID)
	}

	err := r.Run()
	if err != nil {
		log.Fatalf("error while running server")
	}
}
