package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 01
const selectExample = false

type dial struct {
	number int
	limit  int
}

func (d *dial) turn(dir rune, num int) {
	if dir == 'L' {
		num = -num
	}

	d.number += num
	d.number %= d.limit
}

func handleLines(lines []string) int {
	d := dial{50, 100}
	timesAtZero := 0
	for _, line := range lines {
		var dir rune
		num := 0
		_, err := fmt.Sscanf(line, "%c%d", &dir, &num)
		aoc.Check(err)

		d.turn(dir, num)
		if d.number == 0 {
			timesAtZero++
		}
	}
	return timesAtZero
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
}
