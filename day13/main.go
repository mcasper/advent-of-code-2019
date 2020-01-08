package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mcasper/advent-of-code-2019/shared"
)

type Screen struct {
	pixels            [][]string
	outputInstruction []int
	score             int
}

func (s Screen) Print() {
	fmt.Println("\033[2J")

	for _, line := range s.pixels {
		fmt.Printf("%v\n", strings.Join(line, ""))
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Print("\n\n\n\n")
}

func (s Screen) FindCoord(tile string) []int {
	for y, line := range s.pixels {
		for x, pixel := range line {
			if pixel == tile {
				return []int{x, y}
			}
		}
	}
	return []int{}
}

func (s Screen) CountBlocks() int {
	result := 0
	for _, line := range s.pixels {
		for _, pixel := range line {
			if pixel == "X" {
				result++
			}
		}
	}
	return result
}

func (s *Screen) Read(p []byte) (n int, err error) {
	paddleCoords := s.FindCoord("_")
	ballCoords := s.FindCoord("*")

	var input string
	if ballCoords[0] < paddleCoords[0] {
		input = "-1\n"
	} else if ballCoords[0] == paddleCoords[0] {
		input = "0\n"
	} else {
		input = "1\n"
	}

	for i, char := range input {
		p[i] = byte(char)
	}
	return len(input), nil
}

func (s *Screen) Write(p []byte) (n int, err error) {
	thing := strings.Trim(string(p), "\n")
	i, err := strconv.Atoi(thing)

	if err != nil {
		return 0, err
	}

	s.outputInstruction = append(s.outputInstruction, i)

	if len(s.outputInstruction) == 3 {
		x := s.outputInstruction[0]
		y := s.outputInstruction[1]
		tileId := s.outputInstruction[2]

		if x == -1 && y == 0 {
			fmt.Printf("Score is %v\n\n\n\n", tileId)
			s.score = tileId
		} else {
			var newStuff bool

			switch tileId {
			case 0:
				newStuff = s.WritePixel(y, x, " ")
			case 1:
				newStuff = s.WritePixel(y, x, "|")
			case 2:
				newStuff = s.WritePixel(y, x, "X")
			case 3:
				newStuff = s.WritePixel(y, x, "_")
			case 4:
				newStuff = s.WritePixel(y, x, "*")
			}

			if newStuff {
				s.Print()
			}
		}

		s.outputInstruction = []int{}
	}

	return len(p), nil
}

func (s *Screen) WritePixel(y int, x int, pixel string) bool {
	if len(s.pixels) <= y {
		length := len(s.pixels)
		for i := 0; i < (y - length + 1); i++ {
			s.pixels = append(s.pixels, []string{})
		}
	}

	if len(s.pixels[y]) <= x {
		length := len(s.pixels[y])
		for i := 0; i < (x - length + 1); i++ {
			s.pixels[y] = append(s.pixels[y], " ")
		}
	}

	if s.pixels[y][x] == pixel {
		s.pixels[y][x] = pixel
		return false
	} else {
		s.pixels[y][x] = pixel
		return true
	}
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

	screen := Screen{pixels: [][]string{}, outputInstruction: []int{}, score: 0}

	computer := shared.Computer{Inputs: ints, Position: 0, RelativeBase: 0, InputStream: &screen, OutputStream: &screen}

	computer.Execute()

	fmt.Printf("Part 1 result: %v\n", screen.CountBlocks())
	fmt.Printf("Part 2 result: %v\n", screen.score)
}
