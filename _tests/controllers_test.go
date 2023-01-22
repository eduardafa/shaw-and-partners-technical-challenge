package _tests

import (
	"net/http"
	"net/http/httptest"
	"shaw-and-partners-technical-challenge/cmd"
	"testing"
)

func Test_user_controllers(t *testing.T) {
	var userDetails = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"user", "/api/users/eduardafa/details", http.StatusOK},
		{"404", "/api/fish", http.StatusNotFound},
	}

	var app cmd.Application
	routes := app.Routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// range through test data
	for _, entry := range userDetails {
		resp, err := ts.Client().Get(ts.URL + entry.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != entry.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got %d", entry.name, entry.expectedStatusCode, resp.StatusCode)
		}
	}
}
