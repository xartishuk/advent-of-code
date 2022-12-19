package main

import (
	"fmt"
	"log"
)

func main() {
	in, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := PressureReleaseDuet(in)

	fmt.Println(result)
}

func PressureReleaseSolo(valves map[string]*Valve) int {
	return NewSolver(valves).MaxRelease(valves["AA"], 30)
}

func PressureReleaseDuet(valves map[string]*Valve) int {
	return NewDecisionPoint(valves, 26).MaxRelease()
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

func NewDecisionPoint(valves map[string]*Valve, time int) *DecisionPoint {
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

	return &DecisionPoint{
		distances: distances,
		closed:    closed,
		actor1: Actor{
			busyTime: 0,
			pos:      valves["AA"],
			action:   Move,
		},
		actor2: Actor{
			busyTime: 0,
			pos:      valves["AA"],
			action:   Move,
		},
		time: time,
	}
}

type DecisionPoint struct {
	distances map[*Valve]map[*Valve]int
	closed    map[*Valve]struct{}

	actor1, actor2 Actor

	time int
}

func (p *DecisionPoint) MaxRelease() int {
	//fmt.Printf("new DP: time=%d\tactor1 will %v %s in %d\tactor2 will %v %s in %d\n",
	//	p.time,
	//	p.actor1.action, p.actor1.pos.name, p.actor1.busyTime,
	//	p.actor2.action, p.actor2.pos.name, p.actor2.busyTime,
	//)

	if p.time <= 1 {
		return 0
	}

	if !p.actor1.IsReady() && !p.actor2.IsReady() {
		panic("none ready")
	}

	var max, this int
	var decisionPoints []*DecisionPoint

	if p.actor1.IsReady() && p.actor2.IsReady() {
		// both are ready
		if p.actor1.action == Open && p.actor2.action == Open {
			// preferably you need to check if they're at the same valve,
			// but it's probably a suboptimal route anyway - so I'll skip for now
			delete(p.closed, p.actor1.pos)
			delete(p.closed, p.actor2.pos)

			this = (p.actor1.pos.rate + p.actor2.pos.rate) * p.time
			if p.actor1.pos == p.actor2.pos {
				this /= 2
			}

			// both are ready again - pick pairs
			pairs := uniqueValvePairs(p.closed)
			for _, pair := range pairs {
				newActor1 := Actor{
					busyTime: p.distances[p.actor1.pos][pair.v1],
					pos:      pair.v1,
					action:   Move,
				}
				newActor2 := Actor{
					busyTime: p.distances[p.actor2.pos][pair.v2],
					pos:      pair.v2,
					action:   Move,
				}

				closestTime := min(newActor1.busyTime, newActor2.busyTime)
				newActor1.Tick(closestTime)
				newActor2.Tick(closestTime)

				newP := p.Copy(newActor1, newActor2, p.time-closestTime)
				decisionPoints = append(decisionPoints, newP)
			}
		} else if p.actor1.action == Move && p.actor2.action == Move {
			// both just arrived
			if _, ok := p.closed[p.actor1.pos]; p.actor1.pos == p.actor2.pos && ok {
				// same valve
				// pick next valve to move for actor1, while actor2 opens this one
				// situation where this valve is already open can only happen on start valve
				// for that case we use "both arrived to different already open valves" case below
				for v := range p.closed {
					newActor1 := Actor{
						busyTime: p.distances[p.actor1.pos][v],
						pos:      v,
						action:   Move,
					}
					newActor2 := Actor{
						busyTime: 1,
						pos:      p.actor2.pos,
						action:   Open,
					}

					closestTime := min(newActor1.busyTime, newActor2.busyTime)
					newActor1.Tick(closestTime)
					newActor2.Tick(closestTime)

					newP := p.Copy(newActor1, newActor2, p.time-closestTime)
					decisionPoints = append(decisionPoints, newP)
				}
			} else {
				// different valves
				_, closed1 := p.closed[p.actor1.pos]
				_, closed2 := p.closed[p.actor2.pos]

				if closed1 && closed2 {
					// both valves still closed - just open them
					newActor1 := Actor{
						busyTime: 1,
						pos:      p.actor1.pos,
						action:   Open,
					}
					newActor2 := Actor{
						busyTime: 1,
						pos:      p.actor2.pos,
						action:   Open,
					}

					closestTime := min(newActor1.busyTime, newActor2.busyTime)
					newActor1.Tick(closestTime)
					newActor2.Tick(closestTime)

					newP := p.Copy(newActor1, newActor2, p.time-closestTime)
					decisionPoints = append(decisionPoints, newP)
				} else if !closed1 && !closed2 {
					// both already open
					// also handles start valve
					// pick new valves to move to for both actors
					pairs := uniqueValvePairs(p.closed)
					for _, pair := range pairs {
						newActor1 := Actor{
							busyTime: p.distances[p.actor1.pos][pair.v1],
							pos:      pair.v1,
							action:   Move,
						}
						newActor2 := Actor{
							busyTime: p.distances[p.actor2.pos][pair.v2],
							pos:      pair.v2,
							action:   Move,
						}

						closestTime := min(newActor1.busyTime, newActor2.busyTime)
						newActor1.Tick(closestTime)
						newActor2.Tick(closestTime)

						newP := p.Copy(newActor1, newActor2, p.time-closestTime)
						decisionPoints = append(decisionPoints, newP)
					}
				} else {
					// one of the valves is already open
					// determine which is already open
					var alreadyOpen, stillClosed *Actor
					if closed1 {
						stillClosed = &p.actor1
						alreadyOpen = &p.actor2
					} else {
						alreadyOpen = &p.actor1
						stillClosed = &p.actor2
					}

					// pick new valve to move to for already open actor
					for v := range p.closed {
						newActor1 := Actor{
							busyTime: p.distances[alreadyOpen.pos][v],
							pos:      v,
							action:   Move,
						}
						newActor2 := Actor{
							busyTime: 1,
							pos:      stillClosed.pos,
							action:   Open,
						}

						closestTime := min(newActor1.busyTime, newActor2.busyTime)
						newActor1.Tick(closestTime)
						newActor2.Tick(closestTime)

						newP := p.Copy(newActor1, newActor2, p.time-closestTime)
						decisionPoints = append(decisionPoints, newP)
					}
				}
			}
		} else {
			// one of the actors just arrived - the other is set to open
			// determine which moved and which is set to open
			var actorMove, actorOpen *Actor
			if p.actor1.action == Move {
				actorMove = &p.actor1
				actorOpen = &p.actor2
			} else {
				actorOpen = &p.actor1
				actorMove = &p.actor2
			}

			delete(p.closed, actorOpen.pos)
			this = actorOpen.pos.rate * p.time

			if _, ok := p.closed[actorMove.pos]; ok {
				// moved actor's valve is still closed
				// set moved actor to open it
				// pick valve for opened actor to move to
				for v := range p.closed {
					newActor1 := Actor{
						busyTime: 1,
						pos:      actorMove.pos,
						action:   Open,
					}
					newActor2 := Actor{
						busyTime: p.distances[actorOpen.pos][v],
						pos:      v,
						action:   Move,
					}

					closestTime := min(newActor1.busyTime, newActor2.busyTime)
					newActor1.Tick(closestTime)
					newActor2.Tick(closestTime)

					newP := p.Copy(newActor1, newActor2, p.time-closestTime)
					decisionPoints = append(decisionPoints, newP)
				}
			} else {
				// moved actor's valve is already open
				// pick new valves for both to move to
				pairs := uniqueValvePairs(p.closed)
				for _, pair := range pairs {
					newActor1 := Actor{
						busyTime: p.distances[actorMove.pos][pair.v1],
						pos:      pair.v1,
						action:   Move,
					}
					newActor2 := Actor{
						busyTime: p.distances[actorOpen.pos][pair.v2],
						pos:      pair.v2,
						action:   Move,
					}

					closestTime := min(newActor1.busyTime, newActor2.busyTime)
					newActor1.Tick(closestTime)
					newActor2.Tick(closestTime)

					newP := p.Copy(newActor1, newActor2, p.time-closestTime)
					decisionPoints = append(decisionPoints, newP)
				}
			}
		}
	} else {
		// only one is ready
		// determine who is ready
		var actorReady, actorBusy *Actor
		if p.actor1.IsReady() {
			actorReady = &p.actor1
			actorBusy = &p.actor2
		} else {
			actorBusy = &p.actor1
			actorReady = &p.actor2
		}

		if actorReady.action == Move {
			// actor just arrived
			if _, ok := p.closed[actorReady.pos]; ok {
				// actor's valve is still closed
				// set actor to open it
				newActor1 := Actor{
					busyTime: 1,
					pos:      actorReady.pos,
					action:   Open,
				}
				newActor2 := Actor{
					busyTime: actorBusy.busyTime,
					pos:      actorBusy.pos,
					action:   actorBusy.action,
				}

				closestTime := min(newActor1.busyTime, newActor2.busyTime)
				newActor1.Tick(closestTime)
				newActor2.Tick(closestTime)

				newP := p.Copy(newActor1, newActor2, p.time-closestTime)
				decisionPoints = append(decisionPoints, newP)
			} else {
				// actor's valve is already open
				// pick new valve to move to
				for v := range p.closed {
					newActor1 := Actor{
						busyTime: p.distances[actorReady.pos][v],
						pos:      v,
						action:   Move,
					}
					newActor2 := Actor{
						busyTime: actorBusy.busyTime,
						pos:      actorBusy.pos,
						action:   actorBusy.action,
					}

					closestTime := min(newActor1.busyTime, newActor2.busyTime)
					newActor1.Tick(closestTime)
					newActor2.Tick(closestTime)

					newP := p.Copy(newActor1, newActor2, p.time-closestTime)
					decisionPoints = append(decisionPoints, newP)
				}
			}
		} else {
			// actor is set to open
			delete(p.closed, actorReady.pos)
			this = actorReady.pos.rate * p.time

			// pick new valve to move to
			for v := range p.closed {
				newActor1 := Actor{
					busyTime: p.distances[actorReady.pos][v],
					pos:      v,
					action:   Move,
				}
				newActor2 := Actor{
					busyTime: actorBusy.busyTime,
					pos:      actorBusy.pos,
					action:   actorBusy.action,
				}

				closestTime := min(newActor1.busyTime, newActor2.busyTime)
				newActor1.Tick(closestTime)
				newActor2.Tick(closestTime)

				newP := p.Copy(newActor1, newActor2, p.time-closestTime)
				decisionPoints = append(decisionPoints, newP)
			}
		}
	}

	for _, p := range decisionPoints {
		release := p.MaxRelease()
		if release > max {
			max = release
		}
	}

	//fmt.Printf("returning: %d+%d=%d\n", this, max, this+max)

	return this + max
}

func (p *DecisionPoint) Copy(actor1, actor2 Actor, time int) *DecisionPoint {
	closed := make(map[*Valve]struct{}, len(p.closed))
	for v := range p.closed {
		closed[v] = struct{}{}
	}

	return &DecisionPoint{
		distances: p.distances,
		closed:    closed,
		actor1:    actor1,
		actor2:    actor2,
		time:      time,
	}
}

type Actor struct {
	busyTime int
	pos      *Valve
	action   Action
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

type Action int

const (
	Move Action = iota + 1
	Open
)

func (a Action) String() string {
	switch a {
	case Move:
		return "move"
	case Open:
		return "open"
	}
	panic("unreachable")
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

	switch len(valves) {
	case 0:
		return nil
	case 1:
		return []ValvePair{{
			v1: valves[0],
			v2: valves[0],
		}}
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
