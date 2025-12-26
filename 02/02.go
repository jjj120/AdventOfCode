package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 02
const selectExample = false

var validPos = map[aoc.Vec2d]string{
	{X: 2, Y: 0}: "1",
	{X: 1, Y: 1}: "2",
	{X: 2, Y: 1}: "3",
	{X: 3, Y: 1}: "4",
	{X: 0, Y: 2}: "5",
	{X: 1, Y: 2}: "6",
	{X: 2, Y: 2}: "7",
	{X: 3, Y: 2}: "8",
	{X: 4, Y: 2}: "9",
	{X: 1, Y: 3}: "A",
	{X: 2, Y: 3}: "B",
	{X: 3, Y: 3}: "C",
	{X: 2, Y: 4}: "D",
}

func Step(v, s aoc.Vec2d) aoc.Vec2d {
	v_new := v.Add(s)

	if _, ok := validPos[v_new]; ok {
		return v_new
	}
	return v
}

func handleLines(lines []string) string {
	pos := aoc.Vec2d{X: 0, Y: 2}
	code := ""
	for _, line := range lines {
		for _, r := range line {
			switch r {
			case 'U':
				pos = Step(pos, aoc.Vec2d{X: 0, Y: -1})
			case 'D':
				pos = Step(pos, aoc.Vec2d{X: 0, Y: 1})
			case 'L':
				pos = Step(pos, aoc.Vec2d{X: -1, Y: 0})
			case 'R':
				pos = Step(pos, aoc.Vec2d{X: 1, Y: 0})
			default:
			}
		}

		code += validPos[pos]
	}
	return code
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %s\n", sum)
	if selectExample {
		aoc.Assert(sum == "5DB3", "Example wrong!")
	}
}
