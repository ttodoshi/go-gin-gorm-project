package main

import (
	"github.com/gin-gonic/gin"
	"goGinGormProject/internal/adapters/handler"
	"goGinGormProject/internal/adapters/repository"
	"goGinGormProject/internal/core/ports"
	"goGinGormProject/internal/core/servises"
	"goGinGormProject/pkg/database/postgres"
	"goGinGormProject/pkg/env"
	"gorm.io/gorm"
	"log"
)

var (
	db          *gorm.DB
	postService ports.PostService
)

func init() {
	env.LoadEnvVariables()
	db = postgres.ConnectToDb()
}

// TODO: logging, error handling, dto, tests
func main() {
	postRepository := repository.NewPostRepository(db)
	postService = servises.NewPostService(postRepository)
	initRoutes()
}

func initRoutes() {
	r := gin.Default()

	postHandler := handler.NewPostHandler(postService)
	r.GET("/api/v1/posts/:uuid", postHandler.GetPostByUUID)
	r.GET("/api/v1/posts", postHandler.GetPosts)
	r.POST("/api/v1/posts", postHandler.CreatePost)
	r.PUT("/api/v1/posts/:uuid", postHandler.UpdatePostByUUID)
	r.DELETE("/api/v1/posts/:uuid", postHandler.DeletePostByUUID)

	err := r.Run()
	if err != nil {
		log.Fatalf("error while running server")
	}
}
