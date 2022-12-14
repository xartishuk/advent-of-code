package main

import (
	"fmt"
	"log"
)

func main() {
	start, end, err := readInput("input_test_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := CavePaths(start, end, true)

	fmt.Println(result)
}

func CavePaths(start, end *Cave, allowRepeatOnce bool) int {
	navigator := &Navigator{
		visited:         map[string]struct{}{},
		allowRepeatOnce: allowRepeatOnce,
	}

	return navigator.PathsTo(start, end)
}

type Cave struct {
	name    string
	isSmall bool

	neighbours []*Cave
}

type Navigator struct {
	visited map[string]struct{}

	allowRepeatOnce bool
	repeatedOnce    *Cave
}

func (n *Navigator) PathsTo(from, to *Cave) int {
	if from == to {
		return 1
	}

	if from.isSmall {
		n.visited[from.name] = struct{}{}
	}

	var sum int
	for _, neighbour := range from.neighbours {
		if _, ok := n.visited[neighbour.name]; ok {
			// already been to that small cave
			if n.allowRepeatOnce && n.repeatedOnce == nil {
				// but allowed to repeat and haven't repeated yet
				// so just mark this cave as repeated
				n.repeatedOnce = neighbour
			} else {
				// not allowed to repeat or already repeated - skipping
				continue
			}
		}

		sum += n.PathsTo(neighbour, to)

		// reset repeated
		if neighbour == n.repeatedOnce {
			n.repeatedOnce = nil
		}
	}

	if from.isSmall {
		if from == n.repeatedOnce {
			n.repeatedOnce = nil
		} else {
			delete(n.visited, from.name)
		}
	}

	return sum
}
