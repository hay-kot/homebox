package lookup

import (
	"errors"
	"net/http"
	"strings"
)

// ErrMissingValue can be thrown by follow
// if value with a HTTP header, the value header needs to be set
// if value with URL Query, the query value variable is empty
// if value with a cookie, the value cookie is empty
var ErrMissingValue = errors.New("no value present in request")

// Lookup is a tool that looks up the value from http request, such as token
type Lookup struct {
	extractors MultiExtractor
}

// NewLookup new a lookup.
// lookup is a string in the form of "<source>:<name>[:<prefix>]" that is used
// to extract value from the request.
// use like "header:<name>[:<prefix>],query:<name>,cookie:<name>,param:<name>"
// Optional, Default value "header:Authorization:Bearer" for json web token.
// Possible values:
// - "header:<name>:<prefix>", <prefix> is a special string in the header, Possible value is "Bearer"
// - "query:<name>"
// - "cookie:<name>"
func NewLookup(lookup string) *Lookup {
	if lookup == "" {
		lookup = "header:Authorization:Bearer"
	}
	methods := strings.Split(lookup, ",")
	lookups := make(MultiExtractor, 0, len(methods))
	for _, method := range methods {
		parts := strings.Split(strings.TrimSpace(method), ":")
		if !(len(parts) == 2 || len(parts) == 3) {
			continue
		}
		switch parts[0] {
		case "header":
			prefix := ""
			if len(parts) == 3 {
				prefix = strings.TrimSpace(parts[2])
			}
			lookups = append(lookups, HeaderExtractor{strings.TrimSpace(parts[1]), prefix})
		case "query":
			lookups = append(lookups, ArgumentExtractor(parts[1]))
		case "cookie":
			lookups = append(lookups, CookieExtractor(parts[1]))
		}
	}
	if len(lookups) == 0 {
		lookups = append(lookups, HeaderExtractor{"Authorization", "Bearer"})
	}
	return &Lookup{lookups}
}

// ExtractValue extract value from http request.
func (sf *Lookup) ExtractValue(r *http.Request) (string, error) {
	value, err := sf.extractors.ExtractValue(r)
	if err != nil || value == "" {
		return "", ErrMissingValue
	}
	return value, nil
}

// FromHeader get value from header
// key is a header key, like "Authorization"
// prefix is a string in the header, like "Bearer", if it is empty, only will return value.
func FromHeader(r *http.Request, key, prefix string) (string, error) {
	return HeaderExtractor{key, prefix}.ExtractValue(r)
}

// FromQuery get value from query
// key is a query key
func FromQuery(r *http.Request, key string) (string, error) {
	return ArgumentExtractor(key).ExtractValue(r)
}

// FromCookie get value from Cookie
// key is a cookie key
func FromCookie(r *http.Request, key string) (string, error) {
	return CookieExtractor(key).ExtractValue(r)
}
