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

	excluded, beacon := BeaconExclusionZone(in, 2000000)

	fmt.Println(excluded)
	fmt.Println(beacon)
}

func BeaconExclusionZone(sensors []Sensor, y int) (int, int) {
	// part 1
	excluded := map[int]struct{}{}

	for _, s := range sensors {
		overlap := s.beaconDistance - s.DistanceToCoordinates(s.x, y)
		if overlap >= 0 {
			//fmt.Printf("sensor %d.%d\t->\t%d.%d\treaches y=%d with beacon distance %d\toverlaps on %d\n", s.x, s.y, s.beacon.x, s.beacon.y, y, s.beaconDistance, overlap)
			//fmt.Printf("overlap: %d..%d\n", s.x-overlap, s.x+overlap)
			for x := s.x - overlap; x <= s.x+overlap; x++ {
				excluded[x] = struct{}{}
			}
		} else {
			//fmt.Printf("sensor %d.%d\t->\t%d.%d\tDOESN'T reach y=%d with beacon distance %d\toverlaps on %d\n", s.x, s.y, s.beacon.x, s.beacon.y, y, s.beaconDistance, overlap)
		}
	}

	removeBeacons(excluded, sensors, y)

	// part 2
	bx, by := findBeacon(sensors, y*2)

	return len(excluded), bx*4000000 + by
}

func findBeacon(sensors []Sensor, size int) (int, int) {
	possibleBeacons := map[int]struct{}{}

	for _, s := range sensors {
		x := s.x
		y := s.y - s.beaconDistance - 1
		// top to right
		for x <= s.x+s.beaconDistance+1 && y <= s.y {
			if oob(size, x) || oob(size, y) {
				x++
				y++
				continue
			}
			possibleBeacons[x*(size+1)+y] = struct{}{}

			x++
			y++
		}
		x -= 2
		// right to down
		for x >= s.x && y <= s.y+s.beaconDistance+1 {
			if oob(size, x) || oob(size, y) {
				x--
				y++
				continue
			}
			possibleBeacons[x*(size+1)+y] = struct{}{}

			x--
			y++
		}
		y -= 2
		// down to left
		for x >= s.x-s.beaconDistance-1 && y <= s.y {
			if oob(size, x) || oob(size, y) {
				x--
				y--
				continue
			}
			possibleBeacons[x*(size+1)+y] = struct{}{}

			x--
			y--
		}
		x += 2
		// left to up
		for x <= s.x && y >= s.y-s.beaconDistance-1 {
			if oob(size, x) || oob(size, y) {
				x++
				y--
				continue
			}
			possibleBeacons[x*(size+1)+y] = struct{}{}

			x++
			y--
		}
	}

	fmt.Printf("number of possible beacons: %d\n", len(possibleBeacons))

nextBeacon:
	for i := range possibleBeacons {
		x, y := i/(size+1), i%(size+1)

		for _, s := range sensors {
			if s.DistanceToCoordinates(x, y) <= s.beaconDistance {
				continue nextBeacon
			}
		}

		return x, y
	}

	panic("unreachable")
}

type Sensor struct {
	Point
	beacon         Point
	beaconDistance int
}

func xBounds(sensors []Sensor) (l, r int) {
	var maxDistance, minX, maxX int

	for _, s := range sensors {
		if s.beaconDistance > maxDistance {
			maxDistance = s.beaconDistance
		}

		if s.x < minX {
			minX = s.x
		}
		if s.beacon.x < minX {
			minX = s.beacon.x
		}

		if s.x > maxX {
			maxX = s.x
		}
		if s.beacon.x > maxX {
			maxX = s.beacon.x
		}
	}

	return minX - maxDistance, maxX + maxDistance
}

func removeBeacons(excluded map[int]struct{}, sensors []Sensor, y int) {
	for _, s := range sensors {
		if s.beacon.y == y {
			delete(excluded, s.beacon.x)
		}
	}
}

type Point struct {
	x, y int
}

func (p Point) DistanceToPoint(to Point) int {
	return absDif(p.x, to.x) + absDif(p.y, to.y)
}

func (p Point) DistanceToCoordinates(x, y int) int {
	return absDif(p.x, x) + absDif(p.y, y)
}

func absDif(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func oob(size, v int) bool {
	return v < 0 || v > size
}
