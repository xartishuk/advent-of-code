package main

import "testing"

func TestCRTCPU(t *testing.T) {
	tests := []struct {
		inputFile string
		sum       int
	}{
		{
			inputFile: "input_test.txt",
			sum:       13140,
		},
		{
			inputFile: "input.txt",
			sum:       14520,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		sum := CRTCPU(in)

		if sum != tt.sum {
			t.Errorf("expected %d, got %d", tt.sum, sum)
		}
	}
}
