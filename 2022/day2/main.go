package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	XOffset = 87
	AOffset = 64
	AXDiff  = 23
)

func main() {
	p1, p2, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := RPS(p1, p2)

	fmt.Println(result)
}

func convertXYZAsOutcomes(p1, p2 []rune) {
	for i := range p2 {
		switch p2[i] {
		case 'X':
			// TODO: +1
		case 'Y':
			p2[i] = p1[i]
		case 'Z':
			// TODO: -1
		}
	}
}

func RPS(p1 []rune, p2 []rune) int {
	var score int

	for i := range p1 {
		score += selectScore(p2[i]) + roundScore(p1[i], p2[i])
	}

	return score
}

func selectScore(r rune) int {
	score := int(r) % 3
	if score == 0 {
		score = 3
	}

	fmt.Printf("score for %c is %d\n", r, score)

	return score
}

func roundScore(p1, p2 rune) int {
	switch (p1 - p2 + 2) % 3 {
	case 0:
		fmt.Printf("%c %c: %c draws %c\n", p1, p2, p2, p1)
		return 3
	case -1, 2:
		fmt.Printf("%c %c: %c wins over %c\n", p1, p2, p2, p1)
		return 6
	case 1, -2:
		fmt.Printf("%c %c: %c loses to %c\n", p1, p2, p2, p1)
		return 0
	}

	panic("unreachable")
}

func readInput(filename string) ([]rune, []rune, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	var p1, p2 []rune

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()

		p1 = append(p1, rune(line[0]))
		p2 = append(p2, rune(line[2]))
	}

	return p1, p2, s.Err()
}
