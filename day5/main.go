package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/mcasper/advent-of-code-2019/shared"
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

	computer := shared.Computer{Inputs: ints, Position: 0}
	computer.Execute()
}
