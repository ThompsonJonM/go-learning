package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	ca := CenteredAvg([]int{1, 4, 6, 8, 100})
	if ca != 6 {
		t.Error("Expected 6, got", ca)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 10})
	}
}
