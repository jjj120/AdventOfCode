package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 04
const selectExample = false

func checkMovable(rollsMap [][]bool, x, y int) bool {
	height := len(rollsMap)
	width := len(rollsMap[0])
	adjRolls := 0
	for _, xOff := range []int{-1, 0, 1} {
		for _, yOff := range []int{-1, 0, 1} {
			// check if in bounds
			if x+xOff >= 0 && x+xOff < width && y+yOff >= 0 && y+yOff < height {
				if rollsMap[y+yOff][x+xOff] {
					adjRolls++
				}
			}
		}
	}
	return adjRolls < 5 // +1 because we count the roll itself too
}

func countAccess(rollsMap [][]bool) int {
	movable := 0
	movableMap := make([][]bool, 0, len(rollsMap))
	for y, line := range rollsMap {
		movableLine := make([]bool, len(line))
		for x, e := range line {
			movableLine[x] = false
			if e && checkMovable(rollsMap, x, y) {
				movable++
				movableLine[x] = true
				// fmt.Printf("Found movable roll at %d, %d\n", x, y)
			}
		}
		movableMap = append(movableMap, movableLine)
	}
	printMap(rollsMap, movableMap)
	return movable
}
func printMap(rollsMap [][]bool, movableMap [][]bool) {
	for y, line := range rollsMap {
		for x, e := range line {
			if e && movableMap[y][x] {
				// movable roll
				aoc.ColorPrint(aoc.ConstantToAnsiEscapeString(aoc.ANSI_BRIGHT_GREEN_FG), "@ ")
			} else if e {
				// not movable roll
				aoc.ColorPrint(aoc.ConstantToAnsiEscapeString(aoc.ANSI_BRIGHT_RED_FG), "@ ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
}

func handleLines(lines []string) int {
	paperRolls := make([][]bool, 0, len(lines))
	for _, line := range lines {
		paperRollLine := make([]bool, 0, len(line))
		for _, c := range line {
			paperRollLine = append(paperRollLine, c == '@')
		}
		paperRolls = append(paperRolls, paperRollLine)
	}
	return countAccess(paperRolls)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 13, "Example solution wrong")
	}
}
