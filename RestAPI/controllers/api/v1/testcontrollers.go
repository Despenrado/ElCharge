package v1

import (
	"net/http"
)

// TestAPIV1 ...
func TestAPIV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, http.StatusOK, "TEST")
	}
}
