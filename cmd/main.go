package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/adapters/rest"
	"url-shortener/models"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	shortenURLHandler := rest.ShortenURLHandler{Database: db}
	redirectURLHandler := rest.RedirectURLHandler{Database: db}

	// Endpoint to return shortened URL
	router.HandleFunc("/shorten", shortenURLHandler.ServeHTTP)

	// Endpoint to redirect shortened URL to original URL
	router.HandleFunc("/{shortURL}", redirectURLHandler.ServeHTTP)

	http.Handle("/", router)

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
