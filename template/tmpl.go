package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 0
const selectExample = false

func handleLines(lines []string) int {
	for _, line := range lines {
		fmt.Print(line)
	}
	return 0
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
}
