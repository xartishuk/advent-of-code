package main

import "testing"

func TestTreeHouse(t *testing.T) {
	tests := []struct {
		inputFile string
		visited   int
	}{
		{
			inputFile: "input_test.txt",
			visited:   13,
		},
		{
			inputFile: "input.txt",
			visited:   6498,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		visible := Rope(in)

		if visible != tt.visited {
			t.Errorf("expected %d, got %d", tt.visited, visible)
		}
	}
}
