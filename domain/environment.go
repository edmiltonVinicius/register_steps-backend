package domain

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/edmiltonVinicius/register-steps/api/utils"
	"github.com/joho/godotenv"
)

const (
	DEV  = "development"
	TEST = "test"
	PROD = "production"
)

var Environment *globalEnv

type globalEnv struct {
	ENV            string
	SERVER_PORT    string
	DB_HOST        string
	DB_PORT        string
	DB_USER        string
	DB_PASSWORD    string
	DB_NAME        string
	CTX            context.Context
	REDIS_URL      string
	REDIS_PASSWORD string
	REDIS_USERNAME string
	REDIS_DB       string
}

func LoadEnv(environment string) {
	var env string

	if environment == TEST {
		root := utils.GetRootPath()
		env = string(root) + "/.env.test"
	} else {
		env = ".env"
	}

	err := godotenv.Load(env)

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Environment = &globalEnv{
		ENV:            os.Getenv("ENV"),
		SERVER_PORT:    os.Getenv("SERVER_PORT"),
		DB_HOST:        os.Getenv("DB_HOST"),
		DB_PORT:        os.Getenv("DB_PORT"),
		DB_USER:        os.Getenv("DB_USER"),
		DB_PASSWORD:    os.Getenv("DB_PASSWORD"),
		DB_NAME:        os.Getenv("DB_NAME"),
		REDIS_URL:      os.Getenv("REDIS_URL"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		REDIS_USERNAME: os.Getenv("REDIS_USER"),
		REDIS_DB:       os.Getenv("REDIS_DB"),
		CTX:            context.Background(),
	}

	time.Sleep(time.Second * 5)
}
