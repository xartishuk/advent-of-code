package main

import (
	"bufio"
	"math"
	"os"
)

func readInput(filename string) (start, end *Point, grid [][]*Point, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for i := 0; s.Scan(); i++ {
		line := s.Text()

		row := make([]*Point, 0, len(line))

		for j, h := range line {
			p := &Point{
				i:      i,
				j:      j,
				height: int(h),
				g:      math.MaxInt,
			}

			if h == 'S' {
				end = p
			}
			if h == 'E' {
				start = p
			}

			row = append(row, p)
		}

		grid = append(grid, row)
	}

	end.height = 'a'
	start.height = 'z'
	start.g = 0

	return start, end, grid, s.Err()
}
