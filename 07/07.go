package main

import (
	"fmt"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 07
const selectExample = false

func parseParts(s string) []string {
	matches := strings.FieldsFunc(s, func(r rune) bool { return r == '[' || r == ']' })
	return matches
}

func check4(part string) bool {
	return part[0] == part[3] && part[1] == part[2] && part[0] != part[1]
}

func checkABBA(part string) bool {
	for i := 0; i < len(part)-3; i++ {
		if check4(part[i : i+4]) {
			return true
		}
	}
	return false
}

func handleLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := parseParts(line)
		correct := false
		for i, part := range parts {
			if checkABBA(part) {
				if i%2 == 0 {
					// outside
					correct = true
				} else {
					// inside
					correct = false
					break
				}
			}
		}
		if correct {
			// fmt.Println(line)
			sum += 1
		}
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 2, "Example wrong!")
	}
}
