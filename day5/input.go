package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]Vector, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]Vector, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		pointsString := strings.FieldsFunc(s.Text(), func(c rune) bool {
			switch c {
			case ',', '-', '>', ' ':
				return true
			}
			return false
		})

		var vector Vector
		vector.From.X, _ = strconv.Atoi(pointsString[0])
		vector.From.Y, _ = strconv.Atoi(pointsString[1])
		vector.To.X, _ = strconv.Atoi(pointsString[2])
		vector.To.Y, _ = strconv.Atoi(pointsString[3])
		vector.Direction = direction(vector.From, vector.To)

		res = append(res, vector)
	}

	return res, s.Err()
}

func direction(p1, p2 Point) Point {
	var dir Point
	if p1.X > p2.X {
		dir.X = -1
	}
	if p1.X < p2.X {
		dir.X = 1
	}
	if p1.Y > p2.Y {
		dir.Y = -1
	}
	if p1.Y < p2.Y {
		dir.Y = 1
	}

	return dir
}
