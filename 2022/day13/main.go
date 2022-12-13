package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := DistressSignalDecoder(in)

	fmt.Println(result)
}

func DistressSignalPairs(pairs []PacketPair) int {
	var sum int

	for i := range pairs {
		if pairs[i].l.Compare(pairs[i].r) == -1 {
			sum += i + 1
		}
	}

	return sum
}

func DistressSignalDecoder(pairs []PacketPair) int {
	divider1 := &Set{arr: []SetMember{Set{arr: []SetMember{SetInt(2)}}}}
	divider2 := &Set{arr: []SetMember{Set{arr: []SetMember{SetInt(6)}}}}

	packets := make([]*Set, 0, (len(pairs)*2)+2)
	packets = append(packets, divider1, divider2)

	for i := range pairs {
		packets = append(packets, &pairs[i].l)
		packets = append(packets, &pairs[i].r)
	}

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].compare(*packets[j]) == -1
	})

	var i1, i2 int

	for i := range packets {
		if packets[i] == divider1 {
			i1 = i + 1
		}
		if packets[i] == divider2 {
			i2 = i + 1
		}
	}

	return i1 * i2
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
