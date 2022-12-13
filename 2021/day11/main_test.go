package main

import "testing"

func TestOctopusFlashes(t *testing.T) {
	tests := []struct {
		inputFile string
		flashes   int
		firstSync int
	}{
		{
			inputFile: "input_test.txt",
			flashes:   1656,
			firstSync: 195,
		},
		{
			inputFile: "input.txt",
			flashes:   1785,
			firstSync: 354,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		flashes, firstSync := OctopusFlashes(in)

		if flashes != tt.flashes {
			t.Errorf("expected %d, got %d", tt.flashes, flashes)
		}
		if firstSync != tt.firstSync {
			t.Errorf("expected %d, got %d", tt.firstSync, firstSync)
		}
	}
}
