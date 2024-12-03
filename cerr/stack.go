package cerr

import "runtime"

type StackPCS []uintptr

const maxStackDepth = 32

func Callers(skip ...int) StackPCS {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}
