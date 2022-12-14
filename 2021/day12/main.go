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

	result := CavePaths(start, end)

	fmt.Println(result)
}

func CavePaths(start, end *Cave) int {
	return start.PathsTo(map[string]struct{}{}, end)
}

type Cave struct {
	name    string
	isSmall bool

	neighbours []*Cave
}

func (c *Cave) PathsTo(smallVisited map[string]struct{}, to *Cave) int {
	if c == to {
		return 1
	}

	if c.isSmall {
		smallVisited[c.name] = struct{}{}
		defer delete(smallVisited, c.name)
	}

	var sum int
	for _, neighbour := range c.neighbours {
		if _, ok := smallVisited[neighbour.name]; ok {
			continue
		}
		sum += neighbour.PathsTo(smallVisited, to)
	}

	return sum
}
