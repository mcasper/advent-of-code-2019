package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(strings.Trim(string(input), "\n"), "\n")

	wire1 := strings.Split(split[0], ",")
	wire2 := strings.Split(split[1], ",")

	result := Execute(wire1, wire2)
	fmt.Printf("Part 1 result: %v\n", result)
}

func Execute(wire1 []string, wire2 []string) int {
	wire1Coords := instructionsToCoords(wire1)
	wire2Coords := instructionsToCoords(wire2)
	intersections := extractIntersections(wire1Coords, wire2Coords)

	minDistance := 100000000
	for _, coord := range intersections {
		distance := coordinateDistance(Coord{x: 0, y: 0}, coord)
		if minDistance > distance {
			minDistance = distance
		}
	}

	return minDistance
}

func coordinateDistance(coord1 Coord, coord2 Coord) int {
	return int(math.Abs(float64(coord1.x-coord2.x))) + int(math.Abs(float64(coord1.y-coord2.y)))
}

func extractIntersections(coords1, coords2 []Coord) []Coord {
	intersections := []Coord{}

	for _, wire1Coord := range coords1 {
		for _, wire2Coord := range coords2 {
			if wire1Coord.x == wire2Coord.x && wire1Coord.y == wire2Coord.y {
				intersections = append(intersections, wire1Coord)
			}
		}
	}

	return intersections
}

func instructionsToCoords(instructions []string) []Coord {
	coords := []Coord{}

	last_coord := Coord{x: 0, y: 0}

	for _, instruction := range instructions {
		direction := instruction[0]
		value, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 'U':
			for y := 1; y <= value; y++ {
				dest_coord := Coord{x: last_coord.x, y: last_coord.y + 1}
				coords = append(coords, dest_coord)
				last_coord = dest_coord
			}
		case 'D':
			for y := 1; y <= value; y++ {
				dest_coord := Coord{x: last_coord.x, y: last_coord.y - 1}
				coords = append(coords, dest_coord)
				last_coord = dest_coord
			}
		case 'L':
			for x := 1; x <= value; x++ {
				dest_coord := Coord{x: last_coord.x - 1, y: last_coord.y}
				coords = append(coords, dest_coord)
				last_coord = dest_coord
			}
		case 'R':
			for x := 1; x <= value; x++ {
				dest_coord := Coord{x: last_coord.x + 1, y: last_coord.y}
				coords = append(coords, dest_coord)
				last_coord = dest_coord
			}
		}
	}

	return coords
}
