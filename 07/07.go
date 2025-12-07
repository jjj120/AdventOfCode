package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 07
const selectExample = false

func runBeamRec(lines []string, start aoc.Vec2d, cache *map[aoc.Vec2d]int) int {
	currEntry := start
	if val, ok := (*cache)[currEntry]; ok {
		return val
	}

	if lines[currEntry.Y][currEntry.X] == 'S' || lines[currEntry.Y][currEntry.X] == '.' {
		// empty space --> just move down
		currEntry.Y++
		if currEntry.Y < len(lines) {
			currTimelines := runBeamRec(lines, currEntry, cache)

			currEntry.Y--
			(*cache)[currEntry] = currTimelines
			return currTimelines
		}
	} else if lines[currEntry.Y][currEntry.X] == '^' {
		// splitter --> add new entry at x+1 and x-1
		currTimelinesSum := 0
		if currEntry.X+1 < len(lines[0]) {
			currTimelines := runBeamRec(lines, aoc.Vec2d{X: currEntry.X + 1, Y: currEntry.Y}, cache)
			(*cache)[aoc.Vec2d{X: currEntry.X + 1, Y: currEntry.Y}] = currTimelines
			currTimelinesSum += currTimelines
		}
		if currEntry.X-1 < len(lines[0]) {
			currTimelines := runBeamRec(lines, aoc.Vec2d{X: currEntry.X - 1, Y: currEntry.Y}, cache)
			(*cache)[aoc.Vec2d{X: currEntry.X - 1, Y: currEntry.Y}] = currTimelines
			currTimelinesSum += currTimelines
		}

		(*cache)[currEntry] = currTimelinesSum + 1
		return currTimelinesSum + 1
	} else {
		aoc.Assert(false, "Got unexpected entry in lines")
	}
	return (*cache)[currEntry]
}

func runBeam(lines []string, start aoc.Vec2d) int {
	queue := []aoc.Vec2d{start}
	splits := 1
	visited := make(map[aoc.Vec2d]bool)

	for len(queue) > 0 {
		currEntry := queue[0]
		queue = queue[1:]
		if _, ok := visited[currEntry]; ok {
			continue
		}
		visited[currEntry] = true

		fmt.Printf("Currently at %v\n", currEntry)

		if lines[currEntry.Y][currEntry.X] == 'S' || lines[currEntry.Y][currEntry.X] == '.' {
			// empty space --> just move down
			currEntry.Y++
			if currEntry.Y < len(lines) {
				queue = append(queue, currEntry)
			}
		} else if lines[currEntry.Y][currEntry.X] == '^' {
			// splitter --> add new entry at x+1 and x-1
			if currEntry.X+1 < len(lines[0]) {
				queue = append(queue, aoc.Vec2d{X: currEntry.X + 1, Y: currEntry.Y})
			}
			if currEntry.X-1 < len(lines[0]) {
				queue = append(queue, aoc.Vec2d{X: currEntry.X - 1, Y: currEntry.Y})
			}
			splits++
		} else {
			aoc.Assert(false, "Got unexpected entry in lines")
		}
	}

	return splits
}

func handleLines(lines []string) int {
	var start aoc.Vec2d
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				start = aoc.Vec2d{X: x, Y: y}
			}
		}
	}

	cache := make(map[aoc.Vec2d]int)
	return runBeamRec(lines, start, &cache) + 1
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 40, "Example wrong!")
	}
}
