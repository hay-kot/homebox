// Package adapters provides functions to adapt functions to the server.Handler interface.
package adapters

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

type AdapterFunc[T any, Y any] func(context.Context, T) (Y, error)
type IDFunc[T any, Y any] func(context.Context, uuid.UUID, T) (Y, error)

// Query is a server.Handler that decodes a query from the request and calls the provided function.
func Query[T any, Y any](f AdapterFunc[T, Y], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		q, err := decodeQuery[T](r)
		if err != nil {
			return err
		}

		res, err := f(r.Context(), q)
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}

// QueryID is a server.Handler that decodes a query and an ID from the request and calls the provided function.
func QueryID[T any, Y any](param string, f IDFunc[T, Y], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ID, err := routeUUID(r, param)
		if err != nil {
			return err
		}

		q, err := decodeQuery[T](r)
		if err != nil {
			return err
		}

		res, err := f(r.Context(), ID, q)
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}

// Action is a function that adapts a function to the server.Handler interface.
// It decodes the request body into a value of type T and passes it to the function f.
// The function f is expected to return a value of type Y and an error.
//
// Note: Action differs from Query in that it decodes the request body.
func Action[T any, Y any](f AdapterFunc[T, Y], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		v, err := decode[T](r)
		if err != nil {
			return err
		}

		res, err := f(r.Context(), v)
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}

// ActionID functions the same as Action, but it also decodes a UUID from the URL path.
func ActionID[T any, Y any](param string, f IDFunc[T, Y], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ID, err := routeUUID(r, param)
		if err != nil {
			return err
		}

		v, err := decode[T](r)
		if err != nil {
			return err
		}

		res, err := f(r.Context(), ID, v)
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}