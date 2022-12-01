package main

import (
	"testing"
)

func TestSyntaxScoring(t *testing.T) {
	tests := []struct {
		inputFile  string
		corrupted  int
		incomplete int
	}{
		{
			inputFile:  "input_test.txt",
			corrupted:  26397,
			incomplete: 288957,
		},
		{
			inputFile:  "input.txt",
			corrupted:  319233,
			incomplete: 1118976874,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		corrupted, incomplete := SyntaxScoring(input)

		if corrupted != tt.corrupted {
			t.Errorf("expected %d, got %d", tt.corrupted, corrupted)
		}
		if incomplete != tt.incomplete {
			t.Errorf("expected %d, got %d", tt.incomplete, incomplete)
		}
	}
}
