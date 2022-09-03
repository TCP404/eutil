package eutil

type mapFunc[T any] func(T) T

func Map[T any](fn mapFunc[T], args ...T) []T {
	res := make([]T, 0, len(args))
	for _, v := range args {
		res = append(res, fn(v))
	}
	return res
}
