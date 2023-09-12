package lookup

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// Extractor is an interface for extracting a value from an HTTP request.
// The ExtractValue method should return a value string or an error.
// If no value is present, you must return ErrMissingValue.
type Extractor interface {
	ExtractValue(*http.Request) (string, error)
}

// MultiExtractor tries Extractors in order until one returns a value string or an error occurs
type MultiExtractor []Extractor

func (e MultiExtractor) ExtractValue(req *http.Request) (string, error) {
	// loop over header names and return the first one that contains data
	for _, extractor := range e {
		if val, err := extractor.ExtractValue(req); val != "" {
			return val, nil
		} else if !errors.Is(err, ErrMissingValue) {
			return "", err
		}
	}
	return "", ErrMissingValue
}

// HeaderExtractor is an extractor for finding a value in a header.
// Looks at each specified header in order until there's a match
type HeaderExtractor struct {
	// The key of the header
	// Required
	Key string
	// Strips 'Bearer ' prefix from bearer value string.
	// Possible value is "Bearer"
	// Optional
	Prefix string
}

func (e HeaderExtractor) ExtractValue(r *http.Request) (string, error) {
	// loop over header names and return the first one that contains data
	return stripHeadValuePrefixFromValueString(e.Prefix)(r.Header.Get(e.Key))
}

// ArgumentExtractor extracts a value from request arguments.  This includes a POSTed form or
// GET URL arguments.
// This extractor calls `ParseMultipartForm` on the request
type ArgumentExtractor string

func (e ArgumentExtractor) ExtractValue(r *http.Request) (string, error) {
	// Make sure form is parsed
	_ = r.ParseMultipartForm(10e6)

	tk := strings.TrimSpace(r.Form.Get(string(e)))
	if tk != "" {
		return tk, nil
	}
	return "", ErrMissingValue
}

// CookieExtractor extracts a value from cookie.
type CookieExtractor string

func (e CookieExtractor) ExtractValue(r *http.Request) (string, error) {
	cookie, err := r.Cookie(string(e))
	if err != nil {
		return "", ErrMissingValue
	}
	val, _ := url.QueryUnescape(cookie.Value)
	if val = strings.TrimSpace(val); val != "" {
		return val, nil
	}
	return "", ErrMissingValue
}

// Strips like 'Bearer ' prefix from value string with header name
func stripHeadValuePrefixFromValueString(prefix string) func(string) (string, error) {
	return func(tok string) (string, error) {
		tok = strings.TrimSpace(tok)
		if tok == "" {
			return "", ErrMissingValue
		}
		l := len(prefix)
		if l == 0 {
			return tok, nil
		}
		// Should be a bearer value
		if len(tok) > l && strings.EqualFold(tok[:l], prefix) {
			if tok = strings.TrimSpace(tok[l+1:]); tok != "" {
				return tok, nil
			}
		}
		return "", ErrMissingValue
	}
}
