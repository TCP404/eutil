package etl

import "errors"

// BatchFn 批处理并统计，相当于 for else 语法
func BatchFn[T any](iter <-chan T, size int, fn func([]T) (int, error)) (int, error) {
	var (
		batch = make([]T, 0, size)
		total int
	)
	var errorArr []error
	for v := range iter {
		batch = append(batch, v)
		if len(batch) < size {
			continue
		}
		cnt, err := fn(batch)
		if err != nil {
			errorArr = append(errorArr, err)
		}
		total += cnt
		batch = batch[:0]
	}
	// handle rest
	cnt, err := fn(batch)
	if err != nil {
		errorArr = append(errorArr, err)
	}
	total += cnt
	return total, errors.Join(errorArr...)
}

// BatchUniqueFn 批处理并统计，相当于 for else 语法，会进行去重
func BatchUniqueFn[T any, U comparable](iter <-chan T, size int, uniqueKey func(T) U, fn func([]T) int) int {
	var (
		batch  = make([]T, 0, size)
		unique = make(map[U]struct{})
		total  int
	)
	for v := range iter {
		_, exist := unique[uniqueKey(v)]
		if exist {
			continue
		}
		unique[uniqueKey(v)] = struct{}{}

		batch = append(batch, v)
		if len(batch) < size {
			continue
		}
		cnt := fn(batch)
		total += cnt
		batch = batch[:0]
	}
	// handle rest
	cnt := fn(batch)
	total += cnt
	return total
}
