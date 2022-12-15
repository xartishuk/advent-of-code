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

	result := Sand(in)

	fmt.Println(result)
}

func Sand(fillMap [][]bool) int {
	var i int

	for ; dropGrain(fillMap); i++ {
	}

	return i
}

func dropGrain(fillMap [][]bool) bool {
	floorY := len(fillMap[0]) - 1
	x, y := 500, 0

	for {
		if y == floorY {
			return false
		}

		// assume we are far from left and right bounds

		// center free
		if !fillMap[x][y+1] {
			y++
			continue
		}
		// left free
		if !fillMap[x-1][y+1] {
			y++
			x--
			continue
		}
		// right free
		if !fillMap[x+1][y+1] {
			y++
			x++
			continue
		}

		// settle
		fillMap[x][y] = true
		return true
	}
}
