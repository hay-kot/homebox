package server

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Respond converts a Go value to JSON and sends it to the client.
// Adapted from https://github.com/ardanlabs/service/tree/master/foundation/web
func Respond(w http.ResponseWriter, statusCode int, data interface{}) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", ContentJSON)

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
