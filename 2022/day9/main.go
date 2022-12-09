package main

import (
	"fmt"
	"log"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	visited := Rope(in)

	fmt.Println(visited)
}

func Rope(moves string) int {
	var hi, hj, ti, tj int
	tailVisitedGrid := map[int]map[int]struct{}{0: {0: struct{}{}}}

	for _, move := range moves {
		switch move {
		case 'U':
			hi++
			if absDiff(hi, ti) > 1 || absDiff(hj, tj) > 1 {
				ti, tj = hi-1, hj
			}
		case 'D':
			hi--
			if absDiff(hi, ti) > 1 || absDiff(hj, tj) > 1 {
				ti, tj = hi+1, hj
			}
		case 'R':
			hj++
			if absDiff(hi, ti) > 1 || absDiff(hj, tj) > 1 {
				ti, tj = hi, hj-1
			}
		case 'L':
			hj--
			if absDiff(hi, ti) > 1 || absDiff(hj, tj) > 1 {
				ti, tj = hi, hj+1
			}
		}

		// save tail path
		if tailVisitedRow, ok := tailVisitedGrid[ti]; ok {
			tailVisitedRow[tj] = struct{}{}
		} else {
			tailVisitedGrid[ti] = map[int]struct{}{tj: {}}
		}
	}

	var visited int

	for _, tailVisitedRow := range tailVisitedGrid {
		visited += len(tailVisitedRow)
	}

	return visited
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
