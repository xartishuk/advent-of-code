package main

import "testing"

func TestDeviceSpace(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test.txt",
			expected:  95437,
		},
		{
			inputFile: "input.txt",
			expected:  1749646,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := DeviceSpace(in)

		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}
