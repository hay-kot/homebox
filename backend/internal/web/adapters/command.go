package adapters

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

type CommandFunc[T any] func(context.Context) (T, error)
type CommandIDFunc[T any] func(context.Context, uuid.UUID) (T, error)

// Command is an HandlerAdapter that returns a server.HandlerFunc that
// The command adapters are used to handle commands that do not accept a body
// or a query. You can think of them as a way to handle RPC style Rest Endpoints.
//
// Example:
//
//	fn := func(ctx context.Context) (interface{}, error) {
//		// do something
//		return nil, nil
//	}
//
//  r.Get("/foo", adapters.Command(fn, http.NoContent))
func Command[T any](f CommandFunc[T], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		res, err := f(r.Context())
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}

// CommandID is the same as the Command adapter but it accepts a UUID as a parameter
// in the URL. The parameter name is passed as the first argument.
//
// Example:
//
//	fn := func(ctx context.Context, id uuid.UUID) (interface{}, error) {
//		// do something
//		return nil, nil
//	}
//
//	r.Get("/foo/{id}", adapters.CommandID("id", fn, http.NoContent))
func CommandID[T any](param string, f CommandIDFunc[T], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ID, err := routeUUID(r, param)
		if err != nil {
			return err
		}

		res, err := f(r.Context(), ID)
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}
