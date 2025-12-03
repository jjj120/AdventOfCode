package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 03
const selectExample = false

func maxInSlice(s []int) (int, int) {
	if len(s) <= 0 {
		fmt.Printf("%v\n", s)
		panic("Slice is empty!")
	}
	maxInd := 0
	maxVal := s[0]

	for i, v := range s {
		if v > maxVal {
			maxVal = v
			maxInd = i
		}
	}
	return maxVal, maxInd
}

func getMaxJoltage(batteriesStr string, numTurnOn int) int {
	batteriesSplit := strings.Split(batteriesStr, "")
	batteries := make([]int, 0, len(batteriesStr))

	for _, bat := range batteriesSplit {
		batInt, err := strconv.Atoi(bat)
		aoc.Check(err)
		batteries = append(batteries, batInt)
	}

	maxVal, maxInd := maxInSlice(batteries[0 : len(batteries)-numTurnOn+1])
	fmt.Printf("Got maxVal %d at %d in %v for index %d\n", maxVal, maxInd, batteries, 0)

	maxValues := make([]int, numTurnOn, numTurnOn)
	maxValues[0] = maxVal

	for maxValueIndex := 1; maxValueIndex < numTurnOn; maxValueIndex++ {
		maxVal, newMaxInd := maxInSlice(batteries[maxInd+1 : len(batteries)-numTurnOn+maxValueIndex+1])
		maxInd += newMaxInd + 1

		fmt.Printf("Got maxVal %d at %d in %v for index %d\n", maxVal, maxInd, batteries, maxValueIndex)
		maxValues[maxValueIndex] = maxVal
	}

	maxValue := 0
	for i, v := range maxValues {
		maxValue += int(math.Pow(10., float64(numTurnOn-i-1))) * v
	}

	fmt.Printf("Got max %d in %v\n", maxValue, batteries)

	return maxValue
}

func handleLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getMaxJoltage(line, 12)
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 3121910778619, "Example is wrong")
	}
}
