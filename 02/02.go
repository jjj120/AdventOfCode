package main

import (
	"fmt"
	"math"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 02
const selectExample = false

const WIDTH_X = 3
const WIDTH_Y = 3

func Step(v, s aoc.Vec2d) aoc.Vec2d {
	v_new := v.Add(s)
	if v_new.X >= WIDTH_X || v_new.Y >= WIDTH_Y || v_new.X < 0 || v_new.Y < 0 {
		return v
	}
	return v_new
}

func VecToInt(v aoc.Vec2d) int {
	return 1 + v.X + v.Y*WIDTH_X
}

func handleLines(lines []string) int {
	pos := aoc.Vec2d{X: 1, Y: 1}
	code := make([]int, 0, len(lines))
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

		code = append(code, VecToInt(pos))
	}
	sum := 0
	for i, num := range code {
		sum += int(float64(math.Pow10(len(code)-i-1))) * num
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 1985, "Example wrong!")
	}
}
