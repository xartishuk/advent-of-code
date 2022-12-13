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

	result := DistressSignal(in)

	fmt.Println(result)
}

func DistressSignal(pairs []PacketPair) int {
	var sum int

	for i := range pairs {
		if pairs[i].l.Compare(pairs[i].r) == -1 {
			sum += i + 1
		}
	}

	return sum
}

type PacketPair struct {
	l, r Set
}

type SetMember interface {
	Set() Set
	Compare(SetMember) int
}

type Set struct {
	arr []SetMember
}

func (l Set) Set() Set {
	return l
}

func (l Set) Compare(r SetMember) int {
	return l.compare(r.Set())
}

func (l Set) compare(r Set) int {
	for i := 0; ; i++ {
		// set end
		if i == len(l.arr) || i == len(r.arr) {
			cmp := 0
			if i == len(l.arr) {
				cmp--
			}
			if i == len(r.arr) {
				cmp++
			}

			return cmp
		}

		// compare element i
		if cmp := l.arr[i].Compare(r.arr[i]); cmp != 0 {
			return cmp
		}
	}
}

type SetInt int

func (s SetInt) Set() Set {
	return Set{[]SetMember{s}}
}

func (s SetInt) Compare(r SetMember) int {
	if i, ok := r.(SetInt); ok {
		return unit(s - i)
	}

	return s.Set().Compare(r)
}

func unit(x SetInt) int {
	if x < -1 {
		return -1
	}
	if x > 1 {
		return 1
	}

	return int(x)
}
