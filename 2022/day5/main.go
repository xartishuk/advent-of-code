package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"log"
	"strings"
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
	for _, command := range commands {
		command.ApplyOn(crates)
	}

	// collect top crates
	var top strings.Builder

	for _, c := range crates {
		top.WriteRune(c.Peek().(rune))
	}

	return top.String()
}

func (c Command) ApplyOn(crates []*stack.Stack) {
	for i := 0; i < c.amount; i++ {
		crates[c.to-1].Push(crates[c.from-1].Pop())
	}
}

type Command struct {
	amount int
	from   int
	to     int
}
