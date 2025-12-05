package main

import (
	"errors"
	"fmt"
	"strconv"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 05
const selectExample = false

type dataRange struct {
	low   int
	high  int
	empty bool
}

func (d *dataRange) width() int {
	if !d.empty {
		return d.high - d.low + 1
	}
	return 0
}

func (d *dataRange) includes(v int) bool {
	return v >= d.low && v <= d.high && !d.empty
}

func (d *dataRange) overlaps(nd dataRange) bool {
	return (d.includes(nd.low) || d.includes(nd.high)) && !d.empty && !nd.empty
}

func (d *dataRange) mergeInPlace(nd dataRange) error {
	if !d.overlaps(nd) {
		return errors.New("Merge of non overlapping ranges")
	}
	if nd.empty {
		return nil
	}
	if d.empty {
		d.empty = false
		d.low = nd.low
		d.high = nd.high
		return nil
	}

	d.low = min(d.low, nd.low)
	d.high = max(d.high, nd.high)
	return nil
}

func hasOverlaps(ranges []dataRange) bool {
	for i, r1 := range ranges {
		for j, r2 := range ranges {
			// merge ranges together
			if i != j && r1.overlaps(r2) {
				return true
			}
		}
	}
	return false
}

func countFreshRange(ranges []dataRange, ingredients []int) int {
	for hasOverlaps(ranges) {
		for i, r1 := range ranges {
			for j, r2 := range ranges {
				// merge ranges together
				if i != j && r1.overlaps(r2) {
					ranges[i].mergeInPlace(r2)
					ranges[j].empty = true
				}
			}
		}
	}

	aoc.Assert(!hasOverlaps(ranges), "Still detected overlaps!")

	// sum up ranges
	sum := 0
	nonempty := 0
	for _, r := range ranges {
		sum += r.width()
		if !r.empty {
			nonempty++
		}
	}
	fmt.Printf("Found %d non-overlapping ranges\n", nonempty)
	return sum
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
			break
		}

		if isRanges {
			newRange := dataRange{0, 0, false}
			_, err := fmt.Sscanf(line, "%d-%d", &newRange.low, &newRange.high)
			aoc.Check(err)
			ranges = append(ranges, newRange)
		} else {
			ingredient, err := strconv.Atoi(line)
			aoc.Check(err)
			ingredients = append(ingredients, ingredient)
		}
	}

	// fmt.Printf("Ranges: %v\n", ranges)
	// fmt.Printf("Ingredients: %v\n", ingredients)

	return countFreshRange(ranges, ingredients)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 14, "Example solution is wrong")
	} else {
		aoc.Assert(sum <= 353600416748441, "Real solution is too high")
	}
}
