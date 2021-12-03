package main

import "testing"

func TestPosition(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  150,
		},
		{
			inputFile: "input.txt",
			expected:  1989014,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := Position(input)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
