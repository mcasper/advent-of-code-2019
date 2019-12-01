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
	fmt.Printf("Result is: %v\n", result)
}

func Execute(payloads []int) int {
	var result int
	for _, payload := range payloads {
		fuel_weight := int(math.Floor(float64(payload)/float64(3))) - 2
		result = result + fuel_weight
	}
	return result
}
