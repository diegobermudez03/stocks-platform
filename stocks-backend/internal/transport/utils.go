package transport

import (
	"encoding/json"
	"net/http"
)


func WriteJSON(w http.ResponseWriter, status int, body any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(body)
}

func WriteError(w http.ResponseWriter, status int, err error) error{
	return WriteJSON(
		w,
		status, 
		map[string]any{
			"error" : err.Error(),
		},
	)
}