package shared

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Computer struct {
	Inputs       []int
	Position     int
	InputStream  io.Reader
	OutputStream io.Writer
	RelativeBase int
}

type Instruction struct {
	opcode         int
	parameterModes []int
}

func (c *Computer) Execute() int {
	for {
		instruction := intToInstruction(c.Inputs[c.Position])

		if instruction.opcode == 99 {
			return 99
		}

		if instruction.opcode == 1 {
			c.Add(instruction)
		} else if instruction.opcode == 2 {
			c.Multiply(instruction)
		} else if instruction.opcode == 3 {
			foundInput := c.Input(instruction)
			if !foundInput {
				return 0
			}
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
		} else if instruction.opcode == 9 {
			c.AdjustRelativeBase(instruction)
		} else {
			log.Fatalf("Don't know what to do with opcode %v\n", instruction.opcode)
		}
	}

	return c.Inputs[0]
}

func (c *Computer) Add(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.findPositional(instruction, c.Inputs[c.Position+3], 2)

	c.setVal(val1+val2, dest)
	c.Position += 4
}

func (c *Computer) Multiply(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.findPositional(instruction, c.Inputs[c.Position+3], 2)

	c.setVal(val1*val2, dest)
	c.Position += 4
}

func (c *Computer) Input(instruction Instruction) bool {
	dest := c.findPositional(instruction, c.Inputs[c.Position+1], 0)

	// fmt.Print("Waiting for input: ")
	reader := bufio.NewReader(c.InputStream)
	stringInput, err := reader.ReadString('\n')
	if err != nil {
		return false
		log.Fatal(err)
	}
	stringInput = strings.TrimSuffix(stringInput, "\n")
	c.InputStream = reader

	input, err := strconv.Atoi(stringInput)
	if err != nil {
		fmt.Println("Error parsing input")
		log.Fatal(err)
	}

	c.setVal(input, dest)
	c.Position += 2
	return true
}

func (c *Computer) Output(instruction Instruction) {
	value := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	_, err := io.WriteString(c.OutputStream, fmt.Sprintf("%v\n", strconv.Itoa(value)))
	if err != nil {
		log.Fatal(err)
	}

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
	dest := c.findPositional(instruction, c.Inputs[c.Position+3], 2)

	if val1 < val2 {
		c.setVal(1, dest)
	} else {
		c.setVal(0, dest)
	}

	c.Position += 4
}

func (c *Computer) Equals(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)
	val2 := c.findValue(instruction, c.Inputs[c.Position+2], 1)
	dest := c.findPositional(instruction, c.Inputs[c.Position+3], 2)

	if val1 == val2 {
		c.setVal(1, dest)
	} else {
		c.setVal(0, dest)
	}

	c.Position += 4
}

func (c *Computer) AdjustRelativeBase(instruction Instruction) {
	val1 := c.findValue(instruction, c.Inputs[c.Position+1], 0)

	c.RelativeBase += val1
	c.Position += 2
}

func (c *Computer) findValue(instruction Instruction, parameter int, parameterIndex int) int {
	if len(instruction.parameterModes) > parameterIndex {
		if instruction.parameterModes[parameterIndex] == 1 {
			return parameter
		} else if instruction.parameterModes[parameterIndex] == 2 {
			return c.getVal(parameter + c.RelativeBase)
		}
	}

	return c.getVal(parameter)
}

func (c *Computer) findPositional(instruction Instruction, parameter int, parameterIndex int) int {
	if len(instruction.parameterModes) > parameterIndex {
		if instruction.parameterModes[parameterIndex] == 2 {
			return parameter + c.RelativeBase
		}
	}

	return parameter
}

func (c *Computer) setVal(val, index int) {
	if len(c.Inputs) <= index {
		length := len(c.Inputs)
		for i := 0; i < (index - length + 1); i++ {
			c.Inputs = append(c.Inputs, 0)
		}
	}

	c.Inputs[index] = val
}

func (c *Computer) getVal(index int) int {
	if len(c.Inputs) <= index {
		length := len(c.Inputs)
		for i := 0; i < (index - length + 1); i++ {
			c.Inputs = append(c.Inputs, 0)
		}
	}

	return c.Inputs[index]
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
