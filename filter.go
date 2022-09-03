package eutil

type filterFunc[T any] func(T) bool

func Filter[T any](fn filterFunc[T], args ...T) []T {
	res := make([]T, 0, len(args))
	for _, v := range args {
		if ok := fn(v); ok {
			res = append(res, v)
		}
	}
	return res
}
