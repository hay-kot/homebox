package server

import (
	"net/http"
)

// ErrorBuilder is a helper type to build a response that contains an array of errors.
// Typical use cases are for returning an array of validation errors back to the user.
//
// Example:
//
//
// {
//  "errors": [
//    "invalid id",
//    "invalid name",
//    "invalid description"
//  ],
//  "message": "Unprocessable Entity",
//  "status": 422
// }
//
type ErrorBuilder struct {
	errs []string
}

// HasErrors returns true if the ErrorBuilder has any errors.
func (eb *ErrorBuilder) HasErrors() bool {
	if (eb.errs == nil) || (len(eb.errs) == 0) {
		return false
	}
	return true
}

// AddError adds an error to the ErrorBuilder if an error is not nil. If the
// Error is nil, then nothing is added.
func (eb *ErrorBuilder) AddError(err error) {
	if err != nil {
		if eb.errs == nil {
			eb.errs = make([]string, 0)
		}

		eb.errs = append(eb.errs, err.Error())
	}
}

// Respond sends a JSON response with the ErrorBuilder's errors. If there are no errors, then
// the errors field will be an empty array.
func (eb *ErrorBuilder) Respond(w http.ResponseWriter, statusCode int) {
	Respond(w, statusCode, Wrap(nil).AddError(http.StatusText(statusCode), eb.errs))
}
