package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var pointLines [][]point

	for s.Scan() {
		inputPoints := strings.Split(s.Text(), " -> ")

		pointLine := make([]point, 0, len(inputPoints))

		for _, p := range inputPoints {
			pointLine = append(pointLine, parsePoint(p))
		}

		pointLines = append(pointLines, pointLine)
	}

	maxX, maxY := findMaxes(pointLines)

	fillMap := make([][]bool, maxX+1)
	for x := range fillMap {
		fillMap[x] = make([]bool, maxY+1)
	}

	for _, rock := range pointLines {
		for i := 1; i < len(rock); i++ {
			from, to := rock[i-1], rock[i]

			dir := dir(from, to)

			for cur := from; !equal(cur, to); cur = plus(cur, dir) {
				fillMap[cur.x][cur.y] = true
			}
		}

		last := rock[len(rock)-1]

		fillMap[last.x][last.y] = true
	}

	return fillMap, nil
}

func parsePoint(s string) point {
	parts := strings.Split(s, ",")

	return point{
		x: mustAtoi(parts[0]),
		y: mustAtoi(parts[1]),
	}
}

type point struct {
	x, y int
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func findMaxes(pointLines [][]point) (maxX, maxY int) {
	for _, line := range pointLines {
		for _, p := range line {
			if p.x > maxX {
				maxX = p.x
			}
			if p.y > maxY {
				maxY = p.y
			}
		}
	}

	return maxX, maxY
}

func dir(from, to point) point {
	return point{
		x: floor1(to.x - from.x),
		y: floor1(to.y - from.y),
	}
}

func equal(p1, p2 point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func plus(p1, p2 point) point {
	return point{
		x: p1.x + p2.x,
		y: p1.y + p2.y,
	}
}

func floor1(v int) int {
	if v < -1 {
		return -1
	}
	if v > 1 {
		return 1
	}

	return v
}
