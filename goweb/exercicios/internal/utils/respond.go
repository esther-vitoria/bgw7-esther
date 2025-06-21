package utils

import (
	"encoding/json"
	"net/http"
)

// Func para responder os dados
func Respond(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// Struct que responde
type RespondBodyProduct struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   bool
}
