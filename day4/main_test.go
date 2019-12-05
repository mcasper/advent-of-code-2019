package main

import (
	"testing"
)

func TestValidPassword(t *testing.T) {
	password := 223450
	result := validPassword(password)

	if result {
		t.Errorf("Expected %v to not be a valid password\n", password)
	}

	password = 123789
	result = validPassword(password)

	if result {
		t.Errorf("Expected %v to not be a valid password\n", password)
	}

	password = 122345
	result = validPassword(password)

	if !result {
		t.Errorf("Expected %v to be a valid password\n", password)
	}

	password = 112233
	result = validPassword(password)

	if !result {
		t.Errorf("Expected %v to be a valid password\n", password)
	}

	password = 123444
	result = validPassword(password)

	if result {
		t.Errorf("Expected %v to not be a valid password\n", password)
	}
}
