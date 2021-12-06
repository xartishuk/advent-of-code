package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Hydrothermal(input, false)

	fmt.Println(result)
}

func Hydrothermal(vectors []Vector, filterOrthogonal bool) int {
	// filter orthogonal vectors
	if filterOrthogonal {
		orthogonalVectors := make([]Vector, 0, len(vectors))
		for _, vector := range vectors {
			if vector.IsOrthogonal() {
				orthogonalVectors = append(orthogonalVectors, vector)
			}
		}

		vectors = orthogonalVectors
	}

	max := maxPoint(vectors)
	floor := createFloor(max)

	for _, vector := range vectors {
		floor.ApplyVector(vector)
	}

	return floor.CountDangerous()
}

type Vector struct {
	From      Point
	To        Point
	Direction Point
}

type Point struct {
	X int
	Y int
}

func (v Vector) IsOrthogonal() bool {
	return v.From.X == v.To.X || v.From.Y == v.To.Y
}

type Floor [][]int

func (f Floor) ApplyVector(v Vector) {
	i := v.From.X
	j := v.From.Y
	for i != v.To.X || j != v.To.Y {
		f[i][j]++

		i += v.Direction.X
		j += v.Direction.Y
	}

	f[v.To.X][v.To.Y]++
}

func (f Floor) CountDangerous() int {
	var count int
	for i := range f {
		for j := range f[i] {
			if f[i][j] > 1 {
				count++
			}
		}
	}

	return count
}

func (f Floor) String() string {
	var builder strings.Builder
	for i := range f {
		for j := range f[i] {
			builder.WriteString(strconv.Itoa(f[i][j]))
			builder.WriteRune('\t')
		}
		builder.WriteRune('\n')
	}
	builder.WriteString("==============================================================================")

	return builder.String()
}

func createFloor(p Point) Floor {
	f := make([][]int, p.X+1)
	for i := range f {
		f[i] = make([]int, p.Y+1)
	}

	return f
}

func maxPoint(vectors []Vector) Point {
	var max Point
	for _, v := range vectors {
		if v.To.X > max.X {
			max.X = v.To.X
		}
		if v.From.X > max.X {
			max.X = v.From.X
		}
		if v.To.Y > max.Y {
			max.Y = v.To.Y
		}
		if v.From.Y > max.Y {
			max.Y = v.From.Y
		}
	}

	return max
}
