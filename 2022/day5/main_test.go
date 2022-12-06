package main

import "testing"

func TestCampCleanup(t *testing.T) {
	tests := []struct {
		inputFile string
		res       string
	}{
		{
			inputFile: "input_test.txt",
			res:       "CMZ",
		},
		{
			inputFile: "input.txt",
			res:       "VQZNJMWTR",
		},
	}

	for _, tt := range tests {
		crates, commands, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		res := SupplyStacks(crates, commands)

		if res != tt.res {
			t.Errorf("expected %s, got %s", tt.res, res)
		}
	}
}
