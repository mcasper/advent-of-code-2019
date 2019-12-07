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

func TestExecute2(t *testing.T) {
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
		"K)YOU",
		"I)SAN",
	}
	expected := 4
	result := Execute2(orbits)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}

func TestFindPath(t *testing.T) {
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
	orbitTree := orbitsToTree(orbits)
	expected := []string{"J", "E", "D", "C", "B", "COM"}
	result := findPath(orbitTree, "K")

	if !Equal(expected, result) {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
