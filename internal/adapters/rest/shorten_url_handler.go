package rest

import (
	"database/sql"
	"fmt"
	"net/http"
)

type ShortenURLHandler struct {
	Database *sql.DB
}

func (handler *ShortenURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the long URL from the request
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		WriteError(w, "url not provided", http.StatusBadRequest)
		return
	}

	// Add the URL to the database and return the unique ID of that record
	var id int
	err := handler.Database.QueryRow("INSERT INTO url_shortener(url) VALUES($1) RETURNING id", originalURL).Scan(&id)
	if err != nil {
		WriteError(w, "could not create new database record", http.StatusInternalServerError)
		return
	}

	// Generate short URL
	shortURL := fmt.Sprintf("http://%s.com/", ToBase62(id))

	// Print newly shortened URL
	w.Write([]byte(shortURL))
}
