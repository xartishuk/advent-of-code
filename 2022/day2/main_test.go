package main

import "testing"

func TestRPS_Part1(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  15,
		},
		{
			inputFile: "input.txt",
			expected:  11841,
		},
	}

	for _, tt := range tests {
		p1, p2, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		add2(p1)

		result := RPS(p1, p2)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}

