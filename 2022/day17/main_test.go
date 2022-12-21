package main

import "testing"

func TestRockTower(t *testing.T) {
	tests := []struct {
		inputFile string
		result    int
	}{
		{
			inputFile: "input_test.txt",
			result:    3068,
		},
		{
			inputFile: "input.txt",
			result:    3130,
		},
	}

	for _, tt := range tests {
		in, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		result := RockTower(in)

		if result != tt.result {
			t.Errorf("expected %d, got %d", tt.result, result)
		}
	}
}
