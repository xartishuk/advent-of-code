package main

import "testing"

func TestBeaconExclusionZone(t *testing.T) {
	tests := []struct {
		inputFile string
		y         int
		excluded  int
		beacon    int
	}{
		{
			inputFile: "input_test.txt",
			y:         10,
			excluded:  26,
			beacon:    56000011,
		},
		{
			inputFile: "input.txt",
			y:         2000000,
			excluded:  5256611,
			beacon:    1,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		excluded, beacon := BeaconExclusionZone(in, tt.y)

		if excluded != tt.excluded {
			t.Errorf("expected %d, got %d", tt.excluded, excluded)
		}
		if beacon != tt.beacon {
			t.Errorf("expected %d, got %d", tt.beacon, beacon)
		}
	}
}
