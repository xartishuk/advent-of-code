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

	result := Hydrothermal(input)

	fmt.Println(result)
}

type Vector struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

func Hydrothermal(vectors []Vector) int {
	return 0
}
