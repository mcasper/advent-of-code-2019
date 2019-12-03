package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

	// Special instructions
	ints[1] = 12
	ints[2] = 2

	result := Execute(ints)
	fmt.Printf("Part 1 result is: %v\n", result)
}

func Execute(ints []int) int {
	position := 0

	for {
		opcode := ints[position]

		if opcode == 99 {
			break
		}

		val1_index := ints[position+1]
		val2_index := ints[position+2]
		dest := ints[position+3]

		if opcode == 1 {
			ints[dest] = ints[val1_index] + ints[val2_index]
		} else if opcode == 2 {
			ints[dest] = ints[val1_index] * ints[val2_index]
		} else {
			log.Fatalf("Don't know what to do with opcode %v\n", opcode)
		}

		position += 4
	}

	return ints[0]
}
