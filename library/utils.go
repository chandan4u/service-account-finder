package library

import (
	"encoding/json"
	"net/http"
)

// Message response message handler with code and message
func Message(errorCode string, message string) map[string]interface{} {
	return map[string]interface{}{"STATUS_CODE": errorCode, "MESSAGE": message}
}

// Respond response handler for api's
func Respond(w http.ResponseWriter, data map[string]interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// NotFoundHandler  not found middleware message for invalid api request
var NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
	Respond(w, Message("RESOURCE_NOT_FOUND", "This resources was not found on our server"), 404)
}
