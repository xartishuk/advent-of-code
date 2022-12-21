package main

import (
	"bufio"
	"os"
)

func readInput(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()

	return s.Text(), s.Err()
}
