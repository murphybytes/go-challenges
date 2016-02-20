package main

import "testing"

func TestFib(t *testing.T) {
	result := fib(0, 1, 0, 0)
	if result != 0 {
		t.Errorf("Expected %d got %d\n", 0, result)
	}

	result = fib(0, 1, 0, 1)
	if result != 1 {
		t.Errorf("Expected %d got %d\n", 1, result)
	}

	result = fib(0, 1, 0, 2)
	if result != 1 {
		t.Errorf("Expected %d got %d\n", 1, result)
	}

	result = fib(0, 1, 0, 6)
	if result != 8 {
		t.Errorf("Expected %d got %d\n", 8, result)
	}

	result = fib(0, 1, 0, 43)
	if result != 433494437 {
		t.Errorf("Expected %d got %d\n", 8, result)
	}
}
