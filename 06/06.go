package main

import (
	"fmt"
	"math"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 06
const selectExample = false

func calcSolution(inputs [][]int, ops []string) int {
	totalSum := 0
	for i, op := range ops {
		sum := 0
		if op == "+" {
			for _, elem := range inputs[i] {
				sum += elem
			}
		} else if op == "*" {
			sum = 1
			for _, elem := range inputs[i] {
				sum *= elem
			}
		} else {
			aoc.Assert(1 == 0, "Got unknown op!")
		}
		if selectExample {
			fmt.Println(sum)
		}
		totalSum += sum
	}
	return totalSum
}

func powInt(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func handleLines(lines []string) int {
	ops := strings.Fields(lines[len(lines)-1])
	numCols := len(strings.Fields(lines[0]))

	inputs := make([][]int, numCols)
	maxLen := 0
	for i := range inputs {
		inputs[i] = make([]int, 0, 10)
	}
	for _, line := range lines {
		maxLen = max(maxLen, len(line))
	}

	colIndex := 0

	for i := range maxLen {
		currNumbers := make([]int, 0, 5)
		for _, line := range lines[:len(lines)-1] {
			if i < len(line) && line[i] != ' ' {
				currNumbers = append(currNumbers, int(line[i]-'0'))
			}
		}
		if len(currNumbers) > 0 {
			// there was a number
			currNumber := 0
			for digInd, num := range currNumbers {
				currNumber += num * powInt(10, len(currNumbers)-digInd-1)
			}

			inputs[colIndex] = append(inputs[colIndex], currNumber)
		} else {
			colIndex++
		}
	}

	if selectExample {
		fmt.Printf("%v\n", inputs)
		fmt.Printf("%v\n", ops)
	}

	return calcSolution(inputs, ops)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 3263827, "Example is wrong!")
	}
}
