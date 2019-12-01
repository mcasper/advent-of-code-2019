package main

import (
	"testing"
)

func TestExecuteExample(t *testing.T) {
	payloads := []int{1969}
	result := Execute(payloads)
	if result != 654 {
		t.Errorf("Expected 654, got %v\n", result)
	}

	payloads = []int{100756}
	result = Execute(payloads)
	if result != 33583 {
		t.Errorf("Expected 33583, got %v\n", result)
	}

	payloads = []int{1969, 100756}
	result = Execute(payloads)
	if result != 34237 {
		t.Errorf("Expected 34237, got %v\n", result)
	}
}
