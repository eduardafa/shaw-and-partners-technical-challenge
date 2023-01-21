package main

import (
	"fmt"
	"shaw-and-partners-technical-challenge/core"
	"shaw-and-partners-technical-challenge/routes"
)

func main() {
	fmt.Println("Starting the server with GoLang")
	core.LoadEnv()
	routes.HandleRequest()
}
