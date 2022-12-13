package main

import (
	"fmt"
	"log"
	"strings"
)

const size = 10

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	flashes, firstSync := OctopusFlashes(in)

	fmt.Println(flashes)
	fmt.Println(firstSync)
}

func OctopusFlashes(grid [size][size]*Octopus) (int, int) {
	var flashes int

	fmt.Printf("Before any steps:\n")
	printGrid(grid)

	for i := 0; ; i++ {
		roundFlashes := round(grid)

		fmt.Printf("After step %d:\n", i+1)
		printGrid(grid)

		if i < 100 {
			flashes += roundFlashes
		}

		// all flashed
		if roundFlashes == size*size {
			return flashes, i + 1
		}
	}
}

func round(grid [size][size]*Octopus) int {
	// step 1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j].energy++
		}
	}

	// step 2
	for anyFlashed := false; ; anyFlashed = false {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				cur := grid[i][j]

				if cur.flashed {
					continue
				}

				if cur.energy > 9 {
					anyFlashed = true
					cur.flashed = true
					bumpNeighbours(i, j, grid)
				}
			}
		}

		if !anyFlashed {
			break
		}
	}

	// step 3
	var flashCount int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cur := grid[i][j]

			if cur.flashed {
				cur.flashed = false
				cur.energy = 0

				flashCount++
			}
		}
	}

	return flashCount
}

func bumpNeighbours(i, j int, grid [size][size]*Octopus) {
	if i-1 >= 0 {
		grid[i-1][j].energy++
		if j-1 >= 0 {
			grid[i-1][j-1].energy++
		}
		if j+1 < size {
			grid[i-1][j+1].energy++
		}
	}
	if i+1 < size {
		grid[i+1][j].energy++
		if j-1 >= 0 {
			grid[i+1][j-1].energy++
		}
		if j+1 < size {
			grid[i+1][j+1].energy++
		}
	}
	if j-1 >= 0 {
		grid[i][j-1].energy++
	}
	if j+1 < size {
		grid[i][j+1].energy++
	}
}

type Octopus struct {
	energy int

	flashed bool
}

func printGrid(grid [size][size]*Octopus) {
	var p strings.Builder

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			p.WriteRune(itoa(grid[i][j].energy))
		}
		p.WriteRune('\n')
	}

	fmt.Println(p.String())
}

func itoa(i int) rune {
	return rune(i) + 48
}
