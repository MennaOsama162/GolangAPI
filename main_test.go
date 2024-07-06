package main

import (
	"bytes"
	// "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"library-management/pkg/config"
	"library-management/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()
	config.ConnectDB()
	routes.SetupRoutes(app)
	return app
}

func TestGetBooks(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		log.Fatalf("Could not execute request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateAuthor(t *testing.T) {
	app := setupApp()
	payload := `{"name": "Test34 Author", "email": "test34.author@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/authors", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		log.Fatalf("Could not execute request: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Response: %s", string(body))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
