// Package mymath provides simple math solutions
package mymath

// Sum returns the added total of each parameter of TYPE int
func Sum(x ...int) int {
	s := 0
	for _, k := range x {
		s = s + k
	}

	return s
}

// Difference returns the subtracted total of each parameter of TYPE int
func Difference(x ...int) int {
	s := 0
	for _, k := range x {
		s = s - k
	}

	return s
}

// Product returns the multiplied total of each parameter of TYPE int
func Product(x ...int) int {
	s := 0
	for _, k := range x {
		s = s * k
	}

	return s
}

// Quotient returns the divided total of each parameter of TYPE int
func Quotient(x ...int) int {
	s := 0
	for _, k := range x {
		s = s / k
	}

	return s
}
