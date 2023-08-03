package rest

// import (
// 	"fmt"
// 	"net/http"
// 	"url-shortener/utils"
// )

// var (
// 	shortenedURLs = make(map[string]string)
// )

// func (driver *main.DBClient) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
// 	// Get the long URL from the request
// 	originalURL := r.URL.Query().Get("url")
// 	if originalURL == "" {
// 		WriteError(w, "url not provided", http.StatusBadRequest)
// 		return
// 	}

// 	var id int
// 	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", originalURL).Scan(&id)
// 	fmt.Println(id)

// 	if err != nil {
// 		fmt.Print("oh no")
// 	}

// 	// Generate short URL
// 	shortURL := fmt.Sprintf(utils.ToBase62(id))

// 	// Print newly shortened URL
// 	fmt.Fprintf(w, shortURL)
// }
