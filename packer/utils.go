package packer

func Filter[T any](fn func(T) bool, args ...T) []T {
	res := make([]T, 0, len(args))
	for _, v := range args {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
