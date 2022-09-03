package eutil

type reduceFunc[T any] func(T, T) T

func Reduce[T any](fn reduceFunc[T], args ...T) (res T) {
	if len(args) == 0 {
		return
	}
	res = args[0]
	l := len(args)
	for i := 1; i < l; i++ {
		res = fn(res, args[i])
	}
	return res
}
