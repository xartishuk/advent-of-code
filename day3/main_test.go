package main

import "testing"

func TestPowerConsumption(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  198,
		},
		{
			inputFile: "input.txt",
			expected:  0,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := PowerConsumption(input)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
