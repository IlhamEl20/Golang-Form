package main

import (
	"Si-KP/handler"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAppName(t *testing.T) {
	app := fiber.New()

	app.Get("/", handler.AppName)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to send test request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	expectedBody := `{"app_name":"boilerplate go","desc":"boilerplate go"}`

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if string(body) != expectedBody {
		t.Errorf("Expected response body '%s' but got '%s'", expectedBody, string(body))
	}
}
