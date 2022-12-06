package main

import "testing"

func TestCampCleanup(t *testing.T) {
	tests := []struct {
		inputFile string
		full      int
		partial   int
	}{
		{
			inputFile: "input_test.txt",
			full:      2,
			partial:   4,
		},
		{
			inputFile: "input.txt",
			full:      305,
			partial:   811,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		full, partial := CampCleanup(in)

		if full != tt.full {
			t.Errorf("expected %d, got %d", tt.full, full)
		}
		if partial != tt.partial {
			t.Errorf("expected %d, got %d", tt.partial, partial)
		}
	}
}
