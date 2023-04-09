package adapters

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
)

type CommandFunc[T any] func(*http.Request) (T, error)
type CommandIDFunc[T any] func(*http.Request, uuid.UUID) (T, error)

// Command is an HandlerAdapter that returns a errchain.HandlerFunc that
// The command adapters are used to handle commands that do not accept a body
// or a query. You can think of them as a way to handle RPC style Rest Endpoints.
//
// Example:
//
//		fn := func(r *http.Request) (interface{}, error) {
//			// do something
//			return nil, nil
//		}
//
//	 r.Get("/foo", adapters.Command(fn, http.NoContent))
func Command[T any](f CommandFunc[T], ok int) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		res, err := f(r)
		if err != nil {
			return err
		}

		return server.JSON(w, ok, res)
	}
}

// CommandID is the same as the Command adapter but it accepts a UUID as a parameter
// in the URL. The parameter name is passed as the first argument.
//
// Example:
//
//	fn := func(r *http.Request, id uuid.UUID) (interface{}, error) {
//		// do something
//		return nil, nil
//	}
//
//	r.Get("/foo/{id}", adapters.CommandID("id", fn, http.NoContent))
func CommandID[T any](param string, f CommandIDFunc[T], ok int) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ID, err := RouteUUID(r, param)
		if err != nil {
			return err
		}

		res, err := f(r, ID)
		if err != nil {
			return err
		}

		return server.JSON(w, ok, res)
	}
}
