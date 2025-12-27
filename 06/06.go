package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 06
const selectExample = false

func handleLines(lines []string) string {
	maps := make([]map[rune]int, len(lines[0]))
	for i := range maps {
		maps[i] = make(map[rune]int)
		for r := 'a'; r <= 'z'; r++ {
			maps[i][r] = 0
		}
	}

	for _, line := range lines {
		for i, r := range line {
			maps[i][r] += 1
		}
	}

	sol := ""
	for _, m := range maps {
		minChar := 'a'
		minNum := 1000000
		for k, v := range m {
			if minNum > v && v != 0 {
				minNum = v
				minChar = k
			}
		}
		sol += string(minChar)
	}

	return sol
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %s\n", sum)
	if selectExample {
		aoc.Assert(sum == "advent", "Example wrong!")
	}
}
