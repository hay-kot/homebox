package adapters

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

type CommandFunc[T any] func(context.Context) (T, error)
type CommandIDFunc[T any] func(context.Context, uuid.UUID) (T, error)

func Command[T any](f CommandFunc[T], ok int) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		res, err := f(r.Context())
		if err != nil {
			return err
		}

		return server.Respond(w, ok, res)
	}
}

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
