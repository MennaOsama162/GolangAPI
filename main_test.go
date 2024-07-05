package main

import (
	"bytes"
	"encoding/json"
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
	payload := `{"name": "Test Author", "email": "test.author@example.com"}`
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

func TestCreateBook(t *testing.T) {
	app := setupApp()

	// Create an author first
	authorPayload := `{"name": "Test Author", "email": "test.author@example.com"}`
	authorReq := httptest.NewRequest(http.MethodPost, "/api/authors", bytes.NewBuffer([]byte(authorPayload)))
	authorReq.Header.Set("Content-Type", "application/json")
	authorResp, err := app.Test(authorReq, -1)
	if err != nil {
		log.Fatalf("Could not execute request: %v", err)
	}
	assert.Equal(t, http.StatusCreated, authorResp.StatusCode)

	// Parse the author ID from the response
	var authorResponse map[string]interface{}
	body, _ := ioutil.ReadAll(authorResp.Body)
	json.Unmarshal(body, &authorResponse)
	authorID := int(authorResponse["id"].(float64))

	// Create a book with the author ID
	bookPayload := map[string]interface{}{
		"title":          "Test Book",
		"isbn":           "1234567890",
		"published_date": "2023-01-01",
		"author_id":      authorID,
	}
	bookPayloadBytes, _ := json.Marshal(bookPayload)
	req := httptest.NewRequest(http.MethodPost, "/api/books", bytes.NewBuffer(bookPayloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		log.Fatalf("Could not execute request: %v", err)
	}

	body, _ = ioutil.ReadAll(resp.Body)
	log.Printf("Response: %s", string(body))

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestCreateDuplicateAuthor(t *testing.T) {
	app := setupApp()
	// cleanUp()

	// Create the first author
	payload1 := `{"name": "Test Author", "email": "test.author@example.com"}`
	req1 := httptest.NewRequest(http.MethodPost, "/api/authors", bytes.NewBuffer([]byte(payload1)))
	req1.Header.Set("Content-Type", "application/json")
	resp1, err := app.Test(req1, -1)
	if err != nil {
		log.Fatalf("Could not execute request: %v", err)
	}
	assert.Equal(t, http.StatusCreated, resp1.StatusCode)

	// Create a duplicate author
	payload2 := `{"name": "Test Author Duplicate", "email": "test.author@example.com"}`
	req2 := httptest.NewRequest(http.MethodPost, "/api/authors", bytes.NewBuffer([]byte(payload2)))
	req2.Header.Set("Content-Type", "application/json")
	resp2, err := app.Test(req2, -1)
	if err != nil {
		log.Fatalf("Could not execute request: %v", err)
	}

	body, _ := ioutil.ReadAll(resp2.Body)
	log.Printf("Response: %s", string(body))

	assert.Equal(t, http.StatusConflict, resp2.StatusCode)
}
