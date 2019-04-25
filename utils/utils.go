package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {"status": status, "message": message}
}

func Respond(statusCode int, w http.ResponseWriter, data map[string]interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
