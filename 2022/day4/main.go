package main

import (
	"bufio"
	"fmt"
	"github.com/xartishuk/advent-of-code/util"
	"log"
	"os"
	"strings"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	full, partial := CampCleanup(in)

	fmt.Println(full)
	fmt.Println(partial)
}

func CampCleanup(pairs []AssignmentPair) (full, partial int) {
	for _, pair := range pairs {
		if pair.Contains() {
			full++
			partial++
			continue
		}
		if pair.Intersects() {
			partial++
		}
	}

	return full, partial
}

type AssignmentPair struct {
	A1 Assignment
	A2 Assignment
}

func (p AssignmentPair) Contains() bool {
	return p.A1.IsIn(p.A2) || p.A2.IsIn(p.A1)
}

func (p AssignmentPair) Intersects() bool {
	return p.A1.ContainsSection(p.A2.Start) || p.A2.ContainsSection(p.A1.Start)
}

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) IsIn(b Assignment) bool {
	return b.Start <= a.Start && a.End <= b.End
}

func (a Assignment) ContainsSection(s int) bool {
	return a.Start <= s && s <= a.End
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
				Start: util.MustAtoi(a1Sections[0]),
				End:   util.MustAtoi(a1Sections[1]),
			},
			A2: Assignment{
				Start: util.MustAtoi(a2Sections[0]),
				End:   util.MustAtoi(a2Sections[1]),
			},
		})
	}

	return res, s.Err()
}
