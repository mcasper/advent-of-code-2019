package main

import (
	"bytes"
	"fmt"
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

	largestOutput := 0

	for phase1 := 0; phase1 < 5; phase1++ {
		for phase2 := 0; phase2 < 5; phase2++ {
			if phase1 == phase2 {
				continue
			}
			for phase3 := 0; phase3 < 5; phase3++ {
				if contains([]int{phase1, phase2}, phase3) {
					continue
				}
				for phase4 := 0; phase4 < 5; phase4++ {
					if contains([]int{phase1, phase2, phase3}, phase4) {
						continue
					}
					for phase5 := 0; phase5 < 5; phase5++ {
						if contains([]int{phase1, phase2, phase3, phase4}, phase5) {
							continue
						}
						phase1Inputs := copyInputs(ints)
						var computer1InputStream bytes.Buffer
						computer1InputStream.WriteString(fmt.Sprintf("%v\n", phase1))
						computer1InputStream.WriteString("0\n")

						phase2Inputs := copyInputs(ints)
						var computer2InputStream bytes.Buffer
						computer2InputStream.WriteString(fmt.Sprintf("%v\n", phase2))

						phase3Inputs := copyInputs(ints)
						var computer3InputStream bytes.Buffer
						computer3InputStream.WriteString(fmt.Sprintf("%v\n", phase3))

						phase4Inputs := copyInputs(ints)
						var computer4InputStream bytes.Buffer
						computer4InputStream.WriteString(fmt.Sprintf("%v\n", phase4))

						phase5Inputs := copyInputs(ints)
						var computer5InputStream bytes.Buffer
						computer5InputStream.WriteString(fmt.Sprintf("%v\n", phase5))

						fmt.Printf("%v%v%v%v%v\n", phase1, phase2, phase3, phase4, phase5)

						computer1 := shared.Computer{
							Inputs:       phase1Inputs,
							Position:     0,
							InputStream:  &computer1InputStream,
							OutputStream: &computer2InputStream,
						}
						computer1.Execute()

						computer2 := shared.Computer{
							Inputs:       phase2Inputs,
							Position:     0,
							InputStream:  &computer2InputStream,
							OutputStream: &computer3InputStream,
						}
						computer2.Execute()

						computer3 := shared.Computer{
							Inputs:       phase3Inputs,
							Position:     0,
							InputStream:  &computer3InputStream,
							OutputStream: &computer4InputStream,
						}
						computer3.Execute()

						computer4 := shared.Computer{
							Inputs:       phase4Inputs,
							Position:     0,
							InputStream:  &computer4InputStream,
							OutputStream: &computer5InputStream,
						}
						computer4.Execute()

						var resultBuffer bytes.Buffer

						computer5 := shared.Computer{
							Inputs:       phase5Inputs,
							Position:     0,
							InputStream:  &computer5InputStream,
							OutputStream: &resultBuffer,
						}
						computer5.Execute()

						result, err := strconv.Atoi(strings.TrimSuffix(resultBuffer.String(), "\n"))
						if err != nil {
							log.Fatal(err)
						}

						if result > largestOutput {
							largestOutput = result
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part 1 result: %v\n", largestOutput)

	// var inputStream bytes.Buffer
	// inputStream.Write([]byte("1\n3\n"))

	// outputStream := bytes.Buffer{}

	// computer := shared.Computer{Inputs: ints, Position: 0, InputStream: &inputStream, OutputStream: &outputStream}
	// computer.Execute()

	// fmt.Printf("OutputStream: %v\n", outputStream.String())
}

func copyInputs(inputs []int) []int {
	return append(inputs[:0:0], inputs...)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
