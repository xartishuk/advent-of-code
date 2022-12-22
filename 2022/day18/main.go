package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	in, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := ExternalSurfaceArea(in)

	fmt.Println(result)
}

func TotalSurfaceArea(drops []Drop) int {
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

func ExternalSurfaceArea(drops []Drop) int {
	areaCounter := newExternalAreaCounter(drops)

	return areaCounter.CountFromStart()
}

func newExternalAreaCounter(drops []Drop) *ExternalAreaCounter {
	bounds := newBounds(drops)

	rocks := map[int]struct{}{}

	for _, d := range drops {
		rocks[bounds.I(d.x, d.y, d.z)] = struct{}{}
	}

	return &ExternalAreaCounter{
		bounds: bounds,
		rocks:  rocks,
		water:  map[int]struct{}{},
	}
}

type ExternalAreaCounter struct {
	bounds Bounds
	rocks  map[int]struct{}
	water  map[int]struct{}

	area int
}

func (c *ExternalAreaCounter) CountFromStart() int {
	c.count(c.bounds.minX, c.bounds.minY, c.bounds.minZ)

	return c.area
}

func (c *ExternalAreaCounter) count(x, y, z int) {
	c.water[c.bounds.I(x, y, z)] = struct{}{}

	if c.bounds.InX(x + 1) {
		i := c.bounds.I(x+1, y, z)
		if _, ok := c.rocks[i]; ok {
			c.area++
		} else if _, ok := c.water[i]; !ok {
			c.count(x+1, y, z)
		}
	}
	if c.bounds.InX(x - 1) {
		i := c.bounds.I(x-1, y, z)
		if _, ok := c.rocks[i]; ok {
			c.area++
		} else if _, ok := c.water[i]; !ok {
			c.count(x-1, y, z)
		}
	}
	if c.bounds.InY(y + 1) {
		i := c.bounds.I(x, y+1, z)
		if _, ok := c.rocks[i]; ok {
			c.area++
		} else if _, ok := c.water[i]; !ok {
			c.count(x, y+1, z)
		}
	}
	if c.bounds.InY(y - 1) {
		i := c.bounds.I(x, y-1, z)
		if _, ok := c.rocks[i]; ok {
			c.area++
		} else if _, ok := c.water[i]; !ok {
			c.count(x, y-1, z)
		}
	}
	if c.bounds.InZ(z + 1) {
		i := c.bounds.I(x, y, z+1)
		if _, ok := c.rocks[i]; ok {
			c.area++
		} else if _, ok := c.water[i]; !ok {
			c.count(x, y, z+1)
		}
	}
	if c.bounds.InZ(z - 1) {
		i := c.bounds.I(x, y, z-1)
		if _, ok := c.rocks[i]; ok {
			c.area++
		} else if _, ok := c.water[i]; !ok {
			c.count(x, y, z-1)
		}
	}
}

func newBounds(drops []Drop) Bounds {
	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	var maxX, maxY, maxZ int

	for _, d := range drops {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
		if d.z > maxZ {
			maxZ = d.z
		}
		if d.x < minX {
			minX = d.x
		}
		if d.y < minY {
			minY = d.y
		}
		if d.z < minZ {
			minZ = d.z
		}
	}

	return Bounds{
		minX:  minX - 1,
		minY:  minY - 1,
		minZ:  minZ - 1,
		maxX:  maxX + 1,
		maxY:  maxY + 1,
		maxZ:  maxZ + 1,
		sizeX: maxX - minX + 3,
		sizeY: maxY - minY + 3,
		sizeZ: maxZ - minZ + 3,
	}
}

type Bounds struct {
	minX, minY, minZ, maxX, maxY, maxZ int
	sizeX, sizeY, sizeZ                int
}

func (b Bounds) I(x, y, z int) int {
	return x*b.sizeY*b.sizeZ + y*b.sizeZ + z
}

func (b Bounds) InX(x int) bool { return b.minX <= x && x <= b.maxX }
func (b Bounds) InY(y int) bool { return b.minY <= y && y <= b.maxY }
func (b Bounds) InZ(z int) bool { return b.minZ <= z && z <= b.maxZ }

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
