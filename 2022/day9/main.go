package main

import (
	"fmt"
	"log"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	visited := Rope(in, 10)

	fmt.Println(visited)
}

func Rope(moves string, length int) int {
	head, tail := newRope(length)

	tailVisitedGrid := map[int]map[int]struct{}{0: {0: struct{}{}}}

	for _, move := range moves {
		head.ForceMove(move)

		// save tail path
		if tailVisitedRow, ok := tailVisitedGrid[tail.i]; ok {
			tailVisitedRow[tail.j] = struct{}{}
		} else {
			tailVisitedGrid[tail.i] = map[int]struct{}{tail.j: {}}
		}
	}

	var visited int

	for _, tailVisitedRow := range tailVisitedGrid {
		visited += len(tailVisitedRow)
	}

	return visited
}

func newRope(length int) (head, tail *Knot) {
	// can be remade with pointers
	knots := make([]*Knot, length)

	// from tail to head
	for i := range knots {
		knots[i] = &Knot{}
		if i != 0 {
			knots[i].next = knots[i-1]
		}
	}

	return knots[length-1], knots[0]
}

type Knot struct {
	i int
	j int

	next *Knot
}

func (k *Knot) ForceMove(direction rune) {
	switch direction {
	case 'U':
		k.i++
	case 'D':
		k.i--
	case 'R':
		k.j++
	case 'L':
		k.j--
	}

	if k.next != nil {
		k.next.moveAfter(k)
	}
}

func (k *Knot) moveAfter(h *Knot) {
	if k.isAdjacentTo(h) {
		return
	}

	k.i += floorAbs1(h.i - k.i)
	k.j += floorAbs1(h.j - k.j)

	if k.next != nil {
		k.next.moveAfter(k)
	}
}

func (k *Knot) isAdjacentTo(h *Knot) bool {
	return absDiff(k.i, h.i) <= 1 && absDiff(k.j, h.j) <= 1
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

func floorAbs1(v int) int {
	if v > 1 {
		return 1
	}

	if v < -1 {
		return -1
	}

	return v
}
