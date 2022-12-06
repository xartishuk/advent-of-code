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

	res := SupplyStacks(CrateMover9000{}, crates, commands)

	fmt.Println(res)
}

func SupplyStacks(mover CrateMover, crates []*stack.Stack, commands []Command) string {
	for _, command := range commands {
		mover.Execute(crates, command)
	}

	// collect top crates
	var top strings.Builder

	for _, c := range crates {
		top.WriteRune(c.Peek().(rune))
	}

	return top.String()
}

type CrateMover interface {
	Execute([]*stack.Stack, Command)
}

type CrateMover9000 struct{}

func (_ CrateMover9000) Execute(crates []*stack.Stack, command Command) {
	for i := 0; i < command.Amount; i++ {
		crates[command.To-1].Push(crates[command.From-1].Pop())
	}
}

type CrateMover9001 struct {
	hand stack.Stack
}

func (m CrateMover9001) Execute(crates []*stack.Stack, command Command) {
	// pick up
	for i := 0; i < command.Amount; i++ {
		m.hand.Push(crates[command.From-1].Pop())
	}
	// put down
	for i := 0; i < command.Amount; i++ {
		crates[command.To-1].Push(m.hand.Pop())
	}
}

type Command struct {
	Amount int
	From   int
	To     int
}
