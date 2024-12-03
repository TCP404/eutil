package eutil

import (
	"encoding/json"
	"unsafe"
)

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
//
// Deprecated: Use cmp.Or instead.
func Or[T comparable](args ...T) (res T) {
	var zero = res
	for _, v := range args {
		if v != zero {
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
func OrUnwish[T comparable](unwish T, args ...T) (res T) {
	var zero = res
	for _, v := range args {
		if v != unwish && v != zero {
			return v
		}
	}
	return zero
}

// In checks if the first argument is in the rest of the arguments.
// It returns true if the first argument is in the rest of the arguments.
func In[T comparable](ori T, args ...T) bool {
	for _, arg := range args {
		if arg == ori {
			return true
		}
	}
	return false
}

// B2S means slice of byte to string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B means string to byte slice
func S2B(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func DeepCopy[T any](src T, dst *T) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}
	return nil
}
