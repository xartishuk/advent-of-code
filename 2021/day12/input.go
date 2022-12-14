package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func readInput(filename string) (*Cave, *Cave, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var start, end *Cave
	caves := map[string]*Cave{}

	for i := 0; s.Scan(); i++ {
		pair := strings.Split(s.Text(), "-")

		l, ok := caves[pair[0]]
		if !ok {
			l = &Cave{
				name:    pair[0],
				isSmall: unicode.IsLower(rune(pair[0][0])),
			}
			caves[pair[0]] = l
		}

		r, ok := caves[pair[1]]
		if !ok {
			r = &Cave{
				name:    pair[1],
				isSmall: unicode.IsLower(rune(pair[1][0])),
			}
			caves[pair[1]] = r
		}

		if l.name == "start" {
			start = l
		}
		if l.name == "end" {
			end = l
		}
		if r.name == "start" {
			start = r
		}
		if r.name == "end" {
			end = r
		}

		if r != start {
			l.neighbours = append(l.neighbours, r)
		}
		if l != start {
			r.neighbours = append(r.neighbours, l)
		}
	}

	return start, end, nil
}
