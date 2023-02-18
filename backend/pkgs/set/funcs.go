package set

// Diff returns the difference between two sets
func Diff[T key](a, b Set[T]) Set[T] {
	s := New[T]()
	for k := range a.mp {
		if !b.Contains(k) {
			s.Insert(k)
		}
	}
	return s
}

// Intersect returns the intersection between two sets
func Intersect[T key](a, b Set[T]) Set[T] {
	s := New[T]()
	for k := range a.mp {
		if b.Contains(k) {
			s.Insert(k)
		}
	}
	return s
}

// Union returns the union between two sets
func Union[T key](a, b Set[T]) Set[T] {
	s := New[T]()
	for k := range a.mp {
		s.Insert(k)
	}
	for k := range b.mp {
		s.Insert(k)
	}
	return s
}

// Xor returns the symmetric difference between two sets
func Xor[T key](a, b Set[T]) Set[T] {
	s := New[T]()
	for k := range a.mp {
		if !b.Contains(k) {
			s.Insert(k)
		}
	}
	for k := range b.mp {
		if !a.Contains(k) {
			s.Insert(k)
		}
	}
	return s
}

// Equal returns true if two sets are equal
func Equal[T key](a, b Set[T]) bool {
	if a.Len() != b.Len() {
		return false
	}
	for k := range a.mp {
		if !b.Contains(k) {
			return false
		}
	}
	return true
}

// Subset returns true if a is a subset of b
func Subset[T key](a, b Set[T]) bool {
	if a.Len() > b.Len() {
		return false
	}
	for k := range a.mp {
		if !b.Contains(k) {
			return false
		}
	}
	return true
}

// Superset returns true if a is a superset of b
func Superset[T key](a, b Set[T]) bool {
	if a.Len() < b.Len() {
		return false
	}
	for k := range b.mp {
		if !a.Contains(k) {
			return false
		}
	}
	return true
}

// Disjoint returns true if two sets are disjoint
func Disjoint[T key](a, b Set[T]) bool {
	for k := range a.mp {
		if b.Contains(k) {
			return false
		}
	}
	return true
}
