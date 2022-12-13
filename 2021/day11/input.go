package main

import (
	"bufio"
	"os"
)

func readInput(filename string) ([size][size]*Octopus, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [size][size]*Octopus{}, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var octopuses [size][size]*Octopus

	for i := 0; s.Scan(); i++ {
		for j, v := range s.Text() {
			octopuses[i][j] = &Octopus{energy: mustAtoi(v)}
		}
	}

	return octopuses, nil
}

func mustAtoi(r rune) int {
	return int(r) - 48
}
