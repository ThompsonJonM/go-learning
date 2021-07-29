package mymath

import "testing"

func TestSum(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		got := Sum(2, 2)
		if got != 4 {
			t.Error("expected 4, got ", got)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		got := Sum(-2, -4, -6, -8)
		if got != -20 {
			t.Error("expected -20, got ", got)
		}
	})

	t.Run("mixed numbers", func(t *testing.T) {
		got := Sum(-2, 4, -6, 10)
		if got != 6 {
			t.Error("expected 6, got ", got)
		}
	})
}
