package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 01
const selectExample = false

func IntAbs(v int) int {
	if v <= 0 {
		return -v
	}
	return v
}

type PosHeading struct {
	pos     aoc.Vec2d
	heading aoc.Vec2d
}

func (ph *PosHeading) WalkStraight(m *map[aoc.Vec2d]bool, len int) bool {
	for range len {
		ph.pos = ph.pos.Add(ph.heading)

		if included, ok := (*m)[ph.pos]; included && ok {
			return true
		}
		(*m)[ph.pos] = true
	}
	return false
}

func (ph *PosHeading) TurnRight() {
	ph.heading = ph.heading.Rotate90()
}

func (ph *PosHeading) TurnLeft() {
	ph.heading = ph.heading.Rotate270()
}

func (ph *PosHeading) Walk(m *map[aoc.Vec2d]bool, instr string) bool {
	dir := string(instr[0])
	len, err := strconv.Atoi(string(instr[1:]))
	aoc.Check(err)

	if dir == "R" {
		ph.TurnRight()
	} else if dir == "L" {
		ph.TurnLeft()
	} else {
		aoc.Assert(false, "Got unknown heading!")
	}

	return ph.WalkStraight(m, len)
}

func handleLines(lines []string) int {
	ph := PosHeading{pos: aoc.Vec2d{X: 0, Y: 0}, heading: aoc.Vec2d{X: 0, Y: -1}}
	visited := make(map[aoc.Vec2d]bool)

	for _, line := range lines {
		for _, dir := range strings.Split(line, ", ") {
			// fmt.Println(dir)
			if ph.Walk(&visited, dir) {
				break
			}
			// fmt.Println(ph)
		}
	}

	return IntAbs(ph.pos.X) + IntAbs(ph.pos.Y)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 4, "Example wrong!")
	}
}
