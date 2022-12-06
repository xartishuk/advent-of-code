package main

import (
	"bufio"
	"github.com/golang-collections/collections/stack"
	"os"
	"regexp"
	"strconv"
)

func readInput(filename string) ([]*stack.Stack, []Command, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	var crateLines []string

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		crateLines = append(crateLines, line)
	}

	// fill container stacks
	crates := make([]*stack.Stack, (len(crateLines[0])+1)/4)
	for j := range crates {
		s := stack.New()

		for i := len(crateLines) - 2; i >= 0; i-- {
			c := crateLines[i][(j*4)+1]
			if c == ' ' {
				goto finishStack
			}

			s.Push(rune(c))
		}

	finishStack:
		crates[j] = s
	}

	// read commands
	var commands []Command

	for s.Scan() {
		commands = append(commands, newCommand(s.Text()))
	}

	return crates, commands, s.Err()
}

var commandLayout = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func newCommand(s string) Command {
	parts := commandLayout.FindStringSubmatch(s)

	return Command{
		amount: mustAtoi(parts[1]),
		from:   mustAtoi(parts[2]),
		to:     mustAtoi(parts[3]),
	}
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
