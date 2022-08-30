package server

import (
	"encoding/json"
	"net/http"
)

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value.
func Decode(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}
	return nil
}

// GetId is a shotcut to get the id from the request URL or return a default value
func GetParam(r *http.Request, key, d string) string {
	val := r.URL.Query().Get(key)

	if val == "" {
		return d
	}

	return val
}

// GetSkip is a shotcut to get the skip from the request URL parameters
func GetSkip(r *http.Request, d string) string {
	return GetParam(r, "skip", d)
}

// GetSkip is a shotcut to get the skip from the request URL parameters
func GetId(r *http.Request, d string) string {
	return GetParam(r, "id", d)
}

// GetLimit is a shotcut to get the limit from the request URL parameters
func GetLimit(r *http.Request, d string) string {
	return GetParam(r, "limit", d)
}

// GetQuery is a shotcut to get the sort from the request URL parameters
func GetQuery(r *http.Request, d string) string {
	return GetParam(r, "query", d)
}
