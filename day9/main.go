package main

import (
	"fmt"
	"log"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := SmokeBasin(input)

	fmt.Println(result)
}

func SmokeBasin(floor [][]int) int {
	var danger int
	imax, jmax := len(floor)-1, len(floor[0])-1

	// corners
	if floor[0][0] < floor[0][1] &&
		floor[0][0] < floor[1][0] {
		danger += floor[0][0] + 1
	}
	if floor[0][jmax] < floor[0][jmax-1] &&
		floor[0][jmax] < floor[1][jmax] {
		danger += floor[0][jmax] + 1
	}
	if floor[imax][0] < floor[imax][1] &&
		floor[imax][0] < floor[imax-1][0] {
		danger += floor[imax][0] + 1
	}
	if floor[imax][jmax] < floor[imax-1][jmax] &&
		floor[imax][jmax] < floor[imax][jmax-1] {
		danger += floor[imax][jmax] + 1
	}

	// edges
	for j := 1; j < jmax; j++ {
		if floor[0][j] < floor[0][j-1] &&
			floor[0][j] < floor[0][j+1] &&
			floor[0][j] < floor[1][j] {
			danger += floor[0][j] + 1
		}
	}
	for j := 1; j < jmax; j++ {
		if floor[imax][j] < floor[imax][j-1] &&
			floor[imax][j] < floor[imax][j+1] &&
			floor[imax][j] < floor[imax-1][j] {
			danger += floor[imax][j] + 1
		}
	}
	for i := 1; i < imax; i++ {
		if floor[i][0] < floor[i-1][0] &&
			floor[i][0] < floor[i+1][0] &&
			floor[i][0] < floor[i][1] {
			danger += floor[i][0] + 1
		}
	}
	for i := 1; i < imax; i++ {
		if floor[i][jmax] < floor[i-1][jmax] &&
			floor[i][jmax] < floor[i+1][jmax] &&
			floor[i][jmax] < floor[i][jmax-1] {
			danger += floor[i][jmax] + 1
		}
	}

	// middle
	for i := 1; i < imax; i++ {
		for j := 1; j < jmax; j++ {
			if floor[i][j] < floor[i-1][j] &&
				floor[i][j] < floor[i+1][j] &&
				floor[i][j] < floor[i][j-1] &&
				floor[i][j] < floor[i][j+1] {
				danger += floor[i][j] + 1
			}
		}
	}

	return danger
}
