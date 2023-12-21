package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) ([]bool, int) {
	var result []bool
	var startPoint = -1
	for x, c := range line {
		if c == '#' {
			result = append(result, false)
		} else if c == '.' {
			result = append(result, true)
		} else if c == 'S' {
			result = append(result, true)
			startPoint = x
		} else {
			panic("Unknown character")
		}
	}
	return result, startPoint
}

func findStart(lines [][]bool) [2]int {
	for y, line := range lines {
		for x, c := range line {
			if !c {
				return [2]int{x, y}
			}
		}
	}
	panic("No start found")
}

func contains2(history [][2]int, point [2]int) bool {
	for _, p := range history {
		if p == point {
			return true
		}
	}
	return false
}

func contains3(history [][3]int, point [3]int) bool {
	for _, p := range history {
		if p[0] == point[0] && p[1] == point[1] && p[2] == point[2] {
			return true
		}
	}
	return false
}

func countPossibleSteps(lines [][]bool, stepsLeft int, startPoint [2]int, history [][3]int, endPoints *[][2]int) int {
	stepsLeftOrig := stepsLeft
	queue := [][3]int{{startPoint[0], startPoint[1], stepsLeft}}

	var sum = 0
	for len(queue) > 0 {
		if len(history)%1000 == 0 {
			fmt.Printf("\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r%d: %d-%d, %d/%d", len(queue), queue[0][0], queue[0][1], len(history), len(lines)*len(lines[0])*stepsLeftOrig)
		}
		curr := queue[0]
		queue = queue[1:]

		if curr[2] == 0 {
			if !contains2(*endPoints, [2]int{curr[0], curr[1]}) {
				*endPoints = append(*endPoints, [2]int{curr[0], curr[1]})
			}
			sum++
			continue
		}

		dirs := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		for _, dir := range dirs {
			x := curr[0] + dir[0]
			y := curr[1] + dir[1]

			if x < 0 || y < 0 || x >= len(lines[0]) || y >= len(lines) {
				continue
			}
			if !lines[y][x] {
				continue
			}

			if contains3(history, [3]int{x, y, curr[2] - 1}) {
				continue
			}
			history = append(history, [3]int{x, y, curr[2] - 1})

			queue = append(queue, [3]int{x, y, curr[2] - 1})
		}
	}
	return sum
}

func removeDuplicates(endPoints [][2]int) [][2]int {
	var result [][2]int
	for _, p := range endPoints {
		if !contains2(result, p) {
			result = append(result, p)
		}
	}
	return result
}

func printGarden(lines [][]bool, endPoints [][2]int) {
	for y, line := range lines {
		for x, c := range line {
			if !c {
				fmt.Print("#")
			} else if contains2(endPoints, [2]int{x, y}) {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

const PRINT = false

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
	var lines [][]bool
	var startPointX = -1
	var startPointY = -1
	currLine := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		lineNew, start := handleLine(line)
		lines = append(lines, lineNew)

		if start != -1 {
			startPointX = start
			startPointY = currLine
		}

		currLine++
	}

	history := [][3]int{}
	endPoints := [][2]int{}
	sum = countPossibleSteps(lines, 64, [2]int{startPointX, startPointY}, history, &endPoints)
	endPoints = removeDuplicates(endPoints)
	fmt.Println()

	if PRINT {
		fmt.Printf("History: %d: %v\n", len(history), history)
		fmt.Printf("End points: %d: %v\n", len(endPoints), endPoints)

		printGarden(lines, endPoints)
	}

	// sum = len(endPoints)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
