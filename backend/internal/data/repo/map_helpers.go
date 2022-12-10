package repo

// mapTErrFunc is a factory function that returns a mapper function that
// wraps the given mapper function but first will check for an error and
// return the error if present.
//
// Helpful for wrapping database calls that return both a value and an error
func mapTErrFunc[T any, Y any](fn func(T) Y) func(T, error) (Y, error) {
	return func(t T, err error) (Y, error) {
		if err != nil {
			var zero Y
			return zero, err
		}

		return fn(t), nil
	}
}

func mapTEachFunc[T any, Y any](fn func(T) Y) func([]T) []Y {
	return func(items []T) []Y {
		result := make([]Y, len(items))
		for i, item := range items {
			result[i] = fn(item)
		}

		return result
	}
}

func mapTEachErrFunc[T any, Y any](fn func(T) Y) func([]T, error) ([]Y, error) {
	return func(items []T, err error) ([]Y, error) {
		if err != nil {
			return nil, err
		}

		result := make([]Y, len(items))
		for i, item := range items {
			result[i] = fn(item)
		}

		return result, nil
	}
}

func mapEach[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}
