package main

import "testing"

func TestCavePaths(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test_1.txt",
			result:    10,
		},
		{
			inputFile: "input_test_2.txt",
			result:    19,
		},
		{
			inputFile: "input_test_3.txt",
			result:    226,
		},
		{
			inputFile: "input.txt",
			result:    4773,
		},
	}

	for _, tt := range tests {
		start, end, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := CavePaths(start, end)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
