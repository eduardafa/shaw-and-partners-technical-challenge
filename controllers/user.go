package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ApiUrl = os.Getenv("GITHUB_API_URL")

func GetUserDetails(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	url := fmt.Sprintf(ApiUrl+"/users/%s", username)
	//response, err := http.Get(ApiUrl + "/users/" + username)
	response, err := http.Get(url)
	fmt.Println("response>>>>>>>", response)
	if err != nil {
		fmt.Print(err.Error())
		//os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

//func GetPaintCansByRoom(w http.ResponseWriter, r *http.Request) {
//	roomModel := models.Room{}
//	decoder := json.NewDecoder(r.Body)
//	decoder.DisallowUnknownFields()
//	errDecoder := decoder.Decode(&roomModel)
//	if errDecoder != nil {
//		handler.SendErrorResponse(w, &handler.ApiError{
//			Detail: errDecoder.Error(),
//			Status: http.StatusBadRequest,
//		})
//		return
//	}
//
//	errCalculate, paintCansQuantities := models.CalculateAreaToBePaintedByRoom(roomModel)
//	if errCalculate != nil {
//		handler.SendErrorResponse(w, errCalculate)
//		return
//	}
//	response := handler.NewJsonResponse(paintCansQuantities, nil, handler.MetaResponse{
//		Count: 1,
//	})
//	handler.SendJsonResponse(w, response, http.StatusOK)
//}
