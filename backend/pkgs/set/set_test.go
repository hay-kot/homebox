package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		args args
		want Set[string]
	}{
		{
			name: "new",
			args: args{
				v: []string{"a", "b", "c"},
			},
			want: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
		},
		{
			name: "new empty",
			args: args{
				v: []string{},
			},
			want: Set[string]{
				mp: map[string]struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Insert(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		s    Set[string]
		args args
		want Set[string]
	}{
		{
			name: "insert",
			s: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
			args: args{
				v: []string{"d", "e", "f"},
			},
			want: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
					"d": {},
					"e": {},
					"f": {},
				},
			},
		},
		{
			name: "insert empty",
			s: Set[string]{
				mp: map[string]struct{}{},
			},
			args: args{
				v: []string{"a", "b", "c"},
			},
			want: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Insert(tt.args.v...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Set.Insert() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Delete(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		s    Set[string]
		args args
		want Set[string]
	}{
		{
			name: "insert",
			s: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
					"d": {},
					"e": {},
					"f": {},
				},
			},
			args: args{
				v: []string{"d", "e", "f"},
			},
			want: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
		},
		{
			name: "delete empty",
			s: Set[string]{
				mp: map[string]struct{}{},
			},
			args: args{
				v: []string{},
			},
			want: Set[string]{
				mp: map[string]struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.v...)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Set.Delete() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_ContainsAll(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		s    Set[string]
		args args
		want bool
	}{
		{
			name: "contains",
			s: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
			args: args{
				v: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "contains empty",
			s: Set[string]{
				mp: map[string]struct{}{},
			},
			args: args{
				v: []string{},
			},
			want: true,
		},
		{
			name: "not contains",
			s: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
			args: args{
				v: []string{"d", "e", "f"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ContainsAll(tt.args.v...); got != tt.want {
				t.Errorf("Set.ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Slice(t *testing.T) {
	tests := []struct {
		name string
		s    Set[string]
		want []string
	}{
		{
			name: "slice",
			s: Set[string]{
				mp: map[string]struct{}{
					"a": {},
					"b": {},
					"c": {},
				},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "slice empty",
			s: Set[string]{
				mp: map[string]struct{}{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Slice()

			sort.Strings(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
