package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := CampCleanupFull(in)

	fmt.Println(result)
}

func CampCleanupFull(pairs []AssignmentPair) int {
	var total int

	for _, pair := range pairs {
		if pair.Contains() {
			total++
		}
	}

	return total
}

type AssignmentPair struct {
	A1 Assignment
	A2 Assignment
}

func (p AssignmentPair) Contains() bool {
	return p.A1.In(p.A2) || p.A2.In(p.A1)
}

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) In(b Assignment) bool {
	return b.Start <= a.Start && a.End <= b.End
}

func readInput(filename string) ([]AssignmentPair, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var res []AssignmentPair

	s := bufio.NewScanner(f)

	for s.Scan() {
		assignments := strings.Split(s.Text(), ",")

		a1Sections := strings.Split(assignments[0], "-")
		a2Sections := strings.Split(assignments[1], "-")

		res = append(res, AssignmentPair{
			A1: Assignment{
				Start: mustAtoi(a1Sections[0]),
				End:   mustAtoi(a1Sections[1]),
			},
			A2: Assignment{
				Start: mustAtoi(a2Sections[0]),
				End:   mustAtoi(a2Sections[1]),
			},
		})
	}

	return res, s.Err()
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
