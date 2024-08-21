package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ServerPort string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ServerPort = os.Getenv("SERVER_PORT")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
