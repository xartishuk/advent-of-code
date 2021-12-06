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

	first, last := Bingo(input.numbersDrawn, input.boards)

	fmt.Println(first, last)
}

const boardSize = 5

type Board struct {
	IsCompleted bool

	cells [boardSize][boardSize]*BingoCell
}

type BingoCell struct {
	number    int
	isCrossed bool
}

func Bingo(numbersDrawn []int, boards []Board) (first, last int) {
	first, last = -1, -1

	for _, numberDrawn := range numbersDrawn {
		for i := range boards {
			if boards[i].IsCompleted {
				continue
			}

			boards[i].Cross(numberDrawn)
			if boards[i].IsCompleted {
				// save score
				score := boards[i].SumOfUncrossed() * numberDrawn
				if first == -1 {
					first = score
				}
				last = score
			}
		}
	}

	return first, last
}

// Cross crosses number on board updates IsComplete
func (b *Board) Cross(num int) {
	b.IsCompleted = b.cross(num)
}

func (b *Board) cross(num int) bool {
	for i := range b.cells {
		for j := range b.cells[i] {
			if num == b.cells[i][j].number {
				b.cells[i][j].isCrossed = true

				return b.isCrossedOn(i, j)
			}
		}
	}

	return false
}

func (b *Board) isCrossedOn(row, column int) bool {
	return b.isRowCrossed(row) || b.isColumnCrossed(column)
}

func (b *Board) isRowCrossed(row int) bool {
	for j := 0; j < boardSize; j++ {
		if !b.cells[row][j].isCrossed {
			return false
		}
	}

	return true
}

func (b *Board) isColumnCrossed(column int) bool {
	for i := 0; i < boardSize; i++ {
		if !b.cells[i][column].isCrossed {
			return false
		}
	}

	return true
}

func (b *Board) SumOfUncrossed() int {
	var sum int
	for i := range b.cells {
		for j := range b.cells[i] {
			if !b.cells[i][j].isCrossed {
				sum += b.cells[i][j].number
			}
		}
	}

	return sum
}
