package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"url-shortener/utils"

	"github.com/gorilla/mux"
)

var (
	shortenedURLs = make(map[string]string)
	jsonDBFile    = "database/in_mem_database.json"
)

type Error struct {
	Message string `json:"message"`
}

func main() {
	// Create router
	router := mux.NewRouter()
	// Endpoint to return shortened URL
	router.HandleFunc("/shorten", ShortenURLHandler)
	// Endpoint to redirect shortened URL to original URL
	router.HandleFunc("/", RedirectURLHandler)

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

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the long URL from the request
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		writeError(w, "url not provided", http.StatusBadRequest)
		return
	}

	// Generate short URL
	shortURL := fmt.Sprintf("http://%s.com/", utils.Base62Encode(rand.Uint64()))

	// Store short URL in map with original URL as value
	shortenedURLs[shortURL] = originalURL

	// Update database
	updatedDatabase, err := json.MarshalIndent(shortenedURLs, "", "    ")
	if err != nil {
		writeError(w, "error marshalling data into db", http.StatusBadRequest)
		return
	}
	os.WriteFile(jsonDBFile, updatedDatabase, 0644)

	// Print newly shortened URL
	fmt.Fprintf(w, shortURL)
}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the shortened URL from the request
	shortURL := r.URL.Query().Get("url")
	if shortURL == "" {
		writeError(w, "url not provided", http.StatusBadRequest)
		return
	}

	// Find original URL in the map if it exists
	originalURL, ok := shortenedURLs[shortURL]
	if !ok {
		writeError(w, "invalid short url provided", http.StatusBadRequest)
		return
	}

	// Redirect request to original/long URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

// writeError is a simple utility function for error responses, used to keep the handler code
// cleaner and avoid duplication.
func writeError(w http.ResponseWriter, msg string, status int) {
	restErr := Error{
		Message: msg,
	}

	errResp, err := json.Marshal(restErr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(errResp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
