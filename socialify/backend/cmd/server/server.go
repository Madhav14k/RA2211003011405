package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Call the standalone API server directly
	http.Handle("/api/auth/register", http.NotFoundHandler())
	http.Handle("/api/auth/token", http.NotFoundHandler())
	http.Handle("/api/users/top", http.NotFoundHandler())
	http.Handle("/api/posts/latest", http.NotFoundHandler())
	http.Handle("/api/posts/popular", http.NotFoundHandler())

	httpServer := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("Server is running on port 8081...")
	log.Fatal(httpServer.ListenAndServe())
}
