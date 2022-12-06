package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in, err := readInput("input_test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := TuningMarker(in, 4)

	fmt.Println(res)
}

func TuningMarker(s string, length int) int {
	for i := length; i <= len(s); i++ {
		if allUnique(s[i-length : i]) {
			return i
		}
	}

	panic("couldn't find marker")
}

func allUnique(s string) bool {
	m := map[uint8]struct{}{}

	for i := range s {
		m[s[i]] = struct{}{}
	}

	return len(m) == len(s)
}

func readInput(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()

	return s.Text(), s.Err()
}
