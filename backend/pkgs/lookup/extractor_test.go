package lookup

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestExtractor(t *testing.T) {
	var extractorTestValue = "testTokenValue"

	var tests = []struct {
		name      string
		extractor Extractor
		headers   map[string]string
		query     url.Values
		cookie    map[string]string
		wantValue string
		wantErr   error
	}{
		{
			name:      "header hit",
			extractor: HeaderExtractor{"token", ""},
			headers:   map[string]string{"token": extractorTestValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestValue,
			wantErr:   nil,
		},
		{
			name:      "header miss",
			extractor: HeaderExtractor{"This-Header-Is-Not-Set", ""},
			headers:   map[string]string{"token": extractorTestValue},
			query:     nil,
			cookie:    nil,
			wantValue: "",
			wantErr:   ErrMissingValue,
		},

		{
			name:      "header filter",
			extractor: HeaderExtractor{"Authorization", "Bearer"},
			headers:   map[string]string{"Authorization": "Bearer " + extractorTestValue},
			query:     nil,
			cookie:    nil,
			wantValue: extractorTestValue,
			wantErr:   nil,
		},
		{
			name:      "header filter miss",
			extractor: HeaderExtractor{"Authorization", "Bearer"},
			headers:   map[string]string{"Authorization": "Bearer   "},
			query:     nil,
			cookie:    nil,
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
		{
			name:      "argument hit",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     url.Values{"token": {extractorTestValue}},
			cookie:    nil,
			wantValue: extractorTestValue,
			wantErr:   nil,
		},
		{
			name:      "argument miss",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    nil,
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
		{
			name:      "cookie hit",
			extractor: CookieExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{"token": extractorTestValue},
			wantValue: extractorTestValue,
			wantErr:   nil,
		},
		{
			name:      "cookie miss",
			extractor: ArgumentExtractor("token"),
			headers:   map[string]string{},
			query:     nil,
			cookie:    map[string]string{},
			wantValue: "",
			wantErr:   ErrMissingValue,
		},
		{
			name:      "cookie miss",
			extractor: ArgumentExtractor("token"),
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
		value, err := e.extractor.ExtractValue(r)
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

func makeTestRequest(method, path string, headers, cookie map[string]string, urlArgs url.Values) *http.Request {
	r, _ := http.NewRequest(method, fmt.Sprintf("%v?%v", path, urlArgs.Encode()), nil)
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	for k, v := range cookie {
		r.AddCookie(&http.Cookie{Name: k, Value: v})
	}
	return r
}
