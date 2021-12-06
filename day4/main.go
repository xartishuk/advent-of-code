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

type Board [boardSize][boardSize]BingoCell

type BingoCell struct {
	number    int
	isCrossed bool
}

func Bingo(numbersDrawn []int, boards []Board) int {
	return 0
}
