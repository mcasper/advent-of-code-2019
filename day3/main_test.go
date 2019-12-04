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
