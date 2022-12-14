package main

import "testing"

func TestCavePaths(t *testing.T) {
	tests := []struct {
		inputFile       string
		allowRepeatOnce bool
		result          int
	}{
		{
			inputFile:       "input_test_1.txt",
			allowRepeatOnce: false,
			result:          10,
		},
		{
			inputFile:       "input_test_2.txt",
			allowRepeatOnce: false,
			result:          19,
		},
		{
			inputFile:       "input_test_3.txt",
			allowRepeatOnce: false,
			result:          226,
		},
		{
			inputFile:       "input.txt",
			allowRepeatOnce: false,
			result:          4773,
		},
		{
			inputFile:       "input_test_1.txt",
			allowRepeatOnce: true,
			result:          36,
		},
		{
			inputFile:       "input_test_2.txt",
			allowRepeatOnce: true,
			result:          103,
		},
		{
			inputFile:       "input_test_3.txt",
			allowRepeatOnce: true,
			result:          3509,
		},
		{
			inputFile:       "input.txt",
			allowRepeatOnce: true,
			result:          116985,
		},
	}

	for _, tt := range tests {
		start, end, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := CavePaths(start, end, tt.allowRepeatOnce)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
