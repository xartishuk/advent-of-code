package main

import "testing"

func TestCampCleanup(t *testing.T) {
	tests := []struct {
		inputFile string
		res       int
	}{
		{
			inputFile: "input_test1.txt",
			res:       7,
		}, {
			inputFile: "input_test2.txt",
			res:       5,
		}, {
			inputFile: "input_test3.txt",
			res:       6,
		}, {
			inputFile: "input_test4.txt",
			res:       10,
		}, {
			inputFile: "input_test5.txt",
			res:       11,
		},
		{
			inputFile: "input.txt",
			res:       1155,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		res := TuningMarker(in)

		if res != tt.res {
			t.Errorf("expected %d, got %d", tt.res, res)
		}
	}
}
