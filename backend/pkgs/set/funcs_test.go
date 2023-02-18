package set

import (
	"reflect"
	"testing"
)

type args struct {
	a Set[string]
	b Set[string]
}

var (
	argsBasic = args{
		a: New("a", "b", "c"),
		b: New("b", "c", "d"),
	}

	argsNoOverlap = args{
		a: New("a", "b", "c"),
		b: New("d", "e", "f"),
	}

	argsIdentical = args{
		a: New("a", "b", "c"),
		b: New("a", "b", "c"),
	}
)

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		args args
		want Set[string]
	}{
		{
			name: "diff basic",
			args: argsBasic,
			want: New("a"),
		},
		{
			name: "diff empty",
			args: argsIdentical,
			want: New[string](),
		},
		{
			name: "diff no overlap",
			args: argsNoOverlap,
			want: New("a", "b", "c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name string
		args args
		want Set[string]
	}{
		{
			name: "intersect basic",
			args: argsBasic,
			want: New("b", "c"),
		},
		{
			name: "identical sets",
			args: argsIdentical,
			want: New("a", "b", "c"),
		},
		{
			name: "no overlap",
			args: argsNoOverlap,
			want: New[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		args args
		want Set[string]
	}{
		{
			name: "intersect basic",
			args: argsBasic,
			want: New("a", "b", "c", "d"),
		},
		{
			name: "identical sets",
			args: argsIdentical,
			want: New("a", "b", "c"),
		},
		{
			name: "no overlap",
			args: argsNoOverlap,
			want: New("a", "b", "c", "d", "e", "f"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		name string
		args args
		want Set[string]
	}{
		{
			name: "xor basic",
			args: argsBasic,
			want: New("a", "d"),
		},
		{
			name: "identical sets",
			args: argsIdentical,
			want: New[string](),
		},
		{
			name: "no overlap",
			args: argsNoOverlap,
			want: New("a", "b", "c", "d", "e", "f"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xor(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal basic",
			args: argsBasic,
			want: false,
		},
		{
			name: "identical sets",
			args: argsIdentical,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	type args struct {
		a Set[string]
		b Set[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "subset basic",
			args: args{
				a: New("a", "b"),
				b: New("a", "b", "c"),
			},
			want: true,
		},
		{
			name: "subset basic false",
			args: args{
				a: New("a", "b", "d"),
				b: New("a", "b", "c"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subset(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuperset(t *testing.T) {
	type args struct {
		a Set[string]
		b Set[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "superset basic",
			args: args{
				a: New("a", "b", "c"),
				b: New("a", "b"),
			},
			want: true,
		},
		{
			name: "superset basic false",
			args: args{
				a: New("a", "b", "c"),
				b: New("a", "b", "d"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Superset(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisjoint(t *testing.T) {
	type args struct {
		a Set[string]
		b Set[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "disjoint basic",
			args: args{
				a: New("a", "b"),
				b: New("c", "d"),
			},
			want: true,
		},
		{
			name: "disjoint basic false",
			args: args{
				a: New("a", "b"),
				b: New("b", "c"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Disjoint(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Disjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
