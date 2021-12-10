package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([][]int, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		numsString := strings.Split(s.Text(), "")

		lineRes := make([]int, 0, len(numsString))

		for _, numString := range numsString {
			num, _ := strconv.Atoi(numString)

			lineRes = append(lineRes, num)
		}

		res = append(res, lineRes)
	}

	return res, s.Err()
}
