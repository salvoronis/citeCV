package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status int, mesType string, message string) (map[string]interface{}) {
	return map[string]interface{}{"code": status, "type": mesType, "message": message}
}

func Respond(w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
