package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// AircraftMapping represents the structure of aircraft data
type AircraftMapping struct {
	AircraftType string `json:"aircraftType"`
	DisplayName  string `json:"displayName"`
	Manufacturer string `json:"manufacturer"`
}

// AircraftDatabase represents the complete database structure
type AircraftDatabase struct {
	AircraftMappings map[string]AircraftMapping `json:"aircraft_mappings"`
}

// AircraftRequest represents the incoming request payload
type AircraftRequest struct {
	FlightCodes []string `json:"flightCodes"`
}

// AircraftResponse represents the response payload
type AircraftResponse struct {
	Results map[string]AircraftMapping `json:"results"`
}

// Handler handles the aircraft mapping API requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET and POST methods
	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Load the database
	database, err := loadDatabase()
	if err != nil {
		http.Error(w, "Failed to load database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "GET" {
		// For GET requests, return all aircraft mappings
		response := AircraftResponse{
			Results: database.AircraftMappings,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Handle POST request
	var request AircraftRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Filter results based on requested flight codes
	results := make(map[string]AircraftMapping)
	for _, flightCode := range request.FlightCodes {
		if aircraft, exists := database.AircraftMappings[flightCode]; exists {
			results[flightCode] = aircraft
		}
	}

	response := AircraftResponse{
		Results: results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// loadDatabase loads the JSON database file
func loadDatabase() (*AircraftDatabase, error) {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	// Try multiple possible paths for the database file
	possiblePaths := []string{
		"data/aircraft_mappings.json",
		filepath.Join(cwd, "data/aircraft_mappings.json"),
		"./data/aircraft_mappings.json",
		"../data/aircraft_mappings.json",
		"../../data/aircraft_mappings.json",
		"/var/task/data/aircraft_mappings.json",
		filepath.Join(cwd, "..", "data/aircraft_mappings.json"),
		filepath.Join(cwd, "..", "..", "data/aircraft_mappings.json"),
	}

	var dataPath string
	var data []byte
	var readErr error

	// Try each possible path
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			data, readErr = ioutil.ReadFile(path)
			if readErr == nil {
				dataPath = path
				break
			}
		}
	}

	if dataPath == "" {
		return nil, fmt.Errorf("database file not found. Tried paths: %v", possiblePaths)
	}

	// Parse the JSON
	var database AircraftDatabase
	if err := json.Unmarshal(data, &database); err != nil {
		return nil, fmt.Errorf("failed to parse JSON from %s: %v", dataPath, err)
	}

	return &database, nil
}
