package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Position(input)

	fmt.Println(result)
}

type Input struct {
	Direction string
	Amount    int
}

func Position(in <-chan Input) int {
	for v := range in {
		_ = v
	}

	return 0
}

func readInput(filename string) (<-chan Input, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	res := make(chan Input)

	go func(f io.ReadCloser, res chan<- Input) {
		defer f.Close()
		defer close(res)

		s := bufio.NewScanner(f)

		for s.Scan() {
			parts := strings.Split(s.Text(), " ")
			amount, _ := strconv.Atoi(parts[1])
			res <- Input{
				Direction: parts[0],
				Amount:    amount,
			}
		}
	}(f, res)

	return res, nil
}
