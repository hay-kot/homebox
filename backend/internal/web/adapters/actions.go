package adapters

import (
	"net/http"

	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
)

// Action is a function that adapts a function to the server.Handler interface.
// It decodes the request body into a value of type T and passes it to the function f.
// The function f is expected to return a value of type Y and an error.
//
// Example:
//
//	type Body struct {
//	    Foo string `json:"foo"`
//	}
//
//	fn := func(r *http.Request, b Body) (any, error) {
//	    // do something with b
//	    return nil, nil
//	}
//
// r.Post("/foo", adapters.Action(fn, http.StatusCreated))
func Action[T any, Y any](f AdapterFunc[T, Y], ok int) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		v, err := DecodeBody[T](r)
		if err != nil {
			return err
		}

		res, err := f(r, v)
		if err != nil {
			return err
		}

		return server.JSON(w, ok, res)
	}
}

// ActionID functions the same as Action, but it also decodes a UUID from the URL path.
//
// Example:
//
//	type Body struct {
//	    Foo string `json:"foo"`
//	}
//
//	fn := func(r *http.Request, ID uuid.UUID, b Body) (any, error) {
//	    // do something with ID and b
//	    return nil, nil
//	}
//
//	r.Post("/foo/{id}", adapters.ActionID(fn, http.StatusCreated))
func ActionID[T any, Y any](param string, f IDFunc[T, Y], ok int) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ID, err := RouteUUID(r, param)
		if err != nil {
			return err
		}

		v, err := DecodeBody[T](r)
		if err != nil {
			return err
		}

		res, err := f(r, ID, v)
		if err != nil {
			return err
		}

		return server.JSON(w, ok, res)
	}
}
