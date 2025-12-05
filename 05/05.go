package main

import (
	"fmt"
	"strconv"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 05
const selectExample = false

type dataRange struct {
	low  int
	high int
}

func (d *dataRange) includes(v int) bool {
	return v >= d.low && v <= d.high
}

func countFresh(ranges []dataRange, ingredients []int) int {
	fresh := 0
	for _, i := range ingredients {
		for _, currRange := range ranges {
			if currRange.includes(i) {
				fresh++
				break
			}
		}
	}
	return fresh
}

func handleLines(lines []string) int {
	ranges := make([]dataRange, 0, 20)
	isRanges := true
	ingredients := make([]int, 0, 20)

	for _, line := range lines {
		if len(line) == 0 {
			isRanges = false
			continue
		}

		if isRanges {
			newRange := dataRange{0, 0}
			_, err := fmt.Sscanf(line, "%d-%d", &newRange.low, &newRange.high)
			aoc.Check(err)
			ranges = append(ranges, newRange)
		} else {
			ingredient, err := strconv.Atoi(line)
			aoc.Check(err)
			ingredients = append(ingredients, ingredient)
		}
	}

	fmt.Printf("Ranges: %v\n", ranges)
	fmt.Printf("Ingredients: %v\n", ingredients)

	return countFresh(ranges, ingredients)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 3, "Example solution is wrong")
	}
}
