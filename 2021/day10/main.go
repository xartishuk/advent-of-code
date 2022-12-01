package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	corrupted, incomplete := SyntaxScoring(input)

	fmt.Println(corrupted)
	fmt.Println(incomplete)
}

var pointsForCorrupted = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var pointsForIncomplete = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func SyntaxScoring(lines []string) (corrupted, middleIncomplete int) {
	incompleteScores := make([]int, 0, len(lines))

	for _, l := range lines {
		line := NewLine(l)

		if line.CorruptedOn() != 0 {
			corrupted += pointsForCorrupted[line.CorruptedOn()]
		} else {
			incompleteScores = append(incompleteScores, incompleteScore(line))
		}
	}

	return corrupted, middle(incompleteScores)
}

func NewLine(l string) Line {
	stack, corruptedOn := formStack(l)

	return Line{
		l:           l,
		stack:       stack,
		corruptedOn: corruptedOn,
	}
}

type Line struct {
	l string

	stack *bracketStack

	corruptedOn rune
}

func formStack(l string) (*bracketStack, rune) {
	var stack bracketStack

	for _, bracket := range l {
		switch bracket {
		case '(', '[', '{', '<':
			stack.Push(bracket)
		case ')', ']', '}', '>':
			top := stack.Pop()

			if !match(bracket, top) {
				return nil, bracket
			}
		default:
			panic(bracket)
		}
	}

	return &stack, 0
}

func (l *Line) CorruptedOn() rune {
	return l.corruptedOn
}

func (l *Line) MissingPart() string {
	closers := strings.Builder{}

	for {
		top := l.stack.Pop()
		if top == 0 {
			break
		}

		closers.WriteRune(opposite(top))
	}

	return closers.String()
}

type bracketStack struct {
	arr []rune
}

func (s *bracketStack) Push(bracket rune) {
	s.arr = append(s.arr, bracket)
}

func (s *bracketStack) Pop() rune {
	if len(s.arr) == 0 {
		return 0
	}

	top := s.arr[len(s.arr)-1]

	s.arr = s.arr[:len(s.arr)-1]

	return top
}

func match(close, open rune) bool {
	return opposite(close) == open
}

func opposite(r rune) rune {
	switch r {
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	case '{':
		return '}'
	case '}':
		return '{'
	case '<':
		return '>'
	case '>':
		return '<'
	default:
		panic(r)
	}
}

func incompleteScore(l Line) (score int) {
	for _, bracket := range l.MissingPart() {
		score *= 5
		score += pointsForIncomplete[bracket]
	}

	return score
}

func middle(values []int) int {
	sort.Ints(values)

	return values[len(values)/2]
}