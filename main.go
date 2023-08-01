package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var shortenedURLs = make(map[string]string)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Error struct {
	Message string `json:"message"`
}

func main() {
	// Create router
	router := mux.NewRouter()
	// Endpoint to return shortened URL
	router.HandleFunc("/shorten", ShortenURLHandler)
	// Endppoint to redirect shortened URL to original URL
	router.HandleFunc("/redirect", RedirectURLHandler)

	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Broadcast the HTTP server on port 8080 of localhost, with an handler on the path "/".
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func Base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URL to be shortened from the request
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		writeError(w, "url not provided", http.StatusBadRequest)
		return
	}

	shortenedURL := shortenURL(originalURL)

	// Print newly shortened URL
	fmt.Fprintf(w, shortenedURL)
}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the shortened URL from the request
	shortURL := r.URL.Query().Get("url")
	if shortURL == "" {
		writeError(w, "url not provided", http.StatusBadRequest)
		return
	}
	// Find original URL in the map
	originalURL := shortenedURLs[shortURL]

	// Redirect request to original/long URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func shortenURL(urlToShorten string) string {
	// get base62 encoded version of a random number
	encodedURL := Base62Encode(rand.Uint64())
	// Add this encoding to a correctly formatted URL path
	shortURL := fmt.Sprintf("http://%s.com", encodedURL)
	// Set new short URL as key in map, and original URL as value
	shortenedURLs[shortURL] = urlToShorten

	return shortURL
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
