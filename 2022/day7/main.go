package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	root, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	sumOver10k, deleteSize := DeviceSpace(root)

	fmt.Println(sumOver10k)
	fmt.Println(deleteSize)
}

const sizeLimit = 100000

func DeviceSpace(root *Directory) (int, int) {
	var res int

	fmt.Println(root.String(0))

	for _, size := range root.ListDirectorySizes() {
		if size <= sizeLimit {
			res += size
		}
	}

	return res, 0
}

type Directory struct {
	subDirs map[string]*Directory
	files   map[string]int

	up *Directory

	cachedSize int
}

func (d *Directory) Size() int {
	if d.cachedSize == 0 {
		for _, sub := range d.subDirs {
			d.cachedSize += sub.Size()
		}
		for _, fileSize := range d.files {
			d.cachedSize += fileSize
		}
	}

	return d.cachedSize
}

// ListDirectorySizes includes itself
func (d *Directory) ListDirectorySizes() []int {
	sizes := make([]int, 0, len(d.subDirs)+1)
	sizes = append(sizes, d.Size())

	for _, sub := range d.subDirs {
		sizes = append(sizes, sub.ListDirectorySizes()...)
	}

	return sizes
}

func (d *Directory) String(level int) string {
	s := strings.Builder{}

	for filename, size := range d.files {
		s.WriteString(fmt.Sprintf("%s%s %d\n",
			strings.Repeat("\t", level), filename, size))
	}
	for dirName, dir := range d.subDirs {
		s.WriteString(fmt.Sprintf("%s/%s %d\n",
			strings.Repeat("\t", level), dirName, dir.Size()))
		s.WriteString(dir.String(level + 1))
	}

	return s.String()
}
