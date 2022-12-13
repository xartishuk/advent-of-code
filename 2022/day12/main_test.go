package main

import "testing"

func TestHillClimb(t *testing.T) {
	tests := []struct {
		inputFile  string
		AnyAWillDo bool
		result     int
	}{
		{
			inputFile:  "input_test.txt",
			AnyAWillDo: false,
			result:     31,
		},
		{
			inputFile:  "input.txt",
			AnyAWillDo: false,
			result:     412,
		},
		{
			inputFile:  "input_test.txt",
			AnyAWillDo: true,
			result:     29,
		},
		{
			inputFile:  "input.txt",
			AnyAWillDo: true,
			result:     402,
		},
	}

	for _, tt := range tests {
		start, end, grid, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		AnyAWillDo = tt.AnyAWillDo

		result := HillClimb(start, end, grid)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
