package responseHandler

import (
	"encoding/json"
	"net/http"
)

type JSONErr struct {
	Code    int    `json:"code"`
	Message string `json:"text"`
}

func WithError(w http.ResponseWriter, error JSONErr) {
	WithJSON(w, error.Code, map[string]string{"error": error.Message})
}

func WithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
