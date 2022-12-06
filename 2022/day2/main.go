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

	convertXYZAsSigns(p2)

	result := RPS(p1, p2)

	fmt.Println(result)
}

func convertXYZAsSigns(p2 []rune) {
	for i := range p2 {
		p2[i] = p2[i] - AXDiff
	}
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
	return int(r) - AOffset
}

func roundScore(p1, p2 rune) int {
	switch p1 - p2 {
	case 0:
		fmt.Printf("%c %c: %s(%c) draws %s(%c)\n", p1, p2, name(p2), p2, name(p1), p1)
		return 3
	case -1, 2:
		fmt.Printf("%c %c: %s(%c) wins over %s(%c)\n", p1, p2, name(p2), p2, name(p1), p1)
		return 6
	case 1, -2:
		fmt.Printf("%c %c: %s(%c) loses to %s(%c)\n", p1, p2, name(p2), p2, name(p1), p1)
		return 0
	}

	panic("unreachable")
}

func name(r rune) string {
	switch r {
	case 'A':
		return "Rock"
	case 'B':
		return "Paper"
	case 'C':
		return "Scissors"
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
