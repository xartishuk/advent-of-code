package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func readInput(filename string) (map[string]*Valve, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	valves := map[string]*Valve{}
	tunnelInputs := map[string]string{}

	for s.Scan() {
		parts := strings.Split(s.Text(), "; ")

		valves[parts[0][6:8]] = &Valve{
			name: parts[0][6:8],
			rate: mustAtoi(parts[0][23:]),
			//opened: mustAtoi(parts[0][23:]) == 0,
		}

		tunnelInputs[parts[0][6:8]] = parts[1]
	}

	for name, ti := range tunnelInputs {
		v := valves[name]
		tunnels := strings.Split(ti[strings.IndexAny(ti, upper):], ", ")

		for _, t := range tunnels {
			v.tunnels = append(v.tunnels, valves[t])
		}
	}

	return valves, s.Err()
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
