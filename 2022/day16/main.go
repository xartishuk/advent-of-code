package main

import (
	"fmt"
	"log"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := PressureReleaseSolo(in)

	fmt.Println(result)
}

func PressureReleaseSolo(valves map[string]*Valve) int {
	return NewSolver(valves).MaxRelease(valves["AA"], 30)
}

func PressureReleaseDuet(valves map[string]*Valve) int {
	return NewDuetSolver(valves).MaxRelease(valves["AA"], 26)
}

func NewSolver(valves map[string]*Valve) *Solver {
	// use Floyd–Warshall for all distances
	distances := make(map[*Valve]map[*Valve]int, len(valves))
	for _, from := range valves {
		distances[from] = make(map[*Valve]int, len(valves))

		for _, to := range valves {
			distances[from][to] = len(valves) * 10
		}

		distances[from][from] = 0

		for _, to := range from.tunnels {
			distances[from][to] = 1
		}
	}

	for _, k := range valves {
		for _, i := range valves {
			for _, j := range valves {
				if distances[i][k]+distances[k][j] < distances[i][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}
	}

	closed := make(map[*Valve]struct{}, len(valves))
	for _, v := range valves {
		if v.rate != 0 {
			closed[v] = struct{}{}
		}
	}

	return &Solver{
		distances: distances,
		closed:    closed,
	}
}

type Solver struct {
	distances map[*Valve]map[*Valve]int
	closed    map[*Valve]struct{}
}

func (s *Solver) MaxRelease(from *Valve, time int) int {
	if time <= 0 {
		return 0
	}

	var max int
	// make a copy of s.closed
	closedOnThisStep := make(map[*Valve]struct{}, len(s.closed))
	for v := range s.closed {
		closedOnThisStep[v] = struct{}{}
	}

	for to := range closedOnThisStep {
		timeLeft := time - s.distances[from][to] - 1 // -1 for opening

		if timeLeft < 0 {
			continue
		}

		delete(s.closed, to)
		{
			toMax := s.MaxRelease(to, timeLeft)

			total := toMax + to.rate*timeLeft

			if total > max {
				max = total
			}
		}
		s.closed[to] = struct{}{}
	}

	return max
}

type DuetSolver struct {
	Solver

	actor1, actor2 Actor
}

func (s *DuetSolver) MaxRelease(from *Valve, time int) int {
	if time <= 0 {
		return 0
	}

	if s.actor1.IsReady() && s.actor2.IsReady() {
		// both are ready - pick pairs
		pairs := uniqueValvePairs(s.closed)

		for _, p := range pairs {

		}
	}

	var max int
	// make a copy of s.closed
	closedOnThisStep := make(map[*Valve]struct{}, len(s.closed))
	for v := range s.closed {
		closedOnThisStep[v] = struct{}{}
	}

	for to := range closedOnThisStep {
		timeLeft := time - s.distances[from][to] - 1 // -1 for opening

		if timeLeft < 0 {
			continue
		}

		delete(s.closed, to)
		{
			toMax := s.MaxRelease(to, timeLeft)

			total := toMax + to.rate*timeLeft

			if total > max {
				max = total
			}
		}
		s.closed[to] = struct{}{}
	}

	return max
}

type Actor struct {
	busyTime int
}

func (a *Actor) Tick(v int) {
	a.busyTime -= v

	if a.busyTime < 0 {
		panic("actor ticked overtime")
	}
}

func (a *Actor) IsReady() bool {
	return a.busyTime == 0
}

type Valve struct {
	name    string
	rate    int
	tunnels []*Valve
}

type ValvePair struct {
	v1, v2 *Valve
}

func uniqueValvePairs(valveMap map[*Valve]struct{}) []ValvePair {
	valves := make([]*Valve, 0, len(valveMap))
	for v := range valveMap {
		valves = append(valves, v)
	}

	pairs := make([]ValvePair, 0, pairNum(len(valves)))

	for i := 0; i < len(valves)-1; i++ {
		for j := i + 1; j < len(valves); j++ {
			pairs = append(pairs, ValvePair{
				v1: valves[i],
				v2: valves[j],
			})
		}
	}

	return pairs
}

func pairNum(x int) int {
	return x * (x - 1) / 2
}
