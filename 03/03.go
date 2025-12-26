package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 03
const selectExample = false

func checkTriangle(l1, l2, l3 int) bool {
	return (l1+l2 > l3 && l2+l3 > l1 && l3+l1 > l2)
}

func handleLines(lines []string) int {
	counter := 0
	prev_0 := []int{0}
	prev_1 := []int{0}
	for i, line := range lines {
		s := strings.Split(line, " ")
		nums := make([]int, 0, 3)
		for i, n := range s {
			s[i] = strings.TrimSpace(n)
			if len(s[i]) > 0 {
				num, err := strconv.Atoi(s[i])
				aoc.Check(err)
				nums = append(nums, num)
			}
		}

		if i%3 == 0 {
			prev_0 = nums
		} else if i%3 == 1 {
			prev_1 = nums
		} else {
			if checkTriangle(prev_0[0], prev_1[0], nums[0]) {
				counter += 1
			}
			if checkTriangle(prev_0[1], prev_1[1], nums[1]) {
				counter += 1
			}
			if checkTriangle(prev_0[2], prev_1[2], nums[2]) {
				counter += 1
			}
		}
	}
	return counter
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 6, "Example wrong!")
	}
}
