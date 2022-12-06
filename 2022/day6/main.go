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

	res := TuningMarker(in)

	fmt.Println(res)
}

func TuningMarker(s string) int {
	for i := 3; i < len(s); i++ {
		if allUnique(s[i-3], s[i-2], s[i-1], s[i]) {
			return i + 1
		}
	}

	panic("couldn't find marker")
}

func allUnique(r1, r2, r3, r4 uint8) bool {
	m := map[uint8]struct{}{
		r1: {},
		r2: {},
		r3: {},
		r4: {},
	}

	return len(m) == 4
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
