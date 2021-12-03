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
	return 0
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
