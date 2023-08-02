package main

import (
	"fmt"
	"net/http"

	"url-shortener/internal/adapters/rest"

	"github.com/gorilla/mux"
)

func main() {
	// Create router
	router := mux.NewRouter()
	handleRequests(router)

	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Broadcast the HTTP server on port 8080 of localhost.
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func handleRequests(router *mux.Router) {
	// Endpoint to return shortened URL
	router.HandleFunc("/shorten", rest.ShortenURLHandler)
	// Endpoint to redirect shortened URL to original URL
	router.HandleFunc("/redirect", rest.RedirectURLHandler)
}
