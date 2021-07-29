package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(1)
	if y != 7 {
		t.Error("Expected 7, got", y)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(2)
	if y != 14 {
		t.Error("Expected 14, got", y)
	}
}

func ExampleYears() {
	fmt.Println(Years(1))
	// Output:
	// 7
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
