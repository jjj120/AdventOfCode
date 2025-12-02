package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 02
const selectExample = false

func checkValid(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	firstHalf := id[:len(id)/2]
	secondHalf := id[len(id)/2:]

	return firstHalf == secondHalf
}

func sumInvalidIDs(start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		if checkValid(strconv.Itoa(i)) {
			sum += i
		}
	}
	return sum
}

func handleRanges(ranges []string) int {
	sum := 0
	for _, r := range ranges {
		var start, end int
		fmt.Sscanf(r, "%d-%d", &start, &end)

		sum += sumInvalidIDs(start, end)
	}
	return sum
}

func handleLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += handleRanges(strings.Split(line, ","))
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 1227775554, "example solution wrong")
	}
}
