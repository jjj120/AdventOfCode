package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coord struct {
	x, y int
}

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func assertf(condition bool, message string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(message, args...))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type NumericalKeypad struct {
	curr rune
}

type DirectionKeypad struct {
	curr rune
}

func handleLine(line string) int {
	numericalValue, err := strconv.Atoi(line[:len(line)-1])
	check(err)
	// fmt.Printf("Numerical value: %d\n", numericalValue)

	code := line

	numKeypad := NumericalKeypad{'A'}
	dirKeypad1 := DirectionKeypad{'A'}
	dirKeypad2 := DirectionKeypad{'A'}

	code = numKeypad.expandCodeNumericalKeypad(code)
	// fmt.Printf("%s 1: %d %s\n", line, len(code), code)
	code = dirKeypad1.expandCodeDirectionKeypad(code)
	// fmt.Printf("%s 2: %d %s\n", line, len(code), code)
	code = dirKeypad2.expandCodeDirectionKeypad(code)
	// fmt.Printf("%s 3: %d %s\n", line, len(code), code)

	fmt.Printf("%s: %d %s\n", line, len(code), code)

	return len(code) * numericalValue
}

func (k *NumericalKeypad) expandCodeNumericalKeypad(code string) string {
	// Keypad matrix:
	// 7 8 9
	// 4 5 6
	// 1 2 3
	// x 0 A
	// starting at A

	expanded := ""
	for i := 0; i < len(code); i++ {
		from := k.curr
		to := rune(code[i])
		expanded += numericalFromTo(from, to)
		expanded += "A"
		k.curr = to
	}

	return expanded
}

func numericalFromTo(from, to rune) string {
	coords := map[rune]Coord{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		'0': {1, 3},
		'A': {2, 3},
	}

	fromCoord := coords[from]
	toCoord := coords[to]

	xDiff := toCoord.x - fromCoord.x
	yDiff := toCoord.y - fromCoord.y

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = toCoord.x - fromCoord.x
	yDiff = toCoord.y - fromCoord.y

	// Priority: < over ^ over v over >

	if fromCoord.y == 3 && toCoord.x == 0 {
		return vertical + horizontal
	} else if fromCoord.x == 0 && toCoord.y == 3 {
		return horizontal + vertical
	} else if xDiff < 0 {
		return horizontal + vertical
	} else if xDiff >= 0 {
		return vertical + horizontal
	}

	panic("Invalid direction")
}

func (k *DirectionKeypad) expandCodeDirectionKeypad(code string) string {
	// Keypad matrix:
	// x ^ A
	// < v >

	expanded := ""
	for i := 0; i < len(code); i++ {
		from := k.curr
		to := rune(code[i])
		expanded += directionFromTo(from, to)
		expanded += "A"
		k.curr = to
	}

	return expanded
}

func directionFromTo(from, to rune) string {
	coords := map[rune]Coord{
		'^': {1, 0},
		'A': {2, 0},
		'<': {0, 1},
		'v': {1, 1},
		'>': {2, 1},
	}

	fromCoord := coords[from]
	toCoord := coords[to]

	xDiff := toCoord.x - fromCoord.x
	yDiff := toCoord.y - fromCoord.y

	vertical := ""
	for yDiff < 0 {
		vertical += "^"
		yDiff++
	}
	for yDiff > 0 {
		vertical += "v"
		yDiff--
	}

	horizontal := ""
	for xDiff < 0 {
		horizontal += "<"
		xDiff++
	}
	for xDiff > 0 {
		horizontal += ">"
		xDiff--
	}

	xDiff = toCoord.x - fromCoord.x
	yDiff = toCoord.y - fromCoord.y

	// Priority: < over ^ over v over >
	if fromCoord.y == 1 && toCoord.y == 1 {
		return horizontal + vertical
	} else if fromCoord.y == 0 && toCoord.x == 0 {
		return vertical + horizontal
	} else if xDiff < 0 {
		return horizontal + vertical
	} else if xDiff >= 0 {
		return vertical + horizontal
	}

	panic("Invalid direction")
}

func main() {
	// Open the file
	file, err := os.Open("21.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		sum += handleLine(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
	assert(sum == 213536, "Sum wrong")
}
