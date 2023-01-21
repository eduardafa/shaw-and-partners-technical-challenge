package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Meta   interface{} `json:"meta"`
	Data   interface{} `json:"data"`
	Errors []ApiError  `json:"errors"`
}

type MetaResponse struct {
	Count int `json:"count"`
}

type MetaResponseWithPagination struct {
	Count int    `json:"count"`
	Link  string `json:"link_for_the_next_page"`
}

var debugApi = false

func SetDebug(setDebug bool) {
	debugApi = setDebug
}

// SendJsonResponse sends a json response from received data, status code and prints current time to standard logger.
// By default, ApiError is not sent with debug info. To send responses with debug, call handler.SetDebug(true)
// To log error details to NewRelic, use SendJsonResponseContext instead
func SendJsonResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	response = prepareResponse(response)

	sendJson(w, response, statusCode)
}

func sendJson(w http.ResponseWriter, response interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Print("| STARTED | status | ", http.StatusInternalServerError, " |")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Print("| STARTED | status | ", statusCode)
	}
}

func prepareResponse(response interface{}) interface{} {
	preparedResponse := response

	jsonResponse, ok := response.(JsonResponse)
	if ok {
		if !debugApi {
			for i, apiError := range jsonResponse.Errors {
				apiError.Debug = nil
				jsonResponse.Errors[i] = apiError
			}
			preparedResponse = jsonResponse
		}
	}

	return preparedResponse
}

// SendErrorResponse sends an error response, in case the response is from type ApiError it returns
// the defined status at the struct, otherwise it sends an answer with status http.StatusInternalServerError
func SendErrorResponse(w http.ResponseWriter, err error) {
	errResponse, status := formatErrResponse(err)
	log.Print("| ERROR | ", err.Error())
	SendJsonResponse(w, errResponse, status)
}

func formatErrResponse(err error) (res JsonResponse, status int) {
	var errors []ApiError
	apiError, ok := err.(*ApiError)

	if ok {
		errors = []ApiError{*apiError}
		status = apiError.Status
	} else {
		errors = []ApiError{{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}}
		status = http.StatusInternalServerError
	}

	return NewJsonResponse(nil, errors, MetaResponse{
		Count: len(errors),
	}), status
}

// NewJsonResponse Defines a standard struct for sending responses, populating the JsonResponse with data, errors and meta received
func NewJsonResponse(data interface{}, errors []ApiError, meta interface{}) JsonResponse {
	if errors == nil {
		errors = []ApiError{}
	}

	if data == nil {
		data = []interface{}{}
	}

	response := JsonResponse{
		Data:   data,
		Meta:   meta,
		Errors: errors,
	}

	return response
}
