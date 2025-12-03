package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 03
const selectExample = false

func getMaxJoltage(batteriesStr string) int {
	batteriesSplit := strings.Split(batteriesStr, "")
	batteries := make([]int, 0, len(batteriesStr))

	maxBat := -1
	maxInd := 0
	for i, bat := range batteriesSplit {
		batInt, err := strconv.Atoi(bat)
		aoc.Check(err)
		batteries = append(batteries, batInt)

		if batInt > maxBat && i != len(batteriesStr)-1 {
			// dont use the last entry
			maxBat = batInt
			maxInd = i
		}
	}

	secondMax := 0
	for i := maxInd + 1; i < len(batteries); i++ {
		secondMax = max(batteries[i], secondMax)
	}
	fmt.Printf("Got max1 %d at %d and max2 %d in %v\n", maxBat, maxInd, secondMax, batteries)

	return maxBat*10 + secondMax
}

func handleLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getMaxJoltage(line)
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 357, "Example is wrong")
	}
}
