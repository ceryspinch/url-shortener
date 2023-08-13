package rest

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type RedirectURLHandler struct {
	Database *sql.DB
}

func (handler *RedirectURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the short URL from the request
	shortURL := mux.Vars(r)["shortURL"]

	// Convert the short URL to the a database ID
	id := ToBase10(shortURL)

	// Get the original URL from the database
	var originalURL string
	err := handler.Database.QueryRow("SELECT url FROM urls WHERE id = $1", id).Scan(&originalURL)
	if err != nil {
		WriteError(w, "could not retrieve original URL from database", http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
