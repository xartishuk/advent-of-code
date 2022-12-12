package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	in, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	visited := CRTCPU(in)

	fmt.Println(visited)
}

const screenWidth = 40

func CRTCPU(instructions []string) int {
	var sum int

	cpu := newCPU(instructions, os.Stdout)

	for cpu.Draw() {
		if (cpu.Cycle()-20)%screenWidth == 0 {
			sum += cpu.SignalStrength()
		}

		cpu.Execute()
	}

	return sum
}

type CPU struct {
	x int

	// instruction pointer doubles as crt beam position
	ip int

	instructions []string

	screen *os.File
}

func newCPU(instructions []string, screen *os.File) *CPU {
	return &CPU{
		x:            1,
		instructions: instructions,
		screen:       screen,
	}
}

func (c *CPU) Draw() bool {
	// sprite intersects CRT beam
	if c.x-1 <= c.beam() && c.beam() <= c.x+1 {
		c.screen.WriteString("#")
	} else {
		c.screen.WriteString(".")
	}

	if (c.ip+1)%screenWidth == 0 {
		c.screen.WriteString("\n")
	}

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

func (c *CPU) beam() int {
	return c.ip % screenWidth
}
