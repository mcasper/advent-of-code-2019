package shared

import (
	"testing"
)

func TestAdd(t *testing.T) {
	computer := Computer{Inputs: []int{1, 0, 0, 0, 99}, Position: 0}
	result := computer.Execute()

	if result != 2 {
		t.Errorf("Expected %v as Execute result, got %v", 2, result)
	}

	if computer.Position != 4 {
		t.Errorf("Expected Position to be %v, got %v", 4, computer.Position)
	}
}

func TestAddWithParameterGroups(t *testing.T) {
	computer := Computer{Inputs: []int{1101, 50, 50, 0, 99}, Position: 0}
	result := computer.Execute()

	if result != 100 {
		t.Errorf("Expected %v as Execute result, got %v", 100, result)
	}

	if computer.Position != 4 {
		t.Errorf("Expected Position to be %v, got %v", 4, computer.Position)
	}
}

func TestMultiply(t *testing.T) {
	computer := Computer{Inputs: []int{2, 0, 0, 0, 99}, Position: 0}
	result := computer.Execute()

	if result != 4 {
		t.Errorf("Expected %v as Execute result, got %v", 4, result)
	}

	if computer.Position != 4 {
		t.Errorf("Expected Position to be %v, got %v", 4, computer.Position)
	}
}

func TestMultiplyWithParameterGroups(t *testing.T) {
	computer := Computer{Inputs: []int{1102, 50, 50, 0, 99}, Position: 0}
	result := computer.Execute()

	if result != 2500 {
		t.Errorf("Expected %v as Execute result, got %v", 2500, result)
	}

	if computer.Position != 4 {
		t.Errorf("Expected Position to be %v, got %v", 4, computer.Position)
	}
}

func TestComplexInstructions(t *testing.T) {
	computer := Computer{Inputs: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, Position: 0}
	expected := 3500
	result := computer.Execute()

	if result != expected {
		t.Errorf("Expected %v as Execute result, got %v", expected, result)
	}
}

func TestIntToInstruction(t *testing.T) {
	result := intToInstruction(1101)
	expectedOpcode := 1
	expectedParameterModes := []int{1, 1}

	if result.opcode != expectedOpcode {
		t.Errorf("Expected opcode %v, got %v", expectedOpcode, result.opcode)
	}

	if !Equal(result.parameterModes, expectedParameterModes) {
		t.Errorf("Expected parameterModes %v, got %v", expectedParameterModes, result.parameterModes)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
