package utils

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

// Respond write respond to network channel
func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	Respond(w, r, code, map[string]string{"error": err.Error()})
}
