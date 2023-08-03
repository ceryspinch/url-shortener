package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"url-shortener/internal/adapters/rest"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gorilla/mux"
)

// DB stores the database session imformation. Needs to be initialized once
type Database struct {
	db *sql.DB
}

func (driver *Database) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the long URL from the request
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		rest.WriteError(w, "url not provided", http.StatusBadRequest)
		return
	}

	var id int
	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", originalURL).Scan(&id)
	if err != nil {
		rest.WriteError(w, "could not create new database record", http.StatusInternalServerError)
		return
	}

	// Generate short URL
	shortURL := fmt.Sprintf(utils.ToBase62(id))

	// Print newly shortened URL
	fmt.Fprintf(w, shortURL)
}

func (driver *Database) RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the shortened URL from the request
	shortURL := r.URL.Query().Get("url")
	if shortURL == "" {
		rest.WriteError(w, "url not provided", http.StatusBadRequest)
		return
	}

	var originalURL string
	id := utils.ToBase10(shortURL)
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&originalURL)
	if err != nil {
		rest.WriteError(w, "could not retrieve original URL from database", http.StatusInternalServerError)
		return
	}

	// Redirect user to the original long URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func main() {
	// Initialise database
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}

	// Initialise database to be used with handlers
	database := &Database{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create router
	router := mux.NewRouter()
	handleRequests(router, database)

	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Broadcast the HTTP server on port 8080 of localhost.
	err = server.ListenAndServe()
	if err != nil {
		return
	}
}

func handleRequests(router *mux.Router, database *Database) {
	// Endpoint to return shortened URL
	router.HandleFunc("/shorten", database.ShortenURLHandler)
	// Endpoint to redirect shortened URL to original URL
	router.HandleFunc("/redirect", database.RedirectURLHandler)
}
