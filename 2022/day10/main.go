package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	in, err := readInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	visited := CRTCPU(in)

	fmt.Println(visited)
}

func CRTCPU(instructions []string) int {
	var sum int

	cpu := newCPU(instructions)

	for cpu.LoadOp() {
		if (cpu.Cycle()-20)%40 == 0 {
			sum += cpu.SignalStrength()
		}

		cpu.Execute()
	}

	return sum
}

type CPU struct {
	x int

	ip int

	instructions []string
}

func newCPU(instructions []string) *CPU {
	return &CPU{
		x:            1,
		instructions: instructions,
	}
}

func (c *CPU) LoadOp() bool {
	return c.ip < len(c.instructions)
}

func (c *CPU) Execute() {
	v, err := strconv.Atoi(c.instructions[c.ip])
	if err == nil {
		c.x += v
	}

	c.ip++
}

func (c *CPU) Cycle() int {
	return c.ip + 1
}

func (c *CPU) SignalStrength() int {
	return c.Cycle() * c.x
}
