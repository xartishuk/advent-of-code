package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type input struct {
	numbersDrawn []int
	boards       []Board
}

func readInput(filename string) (input, error) {
	f, err := os.Open(filename)
	if err != nil {
		return input{}, err
	}
	defer f.Close()

	var res input

	s := bufio.NewScanner(f)

	// scan first line with numbers drawn
	s.Scan()
	numStrings := strings.Split(s.Text(), ",")

	res.numbersDrawn = make([]int, 0, len(numStrings))
	for _, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		res.numbersDrawn = append(res.numbersDrawn, num)
	}

	// scan boards
	for {
		// scan empty line separator
		if !s.Scan() {
			break
		}

		board := initBoard()
		// scan board
		for i := 0; i < boardSize; i++ {
			s.Scan()
			row := strings.Fields(s.Text())

			for j, numString := range row {
				board.cells[i][j].number, _ = strconv.Atoi(numString)
			}
		}

		res.boards = append(res.boards, board)
	}

	return res, s.Err()
}

func initBoard() Board {
	var b Board
	for i := range b.cells {
		for j := range b.cells[i] {
			b.cells[i][j] = new(BingoCell)
		}
	}

	return b
}
