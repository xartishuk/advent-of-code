package main

import (
	"bufio"
	"os"
)

func readInput(filename string) ([][]*Tree, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var trees [][]*Tree

	s := bufio.NewScanner(f)

	for s.Scan() {
		var treeRow []*Tree

		for _, t := range s.Text() {
			treeRow = append(treeRow, newTree(t))
		}

		trees = append(trees, treeRow)
	}

	return trees, s.Err()
}

func digit(r rune) int {
	return int(r) - 48
}

func newTree(t rune) *Tree {
	return &Tree{
		Height: digit(t),
		Scenic: 1,
	}
}
