package _tests

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"net/url"
	"shaw-and-partners-technical-challenge/controllers"
	"strings"
	"testing"
)

func TestUserControllers_Home(t *testing.T) {
	var home = struct {
		name               string
		method             string
		json               string
		param              string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{"home", "GET", "", "", controllers.Home, http.StatusOK}

	var req *http.Request
	if home.json != "" {
		req, _ = http.NewRequest(home.method, "/", nil)
	} else {
		req, _ = http.NewRequest(home.method, "/", strings.NewReader(home.json))
	}

	rr := httptest.NewRecorder()
	handler := home.handler
	handler.ServeHTTP(rr, req)

	if rr.Code != home.expectedStatusCode {
		t.Errorf("%s returned wrong status code: got %v want %v", home.name, rr.Code, home.expectedStatusCode)
	}
}

func TestUserControllers_GetUserDetails(t *testing.T) {
	var userDetails = struct {
		name               string
		method             string
		json               string
		param              string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{"user_details", "GET", "", "eduardafa", controllers.GetUserDetails, http.StatusOK}

	var req *http.Request
	if userDetails.json != "" {
		req, _ = http.NewRequest(userDetails.method, "/api/users/{username}/details", nil)
	} else {
		req, _ = http.NewRequest(userDetails.method, "/api/users/{username}/details", strings.NewReader(userDetails.json))
	}

	if userDetails.param != "" {
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("username", userDetails.param)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	}

	rr := httptest.NewRecorder()
	handler := userDetails.handler
	handler.ServeHTTP(rr, req)

	if rr.Code != userDetails.expectedStatusCode {
		t.Errorf("%s returned wrong status code: got %v want %v", userDetails.name, rr.Code, userDetails.expectedStatusCode)
	}
}

func TestUserControllers_GetUserRepositories(t *testing.T) {
	var userDetails = struct {
		name               string
		method             string
		json               string
		param              string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{"user_repos", "GET", "", "eduardafa", controllers.GetUserRepositories, http.StatusOK}

	var req *http.Request
	if userDetails.json != "" {
		req, _ = http.NewRequest(userDetails.method, "/api/users/{username}/repos", nil)
	} else {
		req, _ = http.NewRequest(userDetails.method, "/api/users/{username}/repos", strings.NewReader(userDetails.json))
	}

	if userDetails.param != "" {
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("username", userDetails.param)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	}

	rr := httptest.NewRecorder()
	handler := userDetails.handler
	handler.ServeHTTP(rr, req)

	if rr.Code != userDetails.expectedStatusCode {
		t.Errorf("%s returned wrong status code: got %v want %v", userDetails.name, rr.Code, userDetails.expectedStatusCode)
	}
}

func TestUserControllers_GetUsers(t *testing.T) {
	var users = struct {
		name               string
		method             string
		json               string
		param              string
		handler            http.HandlerFunc
		expectedStatusCode int
	}{"user_repos", "GET", "", "83236123", controllers.GetAllUsers, http.StatusOK}

	var req *http.Request
	queryParams := url.Values{}
	queryParams.Add("since", users.param)
	if users.json != "" {
		req, _ = http.NewRequest(users.method, "/api/users?"+queryParams.Encode(), nil)
	} else {
		req, _ = http.NewRequest(users.method, "/api/users?"+queryParams.Encode(), strings.NewReader(users.json))
	}

	rr := httptest.NewRecorder()
	handler := users.handler
	handler.ServeHTTP(rr, req)

	if rr.Code != users.expectedStatusCode {
		t.Errorf("%s returned wrong status code: got %v want %v", users.name, rr.Code, users.expectedStatusCode)
	}
}
