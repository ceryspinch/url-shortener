package rest

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

// WriteError is a simple utility function for error responses, used to keep the handler code
// cleaner and avoid duplication.
func WriteError(w http.ResponseWriter, msg string, status int) {
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
