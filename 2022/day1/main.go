package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Top3Calories(input)

	fmt.Println(result)
}

func Top1Calories(deers [][]int) int {
	var maxCalories int

	for _, deer := range deers {
		deerCalories := calories(deer)

		if deerCalories > maxCalories {
			maxCalories = deerCalories
		}
	}

	return maxCalories
}

func Top3Calories(deers [][]int) int {
	var top3Calories [3]int

	for _, deer := range deers {
		deerCalories := calories(deer)

		if deerCalories > top3Calories[0] {
			top3Calories[0], top3Calories[1], top3Calories[2] = deerCalories, top3Calories[0], top3Calories[1]
		} else if deerCalories > top3Calories[1] {
			top3Calories[1], top3Calories[2] = deerCalories, top3Calories[1]
		} else if deerCalories > top3Calories[2] {
			top3Calories[2] = deerCalories
		}
	}

	return top3Calories[0] + top3Calories[1] + top3Calories[2]
}

func calories(foods []int) (calories int) {
	for _, food := range foods {
		calories += food
	}

	return calories
}

func readInput(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([][]int, 0)
	block := make([]int, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		if line == "" {
			// save block and reset
			res = append(res, block)
			block = make([]int, 0)

			continue
		}

		num, _ := strconv.Atoi(line)
		block = append(block, num)
	}

	// save the last block
	res = append(res, block)

	return res, s.Err()
}
