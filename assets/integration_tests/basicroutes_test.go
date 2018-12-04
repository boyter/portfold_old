// +build integration

package integration_tests

import (
	"net/http"
	"testing"
)


var mainRoute = "http://localhost:8080/"
var routes200 = []string{
	"",
	"health-check/",
	"static/",
	"static/css/main.css",
}

func TestRoute200(t *testing.T) {
	for _, route := range routes200 {
		resp, _ := http.Get(mainRoute + route)

		if resp.StatusCode != 200 {
			t.Error("Expected 200 response for", route)
		}
	}
}
