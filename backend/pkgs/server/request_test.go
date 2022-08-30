package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func TestDecode(t *testing.T) {
	type args struct {
		r   *http.Request
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check_error",
			args: args{
				r: &http.Request{
					Body: http.NoBody,
				},
				val: make(map[string]interface{}),
			},
			wantErr: true,
		},
		{
			name: "check_success",
			args: args{
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"test","data":"test"}`)),
				val: TestStruct{
					Name: "test",
					Data: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decode(tt.args.r, &tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetParam(t *testing.T) {
	type args struct {
		r   *http.Request
		key string
		d   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check_default",
			args: args{
				r:   httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"test","data":"test"}`)),
				key: "id",
				d:   "default",
			},
			want: "default",
		},
		{
			name: "check_id",
			args: args{
				r:   httptest.NewRequest("POST", "/item?id=123", strings.NewReader(`{"name":"test","data":"test"}`)),
				key: "id",
				d:   "",
			},
			want: "123",
		},
		{
			name: "check_query",
			args: args{
				r:   httptest.NewRequest("POST", "/item?query=hello-world", strings.NewReader(`{"name":"test","data":"test"}`)),
				key: "query",
				d:   "",
			},
			want: "hello-world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetParam(tt.args.r, tt.args.key, tt.args.d); got != tt.want {
				t.Errorf("GetParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSkip(t *testing.T) {
	type args struct {
		r *http.Request
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check_default",
			args: args{
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"test","data":"test"}`)),
				d: "0",
			},
			want: "0",
		},
		{
			name: "check_skip",
			args: args{
				r: httptest.NewRequest("POST", "/item?skip=107", strings.NewReader(`{"name":"test","data":"test"}`)),
				d: "0",
			},
			want: "107",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSkip(tt.args.r, tt.args.d); got != tt.want {
				t.Errorf("GetSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLimit(t *testing.T) {
	type args struct {
		r *http.Request
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check_default",
			args: args{
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"test","data":"test"}`)),
				d: "0",
			},
			want: "0",
		},
		{
			name: "check_limit",
			args: args{
				r: httptest.NewRequest("POST", "/item?limit=107", strings.NewReader(`{"name":"test","data":"test"}`)),
				d: "0",
			},
			want: "107",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLimit(tt.args.r, tt.args.d); got != tt.want {
				t.Errorf("GetLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetQuery(t *testing.T) {
	type args struct {
		r *http.Request
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check_default",
			args: args{
				r: httptest.NewRequest("POST", "/", strings.NewReader(`{"name":"test","data":"test"}`)),
				d: "0",
			},
			want: "0",
		},
		{
			name: "check_query",
			args: args{
				r: httptest.NewRequest("POST", "/item?query=hello-query", strings.NewReader(`{"name":"test","data":"test"}`)),
				d: "0",
			},
			want: "hello-query",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuery(tt.args.r, tt.args.d); got != tt.want {
				t.Errorf("GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
