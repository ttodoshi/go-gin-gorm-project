package postgres

import (
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
		dsn := os.Getenv("DB_URL")

		var err error

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to database")
		}
	})
	return db
}
