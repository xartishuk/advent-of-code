package main

import (
	"container/heap"
	"fmt"
	"log"
)

func main() {
	start, end, grid, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := HillClimb(start, end, grid)

	fmt.Println(result)
}

func HillClimb(start, end *Point, grid [][]*Point) int {
	iMax, jMax := len(grid)-1, len(grid[0])-1

	discovered := &PointHeap{
		arr: []*Point{start},
		end: end,
	}
	heap.Init(discovered)

	for {
		cur := heap.Pop(discovered).(*Point)

		if cur == end {
			break
		}

		// down exists
		if cur.i+1 <= iMax {
			discoverNeighbour(discovered, cur, grid[cur.i+1][cur.j])
		}
		// up exists
		if cur.i-1 >= 0 {
			discoverNeighbour(discovered, cur, grid[cur.i-1][cur.j])
		}
		// right exists
		if cur.j+1 <= jMax {
			discoverNeighbour(discovered, cur, grid[cur.i][cur.j+1])
		}
		// left exists
		if cur.j-1 >= 0 {
			discoverNeighbour(discovered, cur, grid[cur.i][cur.j-1])
		}
	}

	return end.g
}

func discoverNeighbour(discovered *PointHeap, cur, neighbour *Point) {
	// climbable
	if cur.height-1 <= neighbour.height {
		// path through cur is closer
		if cur.g+1 < neighbour.g {
			neighbour.g = cur.g + 1

			// update in discovered
			if i := discovered.Has(neighbour); i != -1 {
				heap.Remove(discovered, i)
			}

			heap.Push(discovered, neighbour)
		}
	}
}

type Point struct {
	i, j   int
	height int

	// start to this
	g int
}

// F score is start to end through this Point. F = g + heuristic(p, end)
func (p *Point) F(to *Point) int {
	return p.g + absDiff(to.i, p.i) + absDiff(to.j, p.j)
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

type PointHeap struct {
	arr []*Point
	end *Point
}

func (h PointHeap) Len() int           { return len(h.arr) }
func (h PointHeap) Less(i, j int) bool { return h.arr[i].F(h.end) < h.arr[j].F(h.end) }
func (h PointHeap) Swap(i, j int)      { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func (h *PointHeap) Push(p any) {
	h.arr = append(h.arr, p.(*Point))
}

func (h *PointHeap) Pop() any {
	p := h.arr[len(h.arr)-1]

	h.arr = h.arr[:len(h.arr)-1]

	return p
}

func (h *PointHeap) Has(p *Point) int {
	for i := range h.arr {
		if h.arr[i] == p {
			return i
		}
	}

	return -1
}
