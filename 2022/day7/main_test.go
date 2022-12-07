package main

import "testing"

func TestDeviceSpace(t *testing.T) {
	tests := []struct {
		inputFile  string
		sumOver10k int
		deleteSize int
	}{
		{
			inputFile:  "input_test.txt",
			sumOver10k: 95437,
			deleteSize: 24933642,
		},
		{
			inputFile:  "input.txt",
			sumOver10k: 1749646,
			deleteSize: -1,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		sumOver10k, deleteSize := DeviceSpace(in)

		if sumOver10k != tt.sumOver10k {
			t.Errorf("expected %d, got %d", tt.sumOver10k, sumOver10k)
		}
		if deleteSize != tt.deleteSize {
			t.Errorf("expected %d, got %d", tt.deleteSize, deleteSize)
		}
	}
}
