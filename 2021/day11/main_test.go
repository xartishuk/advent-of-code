package main

import "testing"

func TestOctopusFlashes(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    1656,
		},
		{
			inputFile: "input.txt",
			result:    1785,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := OctopusFlashes(in)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
