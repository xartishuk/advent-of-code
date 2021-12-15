package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	linkedFloor := LinkedFloor(input)

	danger := SmokeDanger(linkedFloor)
	basins := SmokeBasins(linkedFloor)

	fmt.Println(danger, basins)
}

func SmokeDanger(floor [][]Tile) int {
	var danger int

	for i := range floor {
		for j := range floor[i] {
			if floor[i][j].IsLocalMinimum() {
				danger += floor[i][j].Height + 1
			}
		}
	}

	return danger
}

func SmokeBasins(floor [][]Tile) int {
	topBasins := make([]int, 4)

	for i := range floor {
		for j := range floor[i] {
			if floor[i][j].IsLocalMinimum() {
				topBasins[0] = floor[i][j].BasinCount()
				sort.Ints(topBasins)
			}
		}
	}

	return topBasins[1] * topBasins[2] * topBasins[3]
}

func LinkedFloor(floor [][]int) [][]Tile {
	imax, jmax := len(floor)-1, len(floor[0])-1

	linked := make([][]Tile, len(floor))

	for i := range floor {
		linked[i] = make([]Tile, len(floor[i]))
		for j := range floor[i] {
			linked[i][j] = Tile{
				Height:         floor[i][j],
				CountedInBasin: floor[i][j] == 9,
			}
		}
	}

	// corners
	linked[0][0].Right = &linked[0][1]
	linked[0][0].Down = &linked[1][0]

	linked[0][jmax].Left = &linked[0][jmax-1]
	linked[0][jmax].Down = &linked[1][jmax]

	linked[imax][0].Right = &linked[imax][1]
	linked[imax][0].Up = &linked[imax-1][0]

	linked[imax][jmax].Left = &linked[imax][jmax-1]
	linked[imax][jmax].Up = &linked[imax-1][jmax]

	// edges
	for j := 1; j < jmax; j++ {
		linked[0][j].Left = &linked[0][j-1]
		linked[0][j].Right = &linked[0][j+1]
		linked[0][j].Down = &linked[1][j]
	}
	for j := 1; j < jmax; j++ {
		linked[imax][j].Left = &linked[imax][j-1]
		linked[imax][j].Right = &linked[imax][j+1]
		linked[imax][j].Up = &linked[imax-1][j]
	}
	for i := 1; i < imax; i++ {
		linked[i][0].Up = &linked[i-1][0]
		linked[i][0].Down = &linked[i+1][0]
		linked[i][0].Right = &linked[i][1]
	}
	for i := 1; i < imax; i++ {
		linked[i][jmax].Up = &linked[i-1][jmax]
		linked[i][jmax].Down = &linked[i+1][jmax]
		linked[i][jmax].Left = &linked[i][jmax-1]
	}

	// middle
	for i := 1; i < imax; i++ {
		for j := 1; j < jmax; j++ {
			linked[i][j].Up = &linked[i-1][j]
			linked[i][j].Down = &linked[i+1][j]
			linked[i][j].Left = &linked[i][j-1]
			linked[i][j].Right = &linked[i][j+1]
		}
	}

	return linked
}

type Tile struct {
	Height int

	CountedInBasin bool

	Up    *Tile
	Down  *Tile
	Right *Tile
	Left  *Tile
}

func (t *Tile) IsLocalMinimum() bool {
	if t.Up != nil && t.Up.Height <= t.Height {
		return false
	}
	if t.Down != nil && t.Down.Height <= t.Height {
		return false
	}
	if t.Left != nil && t.Left.Height <= t.Height {
		return false
	}
	if t.Right != nil && t.Right.Height <= t.Height {
		return false
	}

	return true
}

func (t *Tile) BasinCount() int {
	if t.CountedInBasin {
		return 0
	}

	t.CountedInBasin = true

	count := 1
	if t.Up != nil {
		count += t.Up.BasinCount()
	}
	if t.Down != nil {
		count += t.Down.BasinCount()
	}
	if t.Left != nil {
		count += t.Left.BasinCount()
	}
	if t.Right != nil {
		count += t.Right.BasinCount()
	}

	return count
}

func (t *Tile) String() string {
	return strconv.Itoa(t.Height)
}
