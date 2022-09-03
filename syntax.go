package eutil

func If[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
}

// Or iterates over the given functions and compares the return value to a zero
// value, returning it as soon as it gets a nonzero value. It is a parody of the
// OR operator in Python.
//
// You can get the basic usage on example_test file.
// In additional, there is another solution:
//
//	type oer struct {
//		or   *oer
//		val string
//	}
//
//	func cmdFlag() *oer {
//		var configPath string
//		flag.StringVar(&configPath, "c", "", "")
//		flag.Parse()
//		if configPath != "" {
//			// yes
//		}
//		return &oer{val: configPath}
//	}
//
//	func (o *oer) env() *oer {
//		if o.val != "" {
//			return o
//		}
//		configEnv := os.Getenv("MY_CONFIG")
//		if configEnv != "" {
//			// yes
//		}
//		o.val = configEnv
//		return o
//	}
//
//	func (o *oer) _default() string {
//		if o.val != "" {
//			return o.val
//		}
//		return "default value"
//	}
//
//	func foo() {
//		configPath := cmdFlag().or.env().or._default()
//		// use configPath
//	}
func Or[T comparable](fns ...func() T) (res T) {
	var zero = res
	for _, fn := range fns {
		if v := fn(); v != zero {
			return v
		}
	}
	return zero
}

// OrUnwished like the function Or but can indicate the unwished value. It will
// traverses through the given function and compares to the unwished value,
// returning it as soon as it gets the value is not equal non-unwished value
// and not equal zero value. However, it will return a zero value when every
// function return the unwished value.
func OrUnwish[T comparable](unwish T, fns ...func() T) (res T) {
	var zero = res
	for _, fn := range fns {
		if v := fn(); v != unwish && v != zero {
			return v
		}
	}
	return zero
}
