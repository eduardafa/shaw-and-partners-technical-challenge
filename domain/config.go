package domain

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	GithubApiUrl string
	Ctx          context.Context
	Port         string
}

var GlobalConfig AppConfig

// LoadVariables loads the values to GlobalConfig. When adding a new environment variable, load it inside this function
func (cfg *AppConfig) LoadVariables(envPath string) error {
	cfg.Ctx = context.Background()
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.GithubApiUrl = os.Getenv("GITHUB_API_URL")
	cfg.Port = os.Getenv("PORT")

	return nil
}
