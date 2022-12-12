package main

import "testing"

func TestDeviceSpace(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    31,
		},
		{
			inputFile: "input.txt",
			result:    412,
		},
	}

	for _, tt := range tests {
		start, end, grid, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := HillClimb(start, end, grid)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
