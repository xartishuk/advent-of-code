package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]byte, 0)

	s := bufio.NewScanner(f)

	for s.Scan() {
		numsString := strings.Split(s.Text(), ",")

		for _, numString := range numsString {
			num, _ := strconv.Atoi(numString)

			res = append(res, byte(num))
		}
	}

	return res, s.Err()
}
