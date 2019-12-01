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

func TestExecute2Example(t *testing.T) {
	payloads := []int{1969}
	result := Execute2(payloads)
	if result != 966 {
		t.Errorf("Expected 966, got %v\n", result)
	}

	payloads = []int{100756}
	result = Execute2(payloads)
	if result != 50346 {
		t.Errorf("Expected 50346, got %v\n", result)
	}

	payloads = []int{1969, 100756}
	result = Execute2(payloads)
	if result != 51312 {
		t.Errorf("Expected 51312, got %v\n", result)
	}
}
