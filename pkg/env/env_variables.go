package env

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load(".env.docker")
		if err != nil {
			log.Fatal("error loading .env file")
		}
	}
	err = godotenv.Overload(".env.local")
	if err != nil {
		recover()
	}
}
