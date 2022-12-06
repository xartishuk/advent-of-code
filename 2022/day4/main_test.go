package main

import "testing"

func TestCampCleanupFull(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  2,
		},
		{
			inputFile: "input.txt",
			expected:  305,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := CampCleanupFull(in)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
