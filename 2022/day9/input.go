package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var moves strings.Builder

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := strings.Split(s.Text(), " ")

		moves.WriteString(strings.Repeat(line[0], mustAtoi(line[1])))
	}

	return moves.String(), s.Err()
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
