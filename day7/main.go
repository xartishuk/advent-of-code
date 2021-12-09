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

	result := WhaleChase(input, true)

	fmt.Println(result)
}

func WhaleChase(crabs []int, complexFuelConsumption bool) int {
	crabLine := makeCrabLine(crabs, complexFuelConsumption)

	var fuel, start, end int
	end = len(crabLine)-1

	for start < end {
		startCost := crabLine[start].CostToMove()
		endCost := crabLine[end].CostToMove()
		if startCost < endCost {
			fuel += startCost

			crabLine[start+1].Add(crabLine[start])

			start++
		} else {
			fuel += endCost

			crabLine[end-1].Add(crabLine[end])

			end--
		}
	}

	return fuel
}

type Fleet interface {
	CostToMove() int
	Add(Fleet)
}

func makeCrabLine(crabs []int, complexFuelConsumption bool) []Fleet {
	line := make([]Fleet, max(crabs)+1)
	if complexFuelConsumption {
		complexLine := makeComplexCrabLine(crabs)
		for i := range complexLine {
			line[i] =  &complexLine[i]
		}
	} else {
		simpleLine := makeSimpleCrabLine(crabs)
		for i := range simpleLine {
			line[i] =  &simpleLine[i]
		}
	}

	return line
}

type SimpleFleet struct {
	numOfSubs int
}

func (f *SimpleFleet) CostToMove() int {
	return f.numOfSubs
}

func (f *SimpleFleet) Add(from Fleet) {
	f.numOfSubs += from.CostToMove()
}

func makeSimpleCrabLine(crabs []int) []SimpleFleet {
	crabLine := make([]SimpleFleet, max(crabs)+1)
	for _, position := range crabs {
		crabLine[position].numOfSubs++
	}

	return crabLine
}

type ComplexFleet struct {
	fleetsByCost map[int]*SimpleFleet
}

func (f *ComplexFleet) CostToMove() int {
	var sum int
	for cost, fleet := range f.fleetsByCost {
		sum+=cost*fleet.CostToMove()
	}

	return sum
}

func (f *ComplexFleet) Add(fromI Fleet) {
	from := fromI.(*ComplexFleet)
	for fromCost, fromFleet := range from.fleetsByCost {
		toFleet, ok := f.fleetsByCost[fromCost+1]
		if !ok {
			toFleet = new(SimpleFleet)
		}
		toFleet.Add(fromFleet)
		f.fleetsByCost[fromCost+1] = toFleet
	}
}

func makeComplexCrabLine(crabs []int) []ComplexFleet {
	crabLine := make([]ComplexFleet, max(crabs)+1)
	for i := range crabLine {
		crabLine[i].fleetsByCost = make(map[int]*SimpleFleet, 1)
		crabLine[i].fleetsByCost[1] = new(SimpleFleet)
	}
	for _, position := range crabs {
		crabLine[position].fleetsByCost[1].numOfSubs++
	}

	return crabLine
}


func max(arr []int) int {
	max := -1
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}
