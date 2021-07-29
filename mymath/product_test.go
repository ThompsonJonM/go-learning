package mymath

import "testing"

func TestProduct(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{5, 8}, 40},
		test{[]int{-2, -4, -10}, -80},
		test{[]int{2, -2}, -4},
	}

	for _, test := range tests {
		x := Product(test.data...)
		if x != test.answer {
			t.Errorf("Expected %v, got %v", test.answer, x)
		}
	}
}
