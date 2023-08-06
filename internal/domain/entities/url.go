package entities

// URL represents a URL record in the url_shortener database
type URL struct {
	ID  int    `json:"ID"`
	URL string `json:"URL"`
}
