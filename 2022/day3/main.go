package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Rucksack(lines)

	fmt.Println(result)
}

func Rucksack(rucks []string) int {
	var total int

	for _, ruck := range rucks {
		ruck1, ruck2 := ruck[:len(ruck)/2], ruck[len(ruck)/2:]

		catalog := catalogItems(ruck1)

		t := findRepeatedType(catalog, ruck2)

		total += priority(t)
	}

	return total
}

func catalogItems(ruck string) map[rune]struct{} {
	catalog := make(map[rune]struct{}, len(ruck))
	for _, v := range ruck {
		catalog[v] = struct{}{}
	}
	return catalog
}

func findRepeatedType(catalog map[rune]struct{}, ruck string) rune {
	for _, v := range ruck {
		if _, ok := catalog[v]; ok {
			return v
		}
	}

	panic("couldn't find repeated type")
}

const (
	lowerOffset = 96
	upperOffset = 64
)

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - lowerOffset
	} else if r >= 'A' && r <= 'Z' {
		return int(r) - upperOffset + 26
	}

	panic(fmt.Sprintf("unexpected type in rucksack: %c", r))
}

func readInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var res []string

	s := bufio.NewScanner(f)

	for s.Scan() {
		res = append(res, s.Text())
	}

	return res, s.Err()
}
