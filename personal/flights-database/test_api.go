package main

import (
	"bytes"
	"encoding/json"
	"flights-database/api"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	// Test GET endpoint
	fmt.Println("Testing GET endpoint...")
	testGetEndpoint()

	// Test POST endpoint
	fmt.Println("\nTesting POST endpoint...")
	testPostEndpoint()
}

func testGetEndpoint() {
	req, _ := http.NewRequest("GET", "/api/aircraft", nil)
	w := httptest.NewRecorder()

	api.Handler(w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Response: %s\n", w.Body.String())
}

func testPostEndpoint() {
	requestBody := map[string][]string{
		"flightCodes": {"IC5302", "AI101", "6E1234"},
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/api/aircraft", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	api.Handler(w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Response: %s\n", w.Body.String())
}
