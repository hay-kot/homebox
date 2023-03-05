package adapters

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

var queryDecoder = schema.NewDecoder()

func decodeQuery[T any](r *http.Request) (T, error) {
	var v T
	err := queryDecoder.Decode(&v, r.URL.Query())
	if err != nil {
		return v, err
	}

	err = validate.Check(v)
	if err != nil {
		return v, err
	}

	return v, nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T

	err := server.Decode(r, &v)
	if err != nil {
		return v, err
	}

	err = validate.Check(v)
	if err != nil {
		return v, err
	}

	return v, nil
}

func routeUUID(r *http.Request, key string) (uuid.UUID, error) {
	ID, err := uuid.Parse(chi.URLParam(r, key))
	if err != nil {
		return uuid.Nil, validate.NewRouteKeyError(key)
	}
	return ID, nil
}
