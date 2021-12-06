package main

import (
	"testing"
)

func TestHydrothermal(t *testing.T) {
	tests := []struct {
		inputFile        string
		filterOrthogonal bool
		expected         int
	}{
		{
			inputFile:        "input_test.txt",
			filterOrthogonal: true,
			expected:         5,
		},
		{
			inputFile:        "input.txt",
			filterOrthogonal: true,
			expected:         6225,
		},
		{
			inputFile:        "input_test.txt",
			filterOrthogonal: false,
			expected:         12,
		},
		{
			inputFile:        "input.txt",
			filterOrthogonal: false,
			expected:         22116,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := Hydrothermal(input, tt.filterOrthogonal)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
