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
			expected:  4147524,
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

func TestLifeSupport(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int64
	}{
		{
			inputFile: "input_test.txt",
			expected:  230,
		},
		{
			inputFile: "input.txt",
			expected:  3570354,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := LifeSupport(input)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
