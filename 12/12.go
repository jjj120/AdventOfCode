package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 12
const selectExample = false

type Present struct {
}

func checkFit(presents []Present, width int, height int, nums []int) int {
	sizePresents := 0
	for _, num := range nums {
		sizePresents += num
	}

	if width*height >= sizePresents*9 {
		return 1
	}
	return 0
}

func handleLines(lines []string) int {
	isPresent := true
	presents := make([]Present, 0)
	sum := 0
	for _, line := range lines {
		if isPresent && len(line) >= 3 && line[2] == 'x' {
			isPresent = false
		}

		if !isPresent {
			parts := strings.Split(line, ": ")

			width := 0
			height := 0
			fmt.Sscanf(parts[0], "%dx%d", &width, &height)

			nums := make([]int, 0, 10)
			for _, numS := range strings.Split(parts[1], " ") {
				num, err := strconv.Atoi(numS)
				aoc.Check(err)

				nums = append(nums, num)
			}

			sum += checkFit(presents, width, height, nums)
		}
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
}
