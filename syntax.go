package eutil

func If[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
}

// Or executes the given functions and compares will the zero value, it will
// return the not-zero value in once. It is a parody of the OR operator in Python.
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
func Or[T comparable](unwished T, fns ...func() T) T {
	for _, fn := range fns {
		if v := fn(); v != unwished {
			return v
		}
	}
	return unwished
}
