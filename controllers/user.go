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

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	url := domain.GlobalConfig.GithubApiUrl + "/users/" + username
	responseUrl, errGetReq := http.Get(url)
	if errGetReq != nil {
		fmt.Print(errGetReq.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(responseUrl.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject domain.User
	errUnmarshal := json.Unmarshal(responseData, &responseObject)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	response := handler.NewJsonResponse(responseObject, nil, handler.MetaResponse{
		Count: 1,
	})
	handler.SendJsonResponse(w, response, http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	number := query.Get("since")

	url := domain.GlobalConfig.GithubApiUrl + "/users?since=" + number
	responseUrl, errGetReq := http.Get(url)
	if errGetReq != nil {
		fmt.Print(errGetReq.Error())
		os.Exit(1)
	}
	link := responseUrl.Header.Get("link")

	responseData, err := io.ReadAll(responseUrl.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []domain.SummarizedUserData
	errUnmarshal := json.Unmarshal(responseData, &responseObject)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	response := handler.NewJsonResponse(responseObject, nil, handler.MetaResponseWithPagination{
		Count: len(responseObject),
		Link:  link,
	})
	handler.SendJsonResponse(w, response, http.StatusOK)
}
