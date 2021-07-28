package main

import "testing"

func TestFoo(t *testing.T) {
	got := foo(2, 7)
	if got != 9 {
		t.Errorf("Foo(2, 7) = %d; want 9", got)
	}
}

func TestBar(t *testing.T) {
	t.Run("negative", func(t *testing.T) {
		_, err := bar(-19.23)
		if err == nil {
			t.Error("Bar should throw error with negative number")
		}
	})

	t.Run("positive", func(t *testing.T) {
		got, err := bar(10.23)
		if err != nil {
			t.Errorf("bar(10.23) returns %f; want 10.23", got)
		}
	})
}
