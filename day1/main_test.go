package main

import "testing"

func TestSonarSweep(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  7,
		},
		{
			inputFile: "input.txt",
			expected:  1581,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := SonarSweep(input)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
