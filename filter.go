package eutil

func Filter[T any](fn func(T) (T, bool), args ...T) []T {
	res := make([]T, 0, len(args))
	for _, v := range args {
		if r, ok := fn(v); ok {
			res = append(res, r)
		}
	}
	return res
}
