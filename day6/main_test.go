package main

import (
	"testing"
)

func TestExecute(t *testing.T) {
	orbits := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	expected := 42
	result := Execute(orbits)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}
