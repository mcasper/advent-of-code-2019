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

	ints1 := ints
	ints2 := append(ints1[:0:0], ints1...)

	// Special instructions
	ints1[1] = 12
	ints1[2] = 2

	result := Execute(ints1)
	fmt.Printf("Part 1 result is: %v\n", result)

	result = Execute2(ints2)
	fmt.Printf("Part 2 result is: %v\n", result)
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

func Execute2(ints []int) int {
	expectation := 19690720
	var noun_result int
	var verb_result int

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			ints_copy := append(ints[:0:0], ints...)
			ints_copy[1] = noun
			ints_copy[2] = verb
			result := Execute(ints_copy)
			if result == expectation {
				noun_result = noun
				verb_result = verb
			}
		}
	}

	return (100 * noun_result) + verb_result
}
