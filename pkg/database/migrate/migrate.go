package main

import (
	"goGinGormProject/internal/core/domain"
	"goGinGormProject/pkg/database/postgres"
	"goGinGormProject/pkg/env"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func init() {
	env.LoadEnvVariables()
	db = postgres.ConnectToDb()
}

func main() {
	err := db.AutoMigrate(&domain.Post{})
	if err != nil {
		log.Fatalf("error while migrating")
	}
}
