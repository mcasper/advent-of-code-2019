package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type Asteroid struct {
	x float64
	y float64
}

func (a Asteroid) Total() float64 {
	return a.x + a.y
}

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")

	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	result := Execute(grid)
	fmt.Printf("Part 1 result: %v\n", result)

	result = Execute2(grid)
	fmt.Printf("Part 2 result: %v\n", result)
}

func Execute(grid [][]string) int {
	asteroids := []Asteroid{}

	for y, row := range grid {
		for x, space := range row {
			if space == "#" {
				asteroids = append(asteroids, Asteroid{float64(x), float64(y)})
			}
		}
	}

	mostVisibleAsteroids := 0
	var mostVisibleAsteroid Asteroid

	for _, sourceAsteroid := range asteroids {
		visibleAsteroids := 0

		for _, destAsteroid := range asteroids {
			if sourceAsteroid == destAsteroid {
				continue
			}

			if !anyBlockingAsteroids(sourceAsteroid, destAsteroid, asteroids) {
				visibleAsteroids++
			}
		}

		if mostVisibleAsteroids < visibleAsteroids {
			mostVisibleAsteroids = visibleAsteroids
			mostVisibleAsteroid = sourceAsteroid
		}
	}

	fmt.Printf("Most visible asteroid %v\n", mostVisibleAsteroid)
	return mostVisibleAsteroids
}

func Execute2(grid [][]string) int {
	asteroids := []Asteroid{}

	for y, row := range grid {
		for x, space := range row {
			if space == "#" {
				asteroids = append(asteroids, Asteroid{float64(x), float64(y)})
			}
		}
	}

	destroyedAsteroids := 0
	lastAngle := float64(89.9999999)
	sourceAsteroid := Asteroid{x: 26, y: 36}

	for {
		visibleAsteroids := []Asteroid{}

		for _, destAsteroid := range asteroids {
			if sourceAsteroid == destAsteroid {
				continue
			}

			if !anyBlockingAsteroids(sourceAsteroid, destAsteroid, asteroids) {
				visibleAsteroids = append(visibleAsteroids, destAsteroid)
			}
		}

		closestAngleDelta := float64(360)
		var closestAsteroid Asteroid
		for _, visibleAsteroid := range visibleAsteroids {
			currentAngle := angle(sourceAsteroid, visibleAsteroid)
			var delta float64
			if lastAngle < currentAngle {
				delta = currentAngle - lastAngle
			} else {
				delta = currentAngle + 360 - lastAngle
			}

			if delta < closestAngleDelta {
				closestAngleDelta = delta
				closestAsteroid = visibleAsteroid
			}
		}

		destroyedAsteroids++

		if destroyedAsteroids == 200 {
			return int((closestAsteroid.x * 100) + closestAsteroid.y)
		}

		var destroyedAsteroidIndex int
		for i, asteroid := range asteroids {
			if asteroid == closestAsteroid {
				destroyedAsteroidIndex = i
			}
		}

		copy(asteroids[destroyedAsteroidIndex:], asteroids[destroyedAsteroidIndex+1:])
		asteroids[len(asteroids)-1] = Asteroid{x: 0, y: 0}
		asteroids = asteroids[:len(asteroids)-1]

		lastAngle = angle(sourceAsteroid, closestAsteroid)
	}
}

func anyBlockingAsteroids(sourceAsteroid Asteroid, destAsteroid Asteroid, asteroids []Asteroid) bool {
	for _, potentiallyBlockingAsteroid := range asteroids {
		if sourceAsteroid == potentiallyBlockingAsteroid || destAsteroid == potentiallyBlockingAsteroid {
			continue
		}

		if asteroidBlocksLOS(sourceAsteroid, destAsteroid, potentiallyBlockingAsteroid) {
			return true
		}
	}

	return false
}

func asteroidBlocksLOS(sourceAsteroid, destAsteroid, potentiallyBlockingAsteroid Asteroid) bool {
	if angle(sourceAsteroid, destAsteroid) == angle(sourceAsteroid, potentiallyBlockingAsteroid) {
		return distance(sourceAsteroid, potentiallyBlockingAsteroid) < distance(sourceAsteroid, destAsteroid)
	}

	return false
}

func angle(a, b Asteroid) float64 {
	return ((math.Atan2(b.y-a.y, b.x-a.x) * 180) / math.Pi) + 180
}

func distance(a, b Asteroid) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}
