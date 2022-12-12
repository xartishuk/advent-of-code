package main

import "testing"

func TestMonkeyBusiness(t *testing.T) {
	tests := []struct {
		inputFile string
		rounds    int
		result    int
	}{
		{
			inputFile: "input_test.txt",
			rounds:    20,
			result:    10605,
		},
		{
			inputFile: "input.txt",
			rounds:    20,
			result:    120056,
		},
		{
			inputFile: "input_test.txt",
			rounds:    10000,
			result:    10605,
		},
		{
			inputFile: "input.txt",
			rounds:    10000,
			result:    120056,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := MonkeyBusiness(in, tt.rounds)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
