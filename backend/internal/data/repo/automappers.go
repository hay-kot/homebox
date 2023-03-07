package repo

type MapFunc[T any, U any] func(T) U

func (a MapFunc[T, U]) Map(v T) U {
	return a(v)
}

func (a MapFunc[T, U]) MapEach(v []T) []U {
	result := make([]U, len(v))
	for i, item := range v {
		result[i] = a(item)
	}
	return result
}

func (a MapFunc[T, U]) MapErr(v T, err error) (U, error) {
	if err != nil {
		var zero U
		return zero, err
	}

	return a(v), nil
}

func (a MapFunc[T, U]) MapEachErr(v []T, err error) ([]U, error) {
	if err != nil {
		return nil, err
	}

	return a.MapEach(v), nil
}
