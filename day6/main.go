package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	orbits := strings.Split(strings.Trim(string(input), "\n"), "\n")

	if err != nil {
		log.Fatal(err)
	}

	result := Execute(orbits)
	fmt.Printf("Part 1 result: %v\n", result)
}

func Execute(orbits []string) int {
	orbitTree := map[string][]string{}

	for _, orbit := range orbits {
		split := strings.Split(orbit, ")")
		left := split[0]
		right := split[1]

		if orbitTree[left] != nil {
			orbitTree[left] = append(orbitTree[left], right)
		} else {
			orbitTree[left] = []string{right}
		}
	}

	total := 0
	depth := 1
	stack := orbitTree["COM"]

	for {
		total += (len(stack) * depth)
		tempStack := []string{}

		for _, thing := range stack {
			entries := orbitTree[thing]
			for _, entry := range entries {
				tempStack = append(tempStack, entry)
			}
		}

		if len(tempStack) == 0 {
			break
		}

		depth += 1
		stack = tempStack
	}

	return total
}
