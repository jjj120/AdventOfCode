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

func (d *dial) turn(dir rune, num int) int {
	// returns if the dial passed zero
	fmt.Printf("Turn %c from %d by %d", dir, d.number, num)

	zeroCount := 0
	for range num {
		if dir == 'L' {
			d.number--
		} else {
			d.number++
		}

		d.number += d.limit
		d.number %= d.limit

		if d.number == 0 {
			zeroCount++
		}
	}

	/* //This did not work, dont know why
	if dir == 'L' {
		num = -num
	}
	zeroCount := 0

	d.number += num

	for d.number < 0 {
		// turned left over zero
		d.number += d.limit
		zeroCount++
	}

	if d.number == 0 {
		zeroCount++
	}

	for d.number >= d.limit {
		// turned right over zero
		d.number -= d.limit
		zeroCount++
	}
	*/

	aoc.Assert(d.number >= 0, "Number is less than zero")
	aoc.Assert(d.number < d.limit, "Number is more than limit")

	fmt.Printf(" to %d, hit zero %d times\n", d.number, zeroCount)

	return zeroCount
}

func intMax(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func checkFunction() {
	for i := 1; i < 550; i++ {
		var expectedL int
		var expectedR int

		switch {
		case i < 50:
			expectedL = 0
			expectedR = 0
		case i < 150:
			expectedL = 1
			expectedR = 1
		case i < 250:
			expectedL = 2
			expectedR = 2
		case i < 350:
			expectedL = 3
			expectedR = 3
		case i < 450:
			expectedL = 4
			expectedR = 4
		case i < 550:
			expectedL = 5
			expectedR = 5

		default:
			aoc.Assert(false, "Ran into default in testing")
		}

		d := dial{50, 100}
		num := d.turn('L', i)
		aoc.Assertf(num == expectedL, "Error at %c%d. Got %d, expected %d.", 'L', i, num, expectedL)

		d = dial{50, 100}
		num = d.turn('R', i)
		aoc.Assertf(num == expectedL, "Error at %c%d. Got %d, expected %d.", 'R', i, num, expectedR)
	}

	d := dial{0, 100}
	aoc.Assert(d.turn('L', 100) == 1, "Failed special L 100")
	aoc.Assert(d.turn('R', 100) == 1, "Failed special R 100")
	aoc.Assert(d.turn('L', 200) == 2, "Failed special L 200")
	aoc.Assert(d.turn('R', 200) == 2, "Failed special R 200")
}

func handleLines(lines []string) int {
	d := dial{50, 100}
	timesAtZero := 0
	for _, line := range lines {
		var dir rune
		num := 0
		_, err := fmt.Sscanf(line, "%c%d", &dir, &num)
		aoc.Check(err)

		timesAtZero += d.turn(dir, num)
	}
	return timesAtZero
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	checkFunction()

	sum := handleLines(lines)

	aoc.Assert(sum != 5900, "5900 was a prev answer!")
	aoc.Assert(sum != 6815, "6815 was a prev answer!")
	aoc.Assert(sum != 6402, "6402 was a prev answer!")
	fmt.Printf("Sum: %d\n", sum)
}
