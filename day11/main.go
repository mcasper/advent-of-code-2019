package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/mcasper/advent-of-code-2019/shared"
)

type Robot struct {
	x      int
	y      int
	facing string
}

func (r *Robot) Turn(direction int) {
	switch direction {
	case 0:
		switch r.facing {
		case "up":
			r.facing = "left"
		case "left":
			r.facing = "down"
		case "down":
			r.facing = "right"
		case "right":
			r.facing = "up"
		}
	case 1:
		switch r.facing {
		case "up":
			r.facing = "right"
		case "left":
			r.facing = "up"
		case "down":
			r.facing = "left"
		case "right":
			r.facing = "down"
		}
	}
}

func (r *Robot) Move() {
	switch r.facing {
	case "up":
		r.y++
	case "left":
		r.x--
	case "down":
		r.y--
	case "right":
		r.x++
	}
}

type Point struct {
	x     int
	y     int
	color int
}

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	split := strings.Split(strings.Trim(string(input), "\n"), ",")

	if err != nil {
		log.Fatal(err)
	}

	var ints []int

	for _, line := range split {
		integer, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, integer)
	}

	var inputStream bytes.Buffer
	var outputStream bytes.Buffer
	reader := bufio.NewReader(&outputStream)
	var grid []Point

	inputStream.WriteString("0\n")

	computer := shared.Computer{Inputs: ints, Position: 0, RelativeBase: 0, InputStream: &inputStream, OutputStream: &outputStream}

	robot := Robot{x: 0, y: 0, facing: "up"}

	for {
		result := computer.Execute()
		if result == 99 {
			break
		}

		output1 := readOutput(reader)
		output2 := readOutput(reader)

		existingPoint := false
		for i, point := range grid {
			if point.x == robot.x && point.y == robot.y {
				grid[i].color = output1
				existingPoint = true
			}
		}

		if !existingPoint {
			grid = append(grid, Point{x: robot.x, y: robot.y, color: output1})
		}

		robot.Turn(output2)
		robot.Move()

		nextInput := 0
		for _, point := range grid {
			if point.x == robot.x && point.y == robot.y {
				nextInput = point.color
			}
		}

		inputStream.WriteString(fmt.Sprintf("%v\n", nextInput))
	}

	fmt.Printf("Part 1 result: %v\n", len(grid))
}

func readOutput(reader *bufio.Reader) int {
	output, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	output = strings.TrimSuffix(output, "\n")
	i, err := strconv.Atoi(output)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
