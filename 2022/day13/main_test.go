package main

import "testing"

func TestDistressSignalPairs(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    13,
		},
		{
			inputFile: "input.txt",
			result:    6623,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := DistressSignalPairs(in)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}

func TestDistressSignalDecoder(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    140,
		},
		{
			inputFile: "input.txt",
			result:    23049,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := DistressSignalDecoder(in)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
