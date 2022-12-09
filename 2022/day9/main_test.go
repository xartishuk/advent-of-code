package main

import "testing"

func TestRope(t *testing.T) {
	tests := []struct {
		inputFile string
		length    int
		visited   int
	}{
		{
			inputFile: "input_test.txt",
			length:    2,
			visited:   13,
		},
		{
			inputFile: "input.txt",
			length:    2,
			visited:   6498,
		},
		{
			inputFile: "input_test_big.txt",
			length:    10,
			visited:   36,
		},
		{
			inputFile: "input.txt",
			length:    10,
			visited:   2531,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		visited := Rope(in, tt.length)

		if visited != tt.visited {
			t.Errorf("expected %d, got %d", tt.visited, visited)
		}
	}
}
