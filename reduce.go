package eutil

func Reduce[T any](fn func(T, T) T, args ...T) T {
	tmp := args[0]
	l := len(args)
	for i := 1; i < l; i++ {
		tmp = fn(tmp, args[i])
	}
	return tmp
}
