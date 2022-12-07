package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (*Directory, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	root := newDir(nil)
	cur := root

	s := bufio.NewScanner(f)

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		switch parts[0] {
		case "$":
			if parts[1] != "cd" {
				continue
			}

			switch parts[2] {
			case "/":
				cur = root
				continue
			case "..":
				cur = cur.up
				continue
			default:
				// check if dir already exists
				dir, ok := cur.subDirs[parts[2]]
				if ok {
					cur = dir
					continue
				}

				// create new dir
				newDir := newDir(cur)
				// save under current
				cur.subDirs[parts[2]] = newDir
				// cd to new dir
				cur = newDir
			}
		default:
			if parts[0] != "dir" {
				// file encountered
				cur.files[parts[1]] = mustAtoi(parts[0])
			}
			// don't bother saving directories here - we only care if we cd there, which is already handled
		}
	}

	return root, s.Err()
}

func newDir(up *Directory) *Directory {
	return &Directory{
		subDirs: make(map[string]*Directory),
		files:   make(map[string]int),
		up:      up,
	}
}

func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
