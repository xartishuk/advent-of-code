package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func readInput(filename string) ([]Sensor, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	format := regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)

	var sensors []Sensor

	for s.Scan() {
		parts := format.FindStringSubmatch(s.Text())

		s := Sensor{
			Point: Point{
				x: mustAtoi(parts[1]),
				y: mustAtoi(parts[2]),
			},
			beacon: Point{
				x: mustAtoi(parts[3]),
				y: mustAtoi(parts[4]),
			},
		}

		s.beaconDistance = s.DistanceToPoint(s.beacon)

		sensors = append(sensors, s)
	}

	return sensors, nil
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
