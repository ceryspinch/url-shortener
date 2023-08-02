package rest

import "net/http"

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Get the shortened URL from the request
	shortURL := r.URL.Query().Get("url")
	if shortURL == "" {
		WriteError(w, "url not provided", http.StatusBadRequest)
		return
	}

	// Find original URL in the map if it exists
	originalURL, ok := shortenedURLs[shortURL]
	if !ok {
		WriteError(w, "invalid short url provided", http.StatusBadRequest)
		return
	}

	// Redirect request to original/long URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
