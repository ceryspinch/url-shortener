package rest

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"url-shortener/utils"
)

var (
	shortenedURLs = make(map[string]string)
	jsonDBFile    = "../build/database/in_mem_database.json"
)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the long URL from the request
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		WriteError(w, "url not provided", http.StatusBadRequest)
		return
	}

	// Generate short URL
	shortURL := fmt.Sprintf("http://%s.com/", utils.Base62Encode(rand.Uint64()))

	// Store short URL in map with original URL as value
	shortenedURLs[shortURL] = originalURL

	// Update database
	updatedDatabase, err := json.MarshalIndent(shortenedURLs, "", "    ")
	if err != nil {
		WriteError(w, "error marshalling data into db", http.StatusBadRequest)
		return
	}
	os.WriteFile(jsonDBFile, updatedDatabase, 0644)

	// Print newly shortened URL
	fmt.Fprintf(w, shortURL)
}
