package main

import "testing"

func TestTreeHouse(t *testing.T) {
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

		visible := Rope(in, tt.length)

		if visible != tt.visited {
			t.Errorf("expected %d, got %d", tt.visited, visible)
		}
	}
}
