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

	result := SonarSweep(input)

	fmt.Println(result)
}

func SonarSweep([]int) int {
	return 0
}

func readInput(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]int, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		res = append(res, num)
	}

	return res, s.Err()
}
