package database

import (
	"goGinGormProject/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var db *gorm.DB
var once sync.Once

func ConnectToDb() *gorm.DB {
	once.Do(func() {
		dbUrl := os.Getenv("DB_URL")

		var err error

		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to database")
		}
		err = db.AutoMigrate(&domain.Post{})
		if err != nil {
			log.Fatal("error while migrating")
		}
	})
	return db
}
