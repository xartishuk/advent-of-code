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

	result := Sand(in, true)

	fmt.Println(result)
}

func Sand(fillMap [][]bool, hasFloor bool) int {
	if hasFloor {
		// extend fillMap to the right by height
		for range fillMap[0] {
			fillMap = append(fillMap, make([]bool, len(fillMap[0])))
		}

		// create floor on fillMap
		for x := range fillMap {
			fillMap[x] = append(fillMap[x], false, true)
		}
	}

	var i int

	for ; dropGrain(fillMap, hasFloor); i++ {
	}

	return i
}

func dropGrain(fillMap [][]bool, hasFloor bool) bool {
	floorY := len(fillMap[0]) - 1
	x, y := 500, 0

	for {
		if hasFloor {
			if fillMap[500][0] {
				return false
			}
		} else {
			if y == floorY {
				return false
			}
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
