package main

import "testing"

func TestSonarSweep(t *testing.T) {
	tests := []struct {
		inputFile string
		groupSize int
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			groupSize: 1,
			expected:  7,
		},
		{
			inputFile: "input.txt",
			groupSize: 1,
			expected:  1581,
		},
		{
			inputFile: "input_test.txt",
			groupSize: 3,
			expected:  5,
		},
		{
			inputFile: "input.txt",
			groupSize: 3,
			expected:  1618,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := SonarSweep(input, tt.groupSize)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
