package main

import (
	"fmt"
	"log"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Lanternfish(input)

	fmt.Println(result)
}

func Lanternfish(ages []byte) int {
	school := newSchool(ages)

	for i := 0; i < 80; i++ {
		school = school.AfterADay()
	}

	return school.Total()
}

type School [9]int

func newSchool(ages []byte) School {
	var school School

	for _, age := range ages {
		school[age]++
	}

	return school
}

func (s School) AfterADay() School {
	var next School
	next[0] = s[1]
	next[1] = s[2]
	next[2] = s[3]
	next[3] = s[4]
	next[4] = s[5]
	next[5] = s[6]
	next[6] = s[7] + s[0]
	next[7] = s[8]
	next[8] = s[0]

	return next
}

func (s School) Total() int {
	return s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6] + s[7] + s[8]
}
