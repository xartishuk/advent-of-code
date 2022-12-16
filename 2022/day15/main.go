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
	excluded := map[int]struct{}{}

	x, maxX := xBounds(sensors)

	fmt.Printf("x bound: %d..%d\n", x, maxX)
	//for _, s := range sensors {
	//	overlap := s.beaconDistance - s.DistanceToCoordinates(s.x, y)
	//	if overlap >= 0 {
	//		fmt.Printf("sensor %d.%d\t->\t%d.%d\treaches y=%d with beacon distance %d\toverlaps on %d\n", s.x, s.y, s.beacon.x, s.beacon.y, y, s.beaconDistance, overlap)
	//		fmt.Printf("overlap: %d..%d\n", s.x-overlap, s.x+overlap)
	//		for x := s.x - overlap; x <= s.x+overlap; x++ {
	//			excluded[x] = struct{}{}
	//		}
	//	} else {
	//		fmt.Printf("sensor %d.%d\t->\t%d.%d\tDOESN'T reach y=%d with beacon distance %d\toverlaps on %d\n", s.x, s.y, s.beacon.x, s.beacon.y, y, s.beaconDistance, overlap)
	//	}
	//}
	//
	for ; x <= maxX; x++ {
		for _, s := range sensors {
			if s.DistanceToCoordinates(x, y) <= s.beaconDistance {
				excluded[x] = struct{}{}
			}
		}
	}

	removeBeacons(excluded, sensors, y)

	return len(excluded)
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
