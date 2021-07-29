package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	c := Count("This is a string")
	if c != 4 {
		t.Error("Expected 4, got", c)
	}
}

func TestUseCount(t *testing.T) {

	type test struct {
		data   string
		answer int
	}

	tests := []test{
		{data: "this", answer: 2},
		{data: "string", answer: 1},
		{data: "is", answer: 1},
		{data: "long", answer: 1},
	}

	uc := UseCount("this string is this long")
	for _, pt := range tests {
		if uc[pt.data] != pt.answer {
			t.Errorf("Expected %v, got %v", pt.answer, uc)
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("Test String, Please Ignore"))
	// Output:
	// 4
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("This is a benchmark")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("test string please ignore")
	}
}
