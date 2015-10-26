/*
Package vec implements functions that create or act upon 1D slices of
`float64`.
*/
package vec

import (
	"log"
	"math"
	"runtime"
)

// ElementalFn is a function that takes a float64 and returns a
// `float64`. This function can therefore be applied to each element
// of a 2D `float64` slice, and can be used to construct a new one.
type ElementalFn func(float64) float64

// Ones returns a new 1D slice where all the elements are equal to `1.0`.
func Ones(l int) []float64 {
	o := make([]float64, l)
	ApplyInPlace(func(i float64) float64 { return 1.0 }, o)
	return o
}

// Inc returns a 1D slice, where element `[0] == 0.0`, and each
// subsequent element is incremented by `1.0`.
//
// For example, `m := Inc(3)` is
//
// `[1.0, 2.0 3.0]`.
func Inc(l int) []float64 {
	v := make([]float64, l)
	iter := 0
	for i := 0; i < l; i++ {
		v[i] = float64(iter)
		iter++
	}
	return v
}

// Equals checks if two 1D slices have the same length, and contain the same
// entries at each slot.
func Equal(v1, v2 []float64) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}

// Mul returns a new 1D slice that is the result of element-wise multiplication
// of two 1D slices.
func Mul(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		msg := "vec.%v Error: in %v [%v line %v].\n"
		msg += "Length of the first 1D slice is %v, length of the second 1D slice\n"
		msg += "is %v. They must match.\n"
		p, f, l, _ := runtime.Caller(1)
		log.Fatalf(msg, "Mul", f, runtime.FuncForPC(p).Name(), l, len(v1), len(v2))
	}
	o := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		o[i] = v1[i] * v2[i]
	}
	return o
}

// Add returns a new 1D slice that is the result of element-wise addition
// of two 1D slices.
func Add(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		msg := "vec.%v Error: in %v [%v line %v].\n"
		msg += "Length of the first 1D slice is %v, length of the second 1D slice\n"
		msg += "is %v. They must match.\n"
		p, f, l, _ := runtime.Caller(1)
		log.Fatalf(msg, "Add", f, runtime.FuncForPC(p).Name(), l, len(v1), len(v2))
	}
	o := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		o[i] = v1[i] + v2[i]
	}
	return o
}

// Sub returns a new 1D slice that is the result of element-wise subtraction
// of two 1D slices.
func Sub(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		msg := "vec.%v Error: in %v [%v line %v].\n"
		msg += "Length of the first 1D slice is %v, length of the second 1D slice\n"
		msg += "is %v. They must match.\n"
		p, f, l, _ := runtime.Caller(1)
		log.Fatalf(msg, "Sub", f, runtime.FuncForPC(p).Name(), l, len(v1), len(v2))
	}
	o := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		o[i] = v1[i] - v2[i]
	}
	return o
}

// Div returns a new 1D slice that is the result of element-wise division
// of two 1D slices. If any elements in the 2nd 1D slice are 0, then this
// function call aborts.
func Div(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		msg := "vec.%v Error: in %v [%v line %v].\n"
		msg += "Length of the first 1D slice is %v, length of the second 1D slice\n"
		msg += "is %v. They must match.\n"
		p, f, l, _ := runtime.Caller(1)
		log.Fatalf(msg, "Div", f, runtime.FuncForPC(p).Name(), l, len(v1), len(v2))
	}
	o := make([]float64, len(v1))
	for i := 0; i < len(v1); i++ {
		if v2[i] == 0.0 {
			msg := "vec.%v Error: in %v [%v line %v].\n"
			msg += "Entry %v in the second slice is 0.0. Cannot devide by 0.0\n"
			p, f, l, _ := runtime.Caller(1)
			log.Fatalf(msg, "Div", f, runtime.FuncForPC(p).Name(), l, i)
		}
		o[i] = v1[i] * v2[i]
	}
	return o
}

// ApplyInPlace calls a given elemental function on each Element of a 1D slice,
// returning it afterwards. This function modifies the original 1D slice. If
// a non-mutating operation is desired, use the "Apply" function instead.
func ApplyInPlace(f ElementalFn, v []float64) {
	for i := 0; i < len(v); i++ {
		v[i] = f(v[i])
	}
}

// Apply created a new 1D slice which is populated throw applying the given
// function to the corresponding entries of a given 1D slice. This function
// does not modify its arguments, instead allocating a new 1D slice to
// contain the result. This is a performance hit. If you are OK with mutating
// the original vector, then use the "ApllyInPlace" function instead.
func Apply(f ElementalFn, v []float64) []float64 {
	o := make([]float64, len(v))
	for i := 0; i < len(v); i++ {
		o[i] = f(v[i])
	}
	return o
}

// Dot is the inner product of two 1D slices of `float64`.
func Dot(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		msg := "vec.%v Error: in %v [%v line %v].\n"
		msg += "Length of the first 1D slice is %v, length of the second 1D slice\n"
		msg += "is %v. They must match.\n"
		p, f, l, _ := runtime.Caller(1)
		log.Fatalf(msg, "Dot", f, runtime.FuncForPC(p).Name(), l, len(v1), len(v2))
	}
	var o float64
	for i := 0; i < len(v1); i++ {
		o += v1[i] * v2[i]
	}
	return o
}

// Reset sets the values of all entries in a 2D slice of `float64` to `0.0`.
func Reset(v []float64) {
	ApplyInPlace(func(i float64) float64 { return 0.0 }, v)
	return
}

// Sum returns the sum of the elements of a 1D slice of `float64`.
func Sum(v []float64) float64 {
	var o float64
	for i := 0; i < len(v); i++ {
		o += v[i]
	}
	return o
}

// Norm calculated the norm of a given 1D slice. This is the Euclidean length
// of the slice.
func Norm(v []float64) float64 {
	return math.Sqrt(Sum(Apply(func(i float64) float64 { return i * i }, v)))
}