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

	result := SurfaceArea(in)

	fmt.Println(result)
}

func SurfaceArea(drops []Drop) int {
	for i := 0; i < len(drops); i++ {
		for j := i + 1; j < len(drops); j++ {
			if drops[i].Sides == 0 {
				break
			}
			if drops[j].Sides == 0 {
				continue
			}
			if drops[i].IsTouching(drops[j]) {
				drops[i].Sides--
				drops[j].Sides--
			}
		}
	}

	var area int
	for _, d := range drops {
		area += d.Sides
	}

	return area
}

type Drop struct {
	x, y, z int
	Sides   int
}

func (d *Drop) IsTouching(to Drop) bool {
	return abs(d.x-to.x)+abs(d.y-to.y)+abs(d.z-to.z) == 1
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
