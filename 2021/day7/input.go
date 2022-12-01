package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]int, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		numsString := strings.Split(s.Text(), ",")

		for _, numString := range numsString {
			num, _ := strconv.Atoi(numString)

			res = append(res, num)
		}
	}

	return res, s.Err()
}
