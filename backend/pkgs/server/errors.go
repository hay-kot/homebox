package server

import "errors"

type shutdownError struct {
	message string
}

func (e *shutdownError) Error() string {
	return e.message
}

// ShutdownError returns an error that indicates that the server has lost
// integrity and should be shut down.
func ShutdownError(message string) error {
	return &shutdownError{message}
}

// IsShutdownError returns true if the error is a shutdown error.
func IsShutdownError(err error) bool {
	var e *shutdownError
	return errors.As(err, &e)
}
