package main

import (
	"fmt"
	"net/http"
	"url-shortener/internal/adapters/rest"
	"url-shortener/models"

	"github.com/gorilla/mux"
)

func main() {
	// Initialise database
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create router
	router := mux.NewRouter()

	// Create handlers
	shortenURLHandler := rest.ShortenURLHandler{Database: db}
	redirectURLHandler := rest.RedirectURLHandler{Database: db}

	// Endpoint to return shortened URL
	router.HandleFunc("/shorten", shortenURLHandler.ServeHTTP)

	// Endpoint to redirect shortened URL to original URL
	router.HandleFunc("/redirect", redirectURLHandler.ServeHTTP)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}

	// Broadcast the HTTP server on port 8080 of localhost.
	err = server.ListenAndServe()
	if err != nil {
		return
	}
}
