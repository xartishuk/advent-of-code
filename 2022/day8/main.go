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

	visible, scenic := TreeHouse(trees)

	fmt.Println(visible)
	fmt.Println(scenic)
}

func TreeHouse(trees [][]*Tree) (int, int) {
	calculateVisibility(trees)
	calculateScenic(trees)

	printVisibility(trees)
	printScenic(trees)

	return countVisible(trees), bestScenic(trees)
}

func calculateVisibility(trees [][]*Tree) {
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

func calculateScenic(trees [][]*Tree) {
	gridSize := len(trees)

	for i := 1; i < gridSize-1; i++ {
		for j := 1; j < gridSize-1; j++ {
			calculateScenicForTree(trees, i, j)
		}
	}
}

func calculateScenicForTree(trees [][]*Tree, ti, tj int) {
	gridSize := len(trees)

	// moving right
	for j := tj + 1; ; j++ {
		if j >= gridSize-1 || trees[ti][tj].Height <= trees[ti][j].Height {
			trees[ti][tj].Scenic *= j - tj
			break
		}
	}
	// moving left
	for j := tj - 1; ; j-- {
		if j <= 0 || trees[ti][tj].Height <= trees[ti][j].Height {
			trees[ti][tj].Scenic *= tj - j
			break
		}
	}
	// moving down
	for i := ti + 1; ; i++ {
		if i >= gridSize-1 || trees[ti][tj].Height <= trees[i][tj].Height {
			trees[ti][tj].Scenic *= i - ti
			break
		}
	}
	// moving up
	for i := ti - 1; ; i-- {
		if i <= 0 || trees[ti][tj].Height <= trees[i][tj].Height {
			trees[ti][tj].Scenic *= ti - i
			break
		}
	}

}

func countVisible(trees [][]*Tree) int {
	gridSize := len(trees)

	// count outside visible first
	count := (gridSize - 1) * 4

	for i := 1; i < gridSize-1; i++ {
		for j := 1; j < gridSize-1; j++ {
			if trees[i][j].IsVisible {
				count++
			}
		}
	}

	return count
}

func bestScenic(trees [][]*Tree) int {
	var best int

	gridSize := len(trees)

	for i := 1; i < gridSize-1; i++ {
		for j := 1; j < gridSize-1; j++ {
			if trees[i][j].Scenic > best {
				best = trees[i][j].Scenic
			}
		}
	}

	return best
}

type Tree struct {
	Height    int
	IsVisible bool
	Scenic    int
}

func printVisibility(trees [][]*Tree) {
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

func printScenic(trees [][]*Tree) {
	gridSize := len(trees)

	var p strings.Builder

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			p.WriteString(fmt.Sprintf("%d\t", trees[i][j].Scenic))
		}
		p.WriteRune('\n')
	}

	fmt.Print(p.String())
}
