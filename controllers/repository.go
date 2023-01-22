package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"shaw-and-partners-technical-challenge/domain"
	"shaw-and-partners-technical-challenge/handler"
)

func GetUserRepositories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if len(username) == 0 {
		handler.SendErrorResponse(w, fmt.Errorf("a username is expected"))
		return
	}

	url := domain.GlobalConfig.GithubApiUrl + "/users/" + username + "/repos"
	responseUrl, errGetReq := http.Get(url)
	if errGetReq != nil {
		fmt.Print(errGetReq.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(responseUrl.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject domain.Repository
	errUnmarshal := json.Unmarshal(responseData, &responseObject)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	response := handler.NewJsonResponse(responseObject, nil, handler.MetaResponse{
		Count: len(responseObject),
	})
	handler.SendJsonResponse(w, response, http.StatusOK)
}
