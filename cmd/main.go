package cmd

import (
	"fmt"
	"log"
	"net/http"
	"shaw-and-partners-technical-challenge/core"
	"shaw-and-partners-technical-challenge/domain"
)

type Application struct{}

func main() {
	app := Application{}

	mux := app.Routes()

	fmt.Println("Starting the server with GoLang")
	core.LoadEnv()
	err := http.ListenAndServe(":"+domain.GlobalConfig.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
