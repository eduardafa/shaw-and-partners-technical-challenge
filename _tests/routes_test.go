package _tests

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"shaw-and-partners-technical-challenge/cmd"
	"strings"
	"testing"
)

func Test_app_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/api/*", "GET"},
		//{"/api/users", "GET"},
		//{"/api/users/{username}/details", "GET"},
		//{"/api/users/{username}/repos", "GET"},
	}

	var app cmd.Application
	mux := app.Routes()

	chiRoutes := mux.(chi.Routes)
	for _, route := range registered {
		// check to see if the route exists
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registered", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}
