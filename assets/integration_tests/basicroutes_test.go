// +build integration

package integration_tests

import (
	"bytes"
	"net/http"
	"testing"
)


var mainRoute = "http://localhost:8080/"
var routes200 = []string{
	"",
	"health-check/",
	"user/login/",
	"static/",
	"static/css/main.css",
	"static/img/favicon.ico",
}

func TestGetRoute200(t *testing.T) {
	for _, route := range routes200 {
		resp, _ := http.Get(mainRoute + route)

		if resp.StatusCode != 200 {
			t.Error("Expected 200 response for", route, "got", resp.StatusCode)
		}
	}
}

var routes405 = []string{
	"",
	"health-check/",
}

func TestPostRoute405(t *testing.T) {
	for _, route := range routes405 {
		resp, _ := http.Post(mainRoute + route, "text/html", bytes.NewBuffer([]byte("")))

		if resp.StatusCode != 405 {
			t.Error("Expected 405 response for", route, "got", resp.StatusCode)
		}
	}
}

var routes404 = []string{
	"something/that/should/not/exist/",
	"static/css/notreal.css",
	"static/img/random_image.jpg",
}

func TestRoute404(t *testing.T) {
	for _, route := range routes404 {
		resp, _ := http.Get(mainRoute + route)

		if resp.StatusCode != 404 {
			t.Error("Expected 404 response for", route)
		}
	}
}
