package adapters

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/safeserve/server"
)

var queryDecoder = schema.NewDecoder()

func DecodeQuery[T any](r *http.Request) (T, error) {
	var v T
	err := queryDecoder.Decode(&v, r.URL.Query())
	if err != nil {
		return v, errors.Wrap(err, "decoding error")
	}

	err = validate.Check(v)
	if err != nil {
		return v, errors.Wrap(err, "validation error")
	}

	return v, nil
}

func DecodeBody[T any](r *http.Request) (T, error) {
	var v T

	err := server.Decode(r, &v)
	if err != nil {
		return v, errors.Wrap(err, "body decoding error")
	}

	err = validate.Check(v)
	if err != nil {
		return v, errors.Wrap(err, "validation error")
	}

	return v, nil
}

func RouteUUID(r *http.Request, key string) (uuid.UUID, error) {
	ID, err := uuid.Parse(chi.URLParam(r, key))
	if err != nil {
		return uuid.Nil, validate.NewRouteKeyError(key)
	}
	return ID, nil
}
