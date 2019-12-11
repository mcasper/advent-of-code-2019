package main

import (
	"testing"
)

func TestExecute(t *testing.T) {
	grid := [][]string{
		[]string{
			".", "#", ".", ".", "#",
		},
		[]string{
			".", ".", ".", ".", ".",
		},
		[]string{
			"#", "#", "#", "#", "#",
		},
		[]string{
			".", ".", ".", ".", "#",
		},
		[]string{
			".", ".", ".", "#", "#",
		},
	}
	expected := 8
	result := Execute(grid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	grid = [][]string{
		[]string{
			".", ".", ".", ".", ".", ".", "#", ".", "#", ".",
		},
		[]string{
			"#", ".", ".", "#", ".", "#", ".", ".", ".", ".",
		},
		[]string{
			".", ".", "#", "#", "#", "#", "#", "#", "#", ".",
		},
		[]string{
			".", "#", ".", "#", ".", "#", "#", "#", ".", ".",
		},
		[]string{
			".", "#", ".", ".", "#", ".", ".", ".", ".", ".",
		},
		[]string{
			".", ".", "#", ".", ".", ".", ".", "#", ".", "#",
		},
		[]string{
			"#", ".", ".", "#", ".", ".", ".", ".", "#", ".",
		},
		[]string{
			".", "#", "#", ".", "#", ".", ".", "#", "#", "#",
		},
		[]string{
			"#", "#", ".", ".", ".", "#", ".", ".", "#", ".",
		},
		[]string{
			".", "#", ".", ".", ".", ".", "#", "#", "#", "#",
		},
	}

	expected = 33
	result = Execute(grid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	grid = [][]string{
		[]string{
			"#", ".", "#", ".", ".", ".", "#", ".", "#", ".",
		},
		[]string{
			".", "#", "#", "#", ".", ".", ".", ".", "#", ".",
		},
		[]string{
			".", "#", ".", ".", ".", ".", "#", ".", ".", ".",
		},
		[]string{
			"#", "#", ".", "#", ".", "#", ".", "#", ".", "#",
		},
		[]string{
			".", ".", ".", ".", "#", ".", "#", ".", "#", ".",
		},
		[]string{
			".", "#", "#", ".", ".", "#", "#", "#", ".", "#",
		},
		[]string{
			".", ".", "#", ".", ".", ".", "#", "#", ".", ".",
		},
		[]string{
			".", ".", "#", "#", ".", ".", ".", ".", "#", "#",
		},
		[]string{
			".", ".", ".", ".", ".", ".", "#", ".", ".", ".",
		},
		[]string{
			".", "#", "#", "#", "#", ".", "#", "#", "#", ".",
		},
	}

	expected = 35
	result = Execute(grid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	grid = [][]string{
		[]string{
			".", "#", ".", ".", "#", ".", ".", "#", "#", "#",
		},
		[]string{
			"#", "#", "#", "#", ".", "#", "#", "#", ".", "#",
		},
		[]string{
			".", ".", ".", ".", "#", "#", "#", ".", "#", ".",
		},
		[]string{
			".", ".", "#", "#", "#", ".", "#", "#", ".", "#",
		},
		[]string{
			"#", "#", ".", "#", "#", ".", "#", ".", "#", ".",
		},
		[]string{
			".", ".", ".", ".", "#", "#", "#", ".", ".", "#",
		},
		[]string{
			".", ".", "#", ".", "#", ".", ".", "#", ".", "#",
		},
		[]string{
			"#", ".", ".", "#", ".", "#", ".", "#", "#", "#",
		},
		[]string{
			".", "#", "#", ".", ".", ".", "#", "#", ".", "#",
		},
		[]string{
			".", ".", ".", ".", ".", "#", ".", "#", ".", ".",
		},
	}

	expected = 41
	result = Execute(grid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}

func TestAsteroidBlocksLOS(t *testing.T) {
	sourceAsteroid := Asteroid{x: 1, y: 0}
	destAsteroid := Asteroid{x: 0, y: 2}
	potentiallyBlockingAsteroid := Asteroid{x: 1, y: 2}
	expected := false
	result := asteroidBlocksLOS(sourceAsteroid, destAsteroid, potentiallyBlockingAsteroid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	sourceAsteroid = Asteroid{x: 3, y: 4}
	destAsteroid = Asteroid{x: 1, y: 0}
	potentiallyBlockingAsteroid = Asteroid{x: 2, y: 2}
	expected = true
	result = asteroidBlocksLOS(sourceAsteroid, destAsteroid, potentiallyBlockingAsteroid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}

func TestAngle(t *testing.T) {
	sourceAsteroid := Asteroid{x: 2, y: 2}
	destAsteroid := Asteroid{x: 2, y: 1}
	expected := 0.0
	result := angle(sourceAsteroid, destAsteroid)

	if result != expected {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}
