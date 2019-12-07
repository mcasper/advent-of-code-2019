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

	result = Execute2(orbits)
	fmt.Printf("Part 2 result: %v\n", result)
}

func Execute(orbits []string) int {
	orbitTree := orbitsToTree(orbits)

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

func Execute2(orbits []string) int {
	orbitTree := orbitsToTree(orbits)
	youPath := findPath(orbitTree, "YOU")
	sanPath := findPath(orbitTree, "SAN")

	var youIndex int
	var intersection string
	for i, node := range youPath {
		if contains(sanPath, node) {
			intersection = node
			youIndex = i
			break
		}
	}

	var sanIndex int
	for i, node := range sanPath {
		if node == intersection {
			sanIndex = i
		}
	}

	return sanIndex + youIndex
}

func orbitsToTree(orbits []string) map[string][]string {
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

	return orbitTree
}

func findPath(orbitTree map[string][]string, node string) []string {
	path := []string{}
	currentNode := node

	for {
		for k, v := range orbitTree {
			if contains(v, currentNode) {
				path = append(path, k)
				currentNode = k
			}
		}

		if currentNode == "COM" {
			break
		}
	}

	return path
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
