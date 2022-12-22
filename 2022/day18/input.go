package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]Drop, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var drops []Drop

	for s.Scan() {
		coords := strings.Split(s.Text(), ",")

		drops = append(drops, Drop{
			x:     mustAtoi(coords[0]),
			y:     mustAtoi(coords[1]),
			z:     mustAtoi(coords[2]),
			Sides: 6,
		})
	}

	return drops, s.Err()
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
