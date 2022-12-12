package main

import "testing"

func TestMonkeyBusiness(t *testing.T) {
	tests := []struct {
		inputFile   string
		reduceWorry bool
		rounds      int
		result      int
	}{
		{
			inputFile:   "input_test.txt",
			reduceWorry: true,
			rounds:      20,
			result:      10605,
		},
		{
			inputFile:   "input.txt",
			reduceWorry: true,
			rounds:      20,
			result:      120056,
		},
		{
			inputFile:   "input_test.txt",
			reduceWorry: false,
			rounds:      10000,
			result:      2713310158,
		},
		{
			inputFile:   "input.txt",
			reduceWorry: false,
			rounds:      10000,
			result:      21816744824,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile, tt.reduceWorry)
		if err != nil {
			t.Error(err)
		}

		result := MonkeyBusiness(in, tt.rounds)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
