package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Respond converts a Go value to JSON and sends it to the client.
// Adapted from https://github.com/ardanlabs/service/tree/master/foundation/web
func Respond(w http.ResponseWriter, statusCode int, data interface{}) {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return
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
		panic(err)
	}
}

// ResponseError is a helper function that sends a JSON response of an error message
func RespondError(w http.ResponseWriter, statusCode int, err error) {
	eb := ErrorBuilder{}
	eb.AddError(err)
	eb.Respond(w, statusCode)
}

// RespondServerError is a wrapper around RespondError that sends a 500 internal server error. Useful for
// Sending generic errors when everything went wrong.
func RespondServerError(w http.ResponseWriter) {
	RespondError(w, http.StatusInternalServerError, errors.New("internal server error"))
}

// RespondNotFound is a helper utility for responding with a generic
// "unauthorized" error.
func RespondUnauthorized(w http.ResponseWriter) {
	RespondError(w, http.StatusUnauthorized, errors.New("unauthorized"))
}

// RespondForbidden is a helper utility for responding with a generic
// "forbidden" error.
func RespondForbidden(w http.ResponseWriter) {
	RespondError(w, http.StatusForbidden, errors.New("forbidden"))
}
