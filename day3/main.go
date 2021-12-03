package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := PowerConsumption(input)

	fmt.Println(result)
}

func PowerConsumption(input []string) int {
	bitCount := countBits(input)

	var gamma, epsilon int

	for _, numOfOnes := range bitCount {
		gamma <<= 1
		epsilon <<= 1
		if numOfOnes > len(input)/2 {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	return gamma * epsilon
}

func countBits(input []string) []int {
	counter := make([]int, len(input[0]))

	for _, line := range input {
		for i, bit := range line {
			if bit == '1' {
				counter[i]++
			}
		}
	}

	return counter
}

func readInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]string, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		res = append(res, s.Text())
	}

	return res, s.Err()
}
