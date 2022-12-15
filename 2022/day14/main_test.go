package main

import "testing"

func TestSand(t *testing.T) {
	tests := []struct {
		inputFile string
		hasFloor  bool
		result    int
	}{
		{
			inputFile: "input_test.txt",
			hasFloor:  false,
			result:    24,
		},
		{
			inputFile: "input.txt",
			hasFloor:  false,
			result:    745,
		},
		{
			inputFile: "input_test.txt",
			hasFloor:  true,
			result:    93,
		},
		{
			inputFile: "input.txt",
			hasFloor:  true,
			result:    27551,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := Sand(in, tt.hasFloor)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
