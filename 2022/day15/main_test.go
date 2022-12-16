package main

import "testing"

func TestBeaconExclusionZone(t *testing.T) {
	tests := []struct {
		inputFile string
		y         int
		result    int
	}{
		{
			inputFile: "input_test.txt",
			y:         10,
			result:    26,
		},
		{
			inputFile: "input.txt",
			y:         200000,
			result:    1,
			// 4'959'739 - too low
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := BeaconExclusionZone(in, tt.y)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
