package adapters

import (
	"net/http"

	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
)

// Query is a server.Handler that decodes a query from the request and calls the provided function.
//
// Example:
//
//	type Query struct {
//	    Foo string `schema:"foo"`
//	}
//
//	fn := func(r *http.Request, q Query) (any, error) {
//	    // do something with q
//		return nil, nil
//	}
//
//	r.Get("/foo", adapters.Query(fn, http.StatusOK))
func Query[T any, Y any](f AdapterFunc[T, Y], ok int) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		q, err := DecodeQuery[T](r)
		if err != nil {
			return err
		}

		res, err := f(r, q)
		if err != nil {
			return err
		}

		return server.JSON(w, ok, res)
	}
}

// QueryID is a server.Handler that decodes a query and an ID from the request and calls the provided function.
//
// Example:
//
//	type Query struct {
//	    Foo string `schema:"foo"`
//	}
//
//	fn := func(r *http.Request, ID uuid.UUID, q Query) (any, error) {
//	    // do something with ID and q
//		return nil, nil
//	}
//
//	r.Get("/foo/{id}", adapters.QueryID(fn, http.StatusOK))
func QueryID[T any, Y any](param string, f IDFunc[T, Y], ok int) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ID, err := RouteUUID(r, param)
		if err != nil {
			return err
		}

		q, err := DecodeQuery[T](r)
		if err != nil {
			return err
		}

		res, err := f(r, ID, q)
		if err != nil {
			return err
		}

		return server.JSON(w, ok, res)
	}
}
