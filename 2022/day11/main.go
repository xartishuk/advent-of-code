package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	visited := MonkeyBusiness(in)

	fmt.Println(visited)
}

func MonkeyBusiness(monkeys []*Monkey) int {
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.Turn()
		}
	}

	inspections := make([]int, len(monkeys))

	for i := range monkeys {
		inspections[i] = monkeys[i].Inspections()
	}

	fmt.Println(inspections)

	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

type Monkey struct {
	items     chan int
	inspector Inspector
	tester    Tester

	monkeys []*Monkey
}

func (m *Monkey) Turn() {
	for {
		select {
		case item := <-m.items:
			item = m.inspector.Inspect(item)

			throwTo := m.tester.Test(item)

			m.monkeys[throwTo].items <- item
		default:
			return
		}
	}
}

func (m *Monkey) Inspections() int {
	return m.inspector.numOfInspections
}

type Inspector struct {
	op   string
	l, r string

	numOfInspections int
}

func (i *Inspector) Inspect(old int) int {
	i.numOfInspections++

	var l, r int

	if i.l == "old" {
		l = old
	} else {
		l = mustAtoi(i.l)
	}
	if i.r == "old" {
		r = old
	} else {
		r = mustAtoi(i.r)
	}

	switch i.op {
	case "*":
		return (l * r) / 3
	case "+":
		return (l + r) / 3
	}

	panic("unknown op")
}

type Tester struct {
	divisibleBy int

	t, f int
}

func (t Tester) Test(worry int) int {
	if worry%t.divisibleBy == 0 {
		return t.t
	} else {
		return t.f
	}
}
