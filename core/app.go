package core

import (
	"log"
	"shaw-and-partners-technical-challenge/domain"
)

func LoadEnv() {
	log.Println("Loading .env file")
	err := domain.GlobalConfig.LoadVariables(domain.Utils{}.RootPath() + "/config/.env")
	if err != nil {
		log.Printf("Can't load .env")
		log.Panic(err)
	}
	log.Println(".env file loaded successfully")
}
