package main

import "testing"

func TestCrateMover9000(t *testing.T) {
	tests := []struct {
		inputFile  string
		crateMover CrateMover
		res        string
	}{
		{
			inputFile:  "input_test.txt",
			crateMover: CrateMover9000{},
			res:        "CMZ",
		},
		{
			inputFile:  "input.txt",
			crateMover: CrateMover9000{},
			res:        "VQZNJMWTR",
		},
		{
			inputFile:  "input_test.txt",
			crateMover: CrateMover9001{},
			res:        "MCD",
		},
		{
			inputFile:  "input.txt",
			crateMover: CrateMover9001{},
			res:        "NLCDCLVMQ",
		},
	}

	for _, tt := range tests {
		crates, commands, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		res := SupplyStacks(tt.crateMover, crates, commands)

		if res != tt.res {
			t.Errorf("expected %s, got %s", tt.res, res)
		}
	}
}
