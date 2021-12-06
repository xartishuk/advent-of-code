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

	result := Hydrothermal(input)

	fmt.Println(result)
}

func Hydrothermal(vectors []Vector) int {
	// filter orthogonal vectors
	orthogonalVectors := make([]Vector, 0, len(vectors))
	for _, vector := range vectors {
		if vector.IsOrthogonal() {
			orthogonalVectors = append(orthogonalVectors, vector)
		}
	}

	// make vectors point in the same direction
	for i, v := range orthogonalVectors {
		if v.To.X < v.From.X || v.To.Y < v.From.Y {
			orthogonalVectors[i].SwapDirection()
		}
	}

	max := maxPoint(orthogonalVectors)
	floor := createFloor(max)

	for _, vector := range orthogonalVectors {
		floor.ApplyVector(vector)
	}

	return floor.CountDangerous()
}

type Vector struct {
	From Point
	To   Point
}

type Point struct {
	X int
	Y int
}

func (v *Vector) IsOrthogonal() bool {
	return v.IsVertical() || v.IsHorizontal()
}

func (v *Vector) IsVertical() bool {
	return v.From.X == v.To.X
}

func (v *Vector) IsHorizontal() bool {
	return v.From.Y == v.To.Y
}

func (v *Vector) SwapDirection() {
	v.From, v.To = v.To, v.From
}

type Floor [][]int

func (f Floor) ApplyVector(v Vector) {
	if v.IsVertical() {
		for j := v.From.Y; j <= v.To.Y; j++ {
			f[v.From.X][j]++
		}
	} else if v.IsHorizontal() {
		for i := v.From.X; i <= v.To.X; i++ {
			f[i][v.From.Y]++
		}
	} else {
		panic("unreachable for orthogonal")
	}
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
