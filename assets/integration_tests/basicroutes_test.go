// +build integration

package integration_tests

import (
	"net/http"
	"testing"
)

func TestBaseRouteGet200(t *testing.T) {
	resp, _ := http.Get("http://localhost:8080/")

	if resp.StatusCode != 200 {
		t.Error("Expected 200 response")
	}
}

func TestHealthCheckGet200(t *testing.T) {
	resp, _ := http.Get("http://localhost:8080/health-check/")

	if resp.StatusCode != 200 {
		t.Error("Expected 200 response")
	}
}
