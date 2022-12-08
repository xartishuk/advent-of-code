package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	trees, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	visible := TreeVisibility(trees)

	fmt.Println(visible)
}

func TreeVisibility(trees [][]*Tree) int {
	checkVisibility(trees)

	printGrid(trees)

	return countInsideVisible(trees) + countOutsideVisible(len(trees))
}

func checkVisibility(trees [][]*Tree) {
	gridSize := len(trees)

	for i := 1; i < gridSize-1; i++ {
		// moving right
		highest := trees[i][0].Height
		for j := 1; j < gridSize-1; j++ {
			if trees[i][j].Height > highest {
				highest = trees[i][j].Height
				trees[i][j].IsVisible = true
			}
		}
		// moving left
		highest = trees[i][gridSize-1].Height
		for j := gridSize - 2; j > 0; j-- {
			if trees[i][j].Height > highest {
				highest = trees[i][j].Height
				trees[i][j].IsVisible = true
			}
		}
	}
	for j := 1; j < gridSize-1; j++ {
		// moving down
		highest := trees[0][j].Height
		for i := 1; i < gridSize-1; i++ {
			if trees[i][j].Height > highest {
				highest = trees[i][j].Height
				trees[i][j].IsVisible = true
			}
		}
		// moving up
		highest = trees[gridSize-1][j].Height
		for i := gridSize - 2; i > 0; i-- {
			if trees[i][j].Height > highest {
				highest = trees[i][j].Height
				trees[i][j].IsVisible = true
			}
		}
	}
}

func countInsideVisible(trees [][]*Tree) int {
	var count int

	gridSize := len(trees)

	for i := 1; i < gridSize-1; i++ {
		for j := 1; j < gridSize-1; j++ {
			if trees[i][j].IsVisible {
				count++
			}
		}
	}

	return count
}

func countOutsideVisible(size int) int {
	return (size - 1) * 4
}

type Tree struct {
	Height    int
	IsVisible bool
}

func printGrid(trees [][]*Tree) {
	gridSize := len(trees)

	var p strings.Builder

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if trees[i][j].IsVisible {
				p.WriteRune('T')
			} else {
				p.WriteRune('.')
			}
		}
		p.WriteRune('\n')
	}

	fmt.Print(p.String())
}
