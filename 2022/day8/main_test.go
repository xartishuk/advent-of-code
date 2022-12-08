package main

import "testing"

func TestDeviceSpace(t *testing.T) {
	tests := []struct {
		inputFile string
		visible   int
	}{
		{
			inputFile: "input_test.txt",
			visible:   21,
		},
		{
			inputFile: "input.txt",
			visible:   1835,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		visible := TreeVisibility(in)

		if visible != tt.visible {
			t.Errorf("expected %d, got %d", tt.visible, visible)
		}
	}
}
