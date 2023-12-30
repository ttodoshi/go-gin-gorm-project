package main

import (
	"goGinGormExample/internal/models"
	"goGinGormExample/pkg/initializers"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatalf("error while migrating")
	}
}
