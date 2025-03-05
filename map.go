package eutil

type mapFunc[T, V any] func(T) V

func Map[T, V any](fn mapFunc[T, V], args ...T) []V {
	res := make([]V, 0, len(args))
	for _, v := range args {
		res = append(res, fn(v))
	}
	return res
}
