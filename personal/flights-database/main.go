package main

import (
	"flights-database/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/aircraft", api.Handler)

	fmt.Println("🚀 Flights Database API starting on http://localhost:3000")
	fmt.Println("📖 API endpoints:")
	fmt.Println("   GET  /api/aircraft - Get all aircraft mappings")
	fmt.Println("   POST /api/aircraft - Get specific aircraft mappings")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
