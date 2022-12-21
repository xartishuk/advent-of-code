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

	result := RockTower(in)

	fmt.Println(result)
}

const width = 7

func RockTower(moves string) int {
	tower := map[int]struct{}{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
	}
	var max, mi int
	var dispenser RockDispenser

	for n := 0; n < 2022; n++ {
		r := dispenser.Next(max / width)

		for {
			move := moves[mi]
			mi = (mi + 1) % len(moves)

			switch move {
			case '>':
				if canMoveRight(tower, r) {
					r.MoveRight()
				}
			case '<':
				if canMoveLeft(tower, r) {
					r.MoveLeft()
				}
			default:
				panic("bad move")
			}

			if canFall(tower, r) {
				r.Fall()
			} else {
				// stopping rock
				points := r.All()
				var pi int
				for _, p := range points {
					pi = i(p.x, p.y)
					tower[pi] = struct{}{}
				}

				// last pi is always the highest
				if pi > max {
					max = pi
				}
				break
			}
		}
	}

	return max / width
}

func canMoveRight(tower map[int]struct{}, r Rock) bool {
	colliders := r.RightColliders()
	for _, c := range colliders {
		if c.x == width-1 {
			return false
		}
		if _, ok := tower[i(c.x+1, c.y)]; ok {
			return false
		}
	}

	return true
}

func canMoveLeft(tower map[int]struct{}, r Rock) bool {
	colliders := r.LeftColliders()
	for _, c := range colliders {
		if c.x == 0 {
			return false
		}
		if _, ok := tower[i(c.x-1, c.y)]; ok {
			return false
		}
	}

	return true
}

func canFall(tower map[int]struct{}, r Rock) bool {
	colliders := r.FallColliders()
	for _, c := range colliders {
		if _, ok := tower[i(c.x, c.y-1)]; ok {
			return false
		}
	}

	return true
}

func i(x, y int) int {
	return y*width + x
}
