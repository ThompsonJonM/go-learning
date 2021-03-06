// Package mymath provides simple math solutions
package mymath

// Sum returns the added total of each parameter of TYPE int
func Sum(x ...int) int {
	s := 0
	for _, k := range x {
		s += k
	}

	return s
}

// Difference returns the subtracted total of each parameter of TYPE int
func Difference(x ...int) int {
	s := 0
	for _, k := range x {
		s -= k
	}

	return s
}

// Product returns the multiplied total of each parameter of TYPE int
func Product(x ...int) int {
	f := 1
	for _, v := range x {
		f *= v
	}

	return f
}

// Quotient returns the divided total of each parameter of TYPE int
func Quotient(x, y int) int {
	s := x / y

	return s
}
