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

	result := Bingo(input.numbersDrawn, input.boards)

	fmt.Println(result)
}

const boardSize = 5

type Board [boardSize][boardSize]*BingoCell

type BingoCell struct {
	number    int
	isCrossed bool
}

func Bingo(numbersDrawn []int, boards []Board) int {
	for _, numberDrawn := range numbersDrawn {
		for _, board := range boards {
			if board.Cross(numberDrawn) {
				return board.SumOfUncrossed() * numberDrawn
			}
		}
	}

	return -1
}

// Cross crosses number on board and returns true if board is complete
func (b *Board) Cross(num int) bool {
	for i := range b {
		for j := range b[i] {
			if num == b[i][j].number {
				b[i][j].isCrossed = true

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
		if !b[row][j].isCrossed {
			return false
		}
	}

	return true
}

func (b *Board) isColumnCrossed(column int) bool {
	for i := 0; i < boardSize; i++ {
		if !b[i][column].isCrossed {
			return false
		}
	}

	return true
}

func (b *Board) SumOfUncrossed() int {
	var sum int
	for i := range b {
		for j := range b[i] {
			if !b[i][j].isCrossed {
				sum += b[i][j].number
			}
		}
	}

	return sum
}
