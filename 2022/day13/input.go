package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]PacketPair, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	inputPairs := strings.Split(string(buf), "\n\n")

	var pairs []PacketPair

	for _, inputPair := range inputPairs {
		pairParts := strings.Split(inputPair, "\n")

		l, _ := parseSet(pairParts[0][1:])
		r, _ := parseSet(pairParts[1][1:])

		pairs = append(pairs, PacketPair{
			l: l,
			r: r,
		})
	}

	return pairs, nil
}

func parseSet(s string) (Set, string) {
	var set Set

	for len(s) > 0 {
		if s[0] == '[' {
			sub, rest := parseSet(s[1:])

			set.arr = append(set.arr, sub)

			s = rest
		} else if s[0] == ']' {
			return set, s[1:]
		} else {
			end := strings.IndexAny(s, ",]")
			if end == 0 {
				// found ',' at start - skipping
				s = s[1:]
				continue
			}
			if end == -1 {
				end = len(s) - 1
			}

			set.arr = append(set.arr, mustAtoiSetInt(s[:end]))

			s = s[end:]
		}
	}

	panic("unreachable")
}

func mustAtoiSetInt(s string) SetInt {
	v, _ := strconv.Atoi(s)
	return SetInt(v)
}
