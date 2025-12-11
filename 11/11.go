package main

import (
	"fmt"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 11
const selectExample = false

func countPaths(devices map[string][]string, from, to string) int {
	sum := 0
	for _, d := range devices[from] {
		if d == to {
			sum++
		} else {
			sum += countPaths(devices, d, to)
		}
	}
	return sum
}

func handleLines(lines []string) int {
	devices := make(map[string][]string)
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		deviceName := splitLine[0]
		splitLine = strings.Split(splitLine[1], " ")
		devices[deviceName] = splitLine
	}

	return countPaths(devices, "you", "out")
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 5, "Example wrong!")
	}
}
