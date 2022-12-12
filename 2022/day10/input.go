package main

import (
	"io"
	"os"
	"strings"
)

func readInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	in, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.ReplaceAll(string(in), "\n", " "), " "), nil
}
