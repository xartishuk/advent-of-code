package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	in, err := readInput("input_test.txt", false)
	if err != nil {
		log.Fatal(err)
	}

	visited := MonkeyBusiness(in, 10000)

	fmt.Println(visited)
}

func MonkeyBusiness(monkeys []*Monkey, rounds int) int {
	worryCap := 1
	for _, monkey := range monkeys {
		worryCap *= monkey.tester.divisibleBy
	}

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.Turn(worryCap)
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

func (m *Monkey) Turn(worryCap int) {
	for {
		select {
		case item := <-m.items:
			item = m.inspector.Inspect(item)

			throwTo := m.tester.Test(item)

			m.monkeys[throwTo].items <- item % worryCap
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

	reduceWorry bool
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

	var new int

	switch i.op {
	case "*":
		new = l * r
	case "+":
		new = l + r
	default:
		panic("unknown op")
	}

	if i.reduceWorry {
		new /= 3
	}

	if i.l == "old" && i.r == "old" {
		if old > math.MaxInt32 {
			log.Printf("new = old * old: %d = %d * %d\n", new, old, old)
		}
	}

	return new
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
