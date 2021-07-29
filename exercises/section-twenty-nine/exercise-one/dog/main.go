// Package dog allows us to more fully understad dogs
package dog

// Years converts human years to dog years.
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	c := 0
	for i := 0; i < n; i++ {
		c += 7
	}

	return c
}
