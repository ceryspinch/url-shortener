package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var shortenedURLs = make(map[string]string)

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
	shortURL := r.URL.Query().Get("shorturl")
	if shortURL == "" {
		writeError(w, "shortened url not provided", http.StatusBadRequest)
		return
	}
	// Find original URL in the map
	originalURL := shortenedURLs[shortURL]
	fmt.Print(originalURL)

	// Redirect request to original/long URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func shortenURL(urlToShorten string) string {
	// TODO: Find way to shorten URL instead of this
	shortURL := uuid.New().String()
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
