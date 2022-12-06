package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"log"
)

func main() {
	crates, commands, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := SupplyStacks(crates, commands)

	fmt.Println(res)
}

func SupplyStacks(crates []*stack.Stack, commands []Command) string {
	return "CMZ"
}

type Command struct {
	amount int
	from   int
	to     int
}
