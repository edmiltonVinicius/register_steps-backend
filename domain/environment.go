package domain

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Environment *globalEnv

type globalEnv struct {
	SERVER_PORT string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Environment = &globalEnv{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
    }

	time.Sleep(time.Second * 5)
}