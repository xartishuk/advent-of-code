package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string, reduceWorry bool) ([]*Monkey, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	monkeyInputs := strings.Split(string(buf), "\n\n")

	monkeys := make([]*Monkey, len(monkeyInputs))
	allItems := make([][]string, len(monkeyInputs))

	for i, monkeyInput := range monkeyInputs {
		lines := strings.Split(monkeyInput, "\n")

		for j := range lines {
			lines[j] = strings.TrimSpace(lines[j])
		}

		allItems[i] = strings.Split(strings.Split(lines[1], ": ")[1], ", ")

		inspectionParts := strings.Split(strings.Split(lines[2], "= ")[1], " ")

		monkeys[i] = &Monkey{
			inspector: Inspector{
				op:          inspectionParts[1],
				l:           inspectionParts[0],
				r:           inspectionParts[2],
				reduceWorry: reduceWorry,
			},
			tester: Tester{
				divisibleBy: mustAtoi(strings.Split(lines[3], " ")[3]),
				t:           mustAtoi(strings.Split(lines[4], " ")[5]),
				f:           mustAtoi(strings.Split(lines[5], " ")[5]),
			},
			monkeys: monkeys,
		}
	}

	var totalItems int

	for i := range allItems {
		totalItems += len(allItems[i])
	}

	for i := range allItems {
		monkeys[i].items = make(chan int, totalItems)

		for _, item := range allItems[i] {
			monkeys[i].items <- mustAtoi(item)
		}
	}

	return monkeys, nil
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
