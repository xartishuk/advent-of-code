package main

import (
	"testing"
)

func TestSmokeDanger(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  15,
		},
		{
			inputFile: "input.txt",
			expected:  570,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		floor := LinkedFloor(input)

		result := SmokeDanger(floor)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}

func TestSmokeBasins(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  1134,
		},
		{
			inputFile: "input.txt",
			expected:  899392,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		floor := LinkedFloor(input)

		result := SmokeBasins(floor)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
