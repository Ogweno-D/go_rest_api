package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[API Request] Method: %s | URL: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func main() {
	//ServerMux for routing
	mux := http.NewServeMux()

	// Get all users
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		users := []User{
			{ID: "1", Name: "Alice"},
			{ID: "2", Name: "Bob"},
		}
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			return
		}
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Extract path parameter using Go 1.22+ built-in routing feature
		userID := r.PathValue("id")

		w.Header().Set("Content-Type", "application/json")
		user := User{ID: userID, Name: "User " + userID}
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			return
		}
	})

	// Wrap our router with the Logging Middleware
	handlerWithMiddleware := LoggingMiddleware(mux)

	port := ":8080"
	fmt.Printf("REST API Server starting on http://localhost%s...\n", port)

	// Start server
	if err := http.ListenAndServe(port, handlerWithMiddleware); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
