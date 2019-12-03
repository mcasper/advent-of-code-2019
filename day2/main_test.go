package main

import (
	"testing"
)

func TestExecuteExample(t *testing.T) {
	ints := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	result := Execute(ints)
	expected := 3500
	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	ints = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	result = Execute(ints)
	expected = 30
	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}
