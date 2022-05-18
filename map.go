package eutil

func Map[T any](fn func(T) T, args ...T) []T {
	res := make([]T, 0, len(args))
	for _, v := range args {
		res = append(res, fn(v))
	}
	return res
}