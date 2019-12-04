package main

import (
	"testing"
)

func TestExecuteExample(t *testing.T) {
	wire1 := []string{"R8", "U5", "L5", "D3"}
	wire2 := []string{"U7", "R6", "D4", "L4"}
	expected := 6
	result := Execute(wire1, wire2)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestExecute2Example(t *testing.T) {
	wire1 := []string{"R8", "U5", "L5", "D3"}
	wire2 := []string{"U7", "R6", "D4", "L4"}
	expected := 30
	result := Execute2(wire1, wire2)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	wire1 = []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wire2 = []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	expected = 610
	result = Execute2(wire1, wire2)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInstructionsToCoords(t *testing.T) {
	instructions := []string{"R2", "U2", "L2", "D2"}
	expected := []Coord{
		Coord{x: 1, y: 0},
		Coord{x: 2, y: 0},
		Coord{x: 2, y: 1},
		Coord{x: 2, y: 2},
		Coord{x: 1, y: 2},
		Coord{x: 0, y: 2},
		Coord{x: 0, y: 1},
		Coord{x: 0, y: 0},
	}
	result := instructionsToCoords(instructions)
	if !Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Equal(a, b []Coord) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.x != b[i].x || v.y != b[i].y {
			return false
		}
	}
	return true
}
