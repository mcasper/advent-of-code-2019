package shared

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	Inputs   []int
	Position int
}

type Instruction struct {
	opcode         int
	parameterModes []int
}

func (c *Computer) Execute() int {
	for {
		instruction := intToInstruction(c.Inputs[c.Position])

		if instruction.opcode == 99 {
			break
		}

		if instruction.opcode == 1 {
			c.Add(instruction)
		} else if instruction.opcode == 2 {
			c.Multiply(instruction)
		} else if instruction.opcode == 3 {
			c.Input()
		} else if instruction.opcode == 4 {
			c.Output(instruction)
		} else if instruction.opcode == 5 {
			c.JumpIfTrue(instruction)
		} else if instruction.opcode == 6 {
			c.JumpIfFalse(instruction)
		} else if instruction.opcode == 7 {
			c.LessThan(instruction)
		} else if instruction.opcode == 8 {
			c.Equals(instruction)
		} else {
			log.Fatalf("Don't know what to do with opcode %v\n", instruction.opcode)
		}
	}

	return c.Inputs[0]
}

func (c *Computer) Add(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.Inputs[c.Position+3]

	c.Inputs[dest] = val1 + val2
	c.Position += 4
}

func (c *Computer) Multiply(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.Inputs[c.Position+3]

	c.Inputs[dest] = val1 * val2
	c.Position += 4
}

func (c *Computer) Input() {
	dest := c.Inputs[c.Position+1]

	fmt.Print("Waiting for input: ")
	reader := bufio.NewReader(os.Stdin)
	stringInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	stringInput = strings.TrimSuffix(stringInput, "\n")
	if len(stringInput) > 1 {
		log.Fatal("Too much input")
	}

	input, err := strconv.Atoi(stringInput)
	if err != nil {
		log.Fatal(err)
	}

	c.Inputs[dest] = input
	c.Position += 2
}

func (c *Computer) Output(instruction Instruction) {
	value := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	fmt.Printf("Output: %v\n", value)

	c.Position += 2
}

func (c *Computer) JumpIfTrue(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)

	if val1 != 0 {
		c.Position = val2
	} else {
		c.Position += 3
	}
}

func (c *Computer) JumpIfFalse(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)

	if val1 == 0 {
		c.Position = val2
	} else {
		c.Position += 3
	}
}

func (c *Computer) LessThan(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.Inputs[c.Position+3]

	if val1 < val2 {
		c.Inputs[dest] = 1
	} else {
		c.Inputs[dest] = 0
	}

	c.Position += 4
}

func (c *Computer) Equals(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.Inputs[c.Position+3]

	if val1 == val2 {
		c.Inputs[dest] = 1
	} else {
		c.Inputs[dest] = 0
	}

	c.Position += 4
}

func (c *Computer) findValue(instruction Instruction, parameter int, parameterIndex int) int {
	if len(instruction.parameterModes) > parameterIndex && instruction.parameterModes[parameterIndex] == 1 {
		return parameter
	} else {
		return c.Inputs[parameter]
	}
}

func intToInstruction(integer int) Instruction {
	stringInstruction := strconv.Itoa(integer)
	splitInstruction := strings.Split(stringInstruction, "")
	var stringOpcode string

	if len(splitInstruction) == 1 {
		stringOpcode = splitInstruction[0]
	} else {
		stringOpcode = strings.Join(splitInstruction[len(splitInstruction)-2:], "")
	}

	opcode, err := strconv.Atoi(stringOpcode)
	if err != nil {
		log.Fatal(err)
	}

	parameterModes := []int{}

	if len(splitInstruction) > 2 {
		for i := (len(splitInstruction) - 3); i >= 0; i-- {
			parameterMode, err := strconv.Atoi(splitInstruction[i])
			if err != nil {
				log.Fatal(err)
			}
			parameterModes = append(parameterModes, parameterMode)
		}
	}

	return Instruction{opcode, parameterModes}
}
