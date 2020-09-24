package text

import (
	"encoding/json"
	"net/http"
)

func SerializeResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	jsonResponse(w, payload, http.StatusOK)
}

/*
	Default behavior
*/
func jsonResponse(w http.ResponseWriter, payload interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
