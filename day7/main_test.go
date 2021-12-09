package main

import (
	"testing"
)

func TestWhaleChase(t *testing.T) {
	tests := []struct {
		inputFile string
		complex bool
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			complex: false,
			expected:  37,
		},
		{
			inputFile: "input.txt",
			complex: false,
			expected:  356992,
		},
		{
			inputFile: "input_test.txt",
			complex: true,
			expected:  168,
		},
		{
			inputFile: "input.txt",
			complex: true,
			expected:  101268110,
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := WhaleChase(input, tt.complex)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
