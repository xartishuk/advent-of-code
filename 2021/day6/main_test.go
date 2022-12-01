package main

import (
	"testing"
)

func TestLanternfish(t *testing.T) {
	tests := []struct {
		inputFile string
		days      int
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			days:      80,
			expected:  5934,
		},
		{
			inputFile: "input.txt",
			days:      80,
			expected:  350605,
		},
		{
			inputFile: "input_test.txt",
			days:      256,
			expected:  26984457539,
		},
		{
			inputFile: "input.txt",
			days:      256,
			expected:  1592778185024,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := Lanternfish(input, tt.days)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
