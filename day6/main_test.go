package main

import (
	"testing"
)

func TestLanternfish(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  5934,
		},
		{
			inputFile: "input.txt",
			expected:  350605,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := Lanternfish(input)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
