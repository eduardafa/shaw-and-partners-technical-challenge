package controllers

import (
	"fmt"
	"net/http"
	"shaw-and-partners-technical-challenge/handler"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w,
		"Welcome to this API!\n\n"+
			"The following endpoints are available for tests:\n"+
			"GET - /api/users?since={number} - Returns a list of GitHub users with an ID greater than the defined ID\n"+
			"GET - /api/users/{username}/details - Returns the details of a GitHub user\n"+
			"GET - /api/users/{username}/repos - Returns a list with all GitHub user repositories\n")
	if err != nil {
		handler.SendErrorResponse(w, err)
		return
	}
}
