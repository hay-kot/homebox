package pathlib

import (
	"testing"
)

func Test_hasConflict(t *testing.T) {
	type args struct {
		path      string
		neighbors []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no conflict",
			args: args{
				path:      "foo",
				neighbors: []string{"bar", "baz"},
			},
			want: false,
		},
		{
			name: "conflict",
			args: args{
				path:      "foo",
				neighbors: []string{"bar", "foo"},
			},
			want: true,
		},
		{
			name: "conflict with different case",
			args: args{
				path:      "foo",
				neighbors: []string{"bar", "Foo"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasConflict(tt.args.path, tt.args.neighbors); got != tt.want {
				t.Errorf("hasConflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafePath(t *testing.T) {
	// override dirReader
	dirReader = func(name string) []string {
		return []string{"bar.pdf", "bar (1).pdf", "bar (2).pdf"}
	}

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no conflict",
			args: args{
				path: "/foo/foo.pdf",
			},
			want: "/foo/foo.pdf",
		},
		{
			name: "conflict",
			args: args{
				path: "/foo/bar.pdf",
			},
			want: "/foo/bar (3).pdf",
		},
		{
			name: "conflict with different case",
			args: args{
				path: "/foo/BAR.pdf",
			},
			want: "/foo/BAR (3).pdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Safe(tt.args.path); got != tt.want {
				t.Errorf("SafePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
