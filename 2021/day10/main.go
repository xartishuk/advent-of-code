package main

import (
	"fmt"
	"log"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := SyntaxScoring(input)

	fmt.Println(result)
}

var pointsPerBracket = map[rune]int{
	0:   0,
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func SyntaxScoring(lines []string) int {
	var pointTotal int

	for _, l := range lines {
		pointTotal += pointsPerBracket[Line(l).corruptedOn()]
	}

	return pointTotal
}

type Line string

func (l Line) corruptedOn() rune {
	var stack bracketStack

	for _, bracket := range l {
		switch bracket {
		case '<', '(', '{', '[':
			stack.push(bracket)
		case '>', ')', '}', ']':
			top := stack.pop()

			if !match(bracket, top) {
				return bracket
			}
		default:
			panic(bracket)
		}
	}

	return 0
}

type bracketStack struct {
	arr []rune
}

func (s *bracketStack) push(bracket rune) {
	s.arr = append(s.arr, bracket)
}

func (s *bracketStack) pop() rune {
	top := s.arr[len(s.arr)-1]

	s.arr = s.arr[:len(s.arr)-1]

	return top
}

func match(close, open rune) bool {
	switch close {
	case '>':
		return open == '<'
	case ')':
		return open == '('
	case '}':
		return open == '{'
	case ']':
		return open == '['
	default:
		panic("wrong match usage")
	}
}
