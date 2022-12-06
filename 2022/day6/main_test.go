package main

import "testing"

func TestCampCleanup(t *testing.T) {
	tests := []struct {
		inputFile string
		length    int
		res       int
	}{
		{
			inputFile: "input_test1.txt",
			length:    4,
			res:       7,
		}, {
			inputFile: "input_test2.txt",
			length:    4,
			res:       5,
		}, {
			inputFile: "input_test3.txt",
			length:    4,
			res:       6,
		}, {
			inputFile: "input_test4.txt",
			length:    4,
			res:       10,
		}, {
			inputFile: "input_test5.txt",
			length:    4,
			res:       11,
		},
		{
			inputFile: "input.txt",
			length:    4,
			res:       1155,
		}, {
			inputFile: "input_test1.txt",
			length:    14,
			res:       19,
		}, {
			inputFile: "input_test2.txt",
			length:    14,
			res:       23,
		}, {
			inputFile: "input_test3.txt",
			length:    14,
			res:       23,
		}, {
			inputFile: "input_test4.txt",
			length:    14,
			res:       29,
		}, {
			inputFile: "input_test5.txt",
			length:    14,
			res:       26,
		},
		{
			inputFile: "input.txt",
			length:    14,
			res:       2789,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		res := TuningMarker(in, tt.length)

		if res != tt.res {
			t.Errorf("expected %d, got %d", tt.res, res)
		}
	}
}
