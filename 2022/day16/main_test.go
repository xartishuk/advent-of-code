package main

import "testing"

func TestPressureReleaseSolo(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    1651,
		},
		{
			inputFile: "input.txt",
			result:    1789,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := PressureReleaseSolo(in)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}

func TestPressureReleaseDuet(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    1707,
		},
		{
			inputFile: "input.txt",
			result:    1,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := PressureReleaseDuet(in)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
