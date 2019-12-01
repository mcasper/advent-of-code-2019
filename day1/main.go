package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	split := strings.Split(strings.Trim(string(input), "\n"), "\n")

	if err != nil {
		log.Fatal(err)
	}

	var payloads []int

	for _, line := range split {
		integer, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		payloads = append(payloads, integer)
	}

	result := Execute(payloads)
	fmt.Printf("Part 1 result is: %v\n", result)

	result2 := Execute2(payloads)
	fmt.Printf("Part 2 result is: %v\n", result2)
}

func Execute(payloads []int) int {
	var result int
	for _, payload := range payloads {
		fuel_weight := int(math.Floor(float64(payload)/float64(3))) - 2
		result = result + fuel_weight
	}
	return result
}

func Execute2(payloads []int) int {
	var result int
	var remaining_mass int
	for _, payload := range payloads {
		remaining_mass = payload
		for remaining_mass > 0 {
			fuel_weight := Execute([]int{remaining_mass})
			if fuel_weight > 0 {
				result = result + fuel_weight
			}
			remaining_mass = fuel_weight
		}
	}
	return result
}
