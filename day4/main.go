package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	passwordCount := matchingPasswords(347312, 805915)
	fmt.Printf("Part 1: %v\n", passwordCount)
}

func matchingPasswords(bottom, top int) int {
	validPasswordCount := 0
	for number := bottom; number <= top; number++ {
		if validPassword(number) {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func validPassword(password int) bool {
	stringPassword := strconv.Itoa(password)
	hasIdenticalPair := false
	digitsNeverDecrease := true

	chars := strings.Split(stringPassword, "")
	ints := []int{}
	for _, char := range chars {
		i, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}

	for i, c := range chars {
		if i == (len(stringPassword) - 1) {
			break
		}

		if c == charAt(chars, i+1) && c != charAt(chars, i-1) && c != charAt(chars, i+2) {
			hasIdenticalPair = true
		}

		if ints[i+1] < ints[i] {
			digitsNeverDecrease = false
		}
	}

	return hasIdenticalPair && digitsNeverDecrease
}

func charAt(chars []string, index int) string {
	if index >= 0 && index < len(chars) {
		return chars[index]
	}
	return ""
}
