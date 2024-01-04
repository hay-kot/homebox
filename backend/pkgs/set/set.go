// Package set provides a simple set implementation.
package set

type key interface {
	comparable
}

type Set[T key] struct {
	mp map[T]struct{}
}

func Make[T key](size int) Set[T] {
	return Set[T]{
		mp: make(map[T]struct{}, size),
	}
}

func New[T key](v ...T) Set[T] {
	mp := make(map[T]struct{}, len(v))

	s := Set[T]{mp}

	s.Insert(v...)
	return s
}

func (s Set[T]) Insert(v ...T) {
	for _, e := range v {
		s.mp[e] = struct{}{}
	}
}

func (s Set[T]) Remove(v ...T) {
	for _, e := range v {
		delete(s.mp, e)
	}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s.mp[v]
	return ok
}

func (s Set[T]) ContainsAll(v ...T) bool {
	for _, e := range v {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s.mp))
	for k := range s.mp {
		slice = append(slice, k)
	}
	return slice
}

func (s Set[T]) Len() int {
	return len(s.mp)
}
