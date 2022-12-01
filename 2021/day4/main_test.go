package main

import "testing"

func TestBingo(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  struct {
			first int
			last  int
		}
	}{
		{
			inputFile: "input_test.txt",
			expected: struct {
				first int
				last  int
			}{first: 4512, last: 1924},
		},
		{
			inputFile: "input.txt",
			expected: struct {
				first int
				last  int
			}{first: 58838, last: 6256},
		},
	}

	for _, tt := range tests {
		input, err := readInput(tt.inputFile)
		if err != nil {
			t.Error(err)
		}

		first, last := Bingo(input.numbersDrawn, input.boards)

		if first != tt.expected.first {
			t.Errorf("expected first to be %d, got %d", tt.expected.first, first)
		}
		if last != tt.expected.last {
			t.Errorf("expected last to be %d, got %d", tt.expected.last, last)
		}
	}
}
