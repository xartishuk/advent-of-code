package main

import (
	"testing"
)

func TestSmokeBasin(t *testing.T) {
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

		result := SmokeBasin(input)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
