package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mcasper/advent-of-code-2019/shared"
)

type Screen struct {
	pixels            [][]string
	outputInstruction []int
}

func (s Screen) Print() {
	for _, line := range s.pixels {
		fmt.Printf("%v\n", strings.Join(line, ""))
	}
	fmt.Println("\n\n\n")
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

		switch tileId {
		case 0:
			s.WritePixel(y, x, " ")
		case 1:
			s.WritePixel(y, x, "|")
		case 2:
			s.WritePixel(y, x, "X")
		case 3:
			s.WritePixel(y, x, "_")
		case 4:
			s.WritePixel(y, x, "*")
		}

		s.outputInstruction = []int{}
		s.Print()
	}

	return len(p), nil
}

func (s *Screen) WritePixel(y int, x int, pixel string) {
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

	s.pixels[y][x] = pixel
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

	screen := Screen{pixels: [][]string{}, outputInstruction: []int{}}

	computer := shared.Computer{Inputs: ints, Position: 0, RelativeBase: 0, InputStream: os.Stdin, OutputStream: &screen}

	computer.Execute()

	fmt.Printf("Part 1 result: %v\n", screen.CountBlocks())
}
