package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 02
const selectExample = false

func checkInvalid(id string) bool {
	for l := 1; l <= len(id)/2; l++ {
		if len(id)%l != 0 {
			continue
		}

		numElements := len(id) / l
		prevElement := id[:l]

		equal := true
		for i := range numElements {
			if prevElement != id[i*l:(i+1)*l] {
				equal = false
				break
			}
		}
		if equal {
			return true
		}
	}
	return false
}

func sumInvalidIDs(start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		if checkInvalid(strconv.Itoa(i)) {
			fmt.Printf("Found invalid id %d\n", i)
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
		aoc.Assert(sum == 4174379265, "example solution wrong")
	}
}
