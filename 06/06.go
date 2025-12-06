package main

import (
	"fmt"
	"strconv"
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
			for _, line := range inputs {
				sum += line[i]
			}
		} else if op == "*" {
			sum = 1
			for _, line := range inputs {
				sum *= line[i]
			}
		} else {
			aoc.Assert(1 == 0, "Got unknown op!")
		}
		totalSum += sum
	}
	return totalSum
}

func handleLines(lines []string) int {
	inputs := make([][]int, 0, 10)
	ops := make([]string, 0, 10)
	for _, line := range lines {
		if line[0] != '+' && line[0] != '*' {
			// number line
			splitLines := strings.Fields(line)
			intLines := make([]int, len(splitLines))
			for i, s := range splitLines {
				parsed, err := strconv.Atoi(s)
				aoc.Check(err)
				intLines[i] = parsed
			}
			inputs = append(inputs, intLines)
		} else {
			// op line
			splitLines := strings.Fields(line)
			for _, s := range splitLines {
				ops = append(ops, s)
			}
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
		aoc.Assert(sum == 4277556, "Example is wrong!")
	}
}
