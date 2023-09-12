package lookup

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookup(t *testing.T) {
	var extractorTestTokenValue = "testTokenValue"

	var tests = []struct {
		name      string
		lookup    string
		headers   map[string]string
		query     url.Values
		cookie    map[string]string
		wantValue string
		wantErr   error
	}{
		{
			name:      "invalid lookup but use default",
			lookup:    "xx",
			headers:   map[string]string{"Authorization": "Bearer " + extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "ignore invalid lookup",
			lookup:    "header:Authorization:Bearer,xxxx",
			headers:   map[string]string{"Authorization": "Bearer " + extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "header default hit",
			lookup:    "",
			headers:   map[string]string{"Authorization": "Bearer " + extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "header hit",
			lookup:    "header:token",
			headers:   map[string]string{"token": extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "header miss",
			lookup:    "header:This-Header-Is-Not-Set",
			headers:   map[string]string{"token": extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			wantValue: "",
			wantErr:   ErrMissingValue,
		},

		{
			name:      "header filter",
			lookup:    "header:Authorization:Bearer",
			headers:   map[string]string{"Authorization": "Bearer " + extractorTestTokenValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "header filter miss",
			lookup:    "header:Authorization:Bearer",
			headers:   map[string]string{"Authorization": "Bearer   "},
			query:     nil,
			cookie:    nil,
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
		{
			name:      "argument hit",
			lookup:    "query:token",
			headers:   map[string]string{},
			query:     url.Values{"token": {extractorTestTokenValue}},
			cookie:    nil,
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "argument miss",
			lookup:    "query:token",
			headers:   map[string]string{},
			query:     nil,
			cookie:    nil,
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
		{
			name:      "cookie hit",
			lookup:    "cookie:token",
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{"token": extractorTestTokenValue},
			wantValue: extractorTestTokenValue,
			wantErr:   nil,
		},
		{
			name:      "cookie miss",
			lookup:    "cookie:token",
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{},
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
		{
			name:      "cookie miss",
			lookup:    "cookie:token",
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{"token": " "},
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
	}
	// Bearer token request
	for _, e := range tests {
		// Make request from test struct
		r := makeTestRequest("GET", "/", e.headers, e.cookie, e.query)

		// Test extractor
		value, err := NewLookup(e.lookup).ExtractValue(r)
		if value != e.wantValue {
			t.Errorf("[%v] Expected value '%v'.  Got '%v'", e.name, e.wantValue, value)
			continue
		}
		if err != e.wantErr {
			t.Errorf("[%v] Expected error '%v'.  Got '%v'", e.name, e.wantErr, err)
			continue
		}
	}
}

func TestFrom(t *testing.T) {
	t.Run("from header", func(t *testing.T) {
		r := makeTestRequest("GET", "/", map[string]string{"token": "foo"}, nil, nil)
		tk, err := FromHeader(r, "token", "")
		require.NoError(t, err)
		require.Equal(t, "foo", tk)
	})
	t.Run("from query", func(t *testing.T) {
		r := makeTestRequest("GET", "/", nil, nil, url.Values{"token": {"foo"}})
		tk, err := FromQuery(r, "token")
		require.NoError(t, err)
		require.Equal(t, "foo", tk)
	})
	t.Run("from query", func(t *testing.T) {
		r := makeTestRequest("GET", "/", nil, map[string]string{"token": "foo"}, nil)
		tk, err := FromCookie(r, "token")
		require.NoError(t, err)
		require.Equal(t, "foo", tk)
	})
}
