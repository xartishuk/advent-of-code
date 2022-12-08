package main

import "testing"

func TestTreeHouse(t *testing.T) {
	tests := []struct {
		inputFile string
		visible   int
		scenic    int
	}{
		{
			inputFile: "input_test.txt",
			visible:   21,
			scenic:    8,
		},
		{
			inputFile: "input.txt",
			visible:   1835,
			scenic:    263670,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		visible, scenic := TreeHouse(in)

		if visible != tt.visible {
			t.Errorf("expected %d, got %d", tt.visible, visible)
		}
		if scenic != tt.scenic {
			t.Errorf("expected %d, got %d", tt.scenic, scenic)
		}
	}
}
