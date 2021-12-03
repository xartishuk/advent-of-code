package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//result := PowerConsumption(input)
	result := LifeSupport(input)

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

func LifeSupport(input []string) int64 {
	oxygenRatingString := oxygen(input, 0)
	co2RatingString := co2(input, 0)

	oxygenRating, _ := strconv.ParseInt(oxygenRatingString, 2, 0)
	co2Rating, _ := strconv.ParseInt(co2RatingString, 2, 0)

	return oxygenRating * co2Rating
}

func oxygen(lines []string, column int) string {
	if len(lines) == 1 {
		return lines[0]
	}

	s := columnSum(lines, column)
	if s >= len(lines)/2+len(lines)%2 {
		lines = filter(lines, column, '1')
	} else {
		lines = filter(lines, column, '0')
	}

	return oxygen(lines, column+1)
}

func co2(lines []string, column int) string {
	if len(lines) == 1 {
		return lines[0]
	}

	s := columnSum(lines, column)
	if s >= len(lines)/2+len(lines)%2 {
		lines = filter(lines, column, '0')
	} else {
		lines = filter(lines, column, '1')
	}

	return co2(lines, column+1)
}

func columnSum(lines []string, column int) int {
	var sum int

	for _, line := range lines {
		if line[column] == '1' {
			sum++
		}
	}

	return sum
}

func filter(lines []string, column int, bit byte) []string {
	res := make([]string, 0, len(lines))

	for _, line := range lines {
		if line[column] == bit {
			res = append(res, line)
		}
	}

	return res
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
