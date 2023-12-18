package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type digInstruction struct {
	dir   int
	len   int
	color int
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func parseInstruction(line string) digInstruction {
	splitString := strings.Split(line, " ")
	var instr digInstruction
	switch line[0] {
	case 'U':
		instr.dir = UP
	case 'R':
		instr.dir = RIGHT
	case 'D':
		instr.dir = DOWN
	case 'L':
		instr.dir = LEFT
	}

	num, err := strconv.ParseInt(splitString[1], 10, 0)
	check(err)
	instr.len = int(num)

	splitString[2] = strings.ReplaceAll(splitString[2], "(", "")
	splitString[2] = strings.ReplaceAll(splitString[2], ")", "")
	splitString[2] = strings.ReplaceAll(splitString[2], "#", "")

	num, err = strconv.ParseInt(splitString[2], 16, 0)
	check(err)

	instr.color = int(num)

	return instr
}

func handleLine(digPlan [][]bool, line string, currX, currY *int) {
	inst := parseInstruction(line)

	for i := 0; i < inst.len; i++ {
		switch inst.dir {
		case UP:
			*currY--
		case RIGHT:
			*currX++
		case DOWN:
			*currY++
		case LEFT:
			*currX--
		}
		if *currX < 0 || *currY < 0 || *currX >= len(digPlan) || *currY >= len(digPlan) {
			fmt.Printf("Out of bounds: %d, %d\n", *currX, *currY)
			os.Exit(1)
		}
		fmt.Printf("X: %d, Y: %d\n", *currX, *currY)
		digPlan[*currY][*currX] = true
	}
}

func printDigPlan(digPlan [][]bool) {
	for _, row := range digPlan {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func floodFill(digPlan [][]bool, x, y int) {
	if x < 0 || y < 0 || x >= len(digPlan) || y >= len(digPlan) {
		return
	}

	if !digPlan[y][x] {
		digPlan[y][x] = true
		floodFill(digPlan, x-1, y)
		floodFill(digPlan, x+1, y)
		floodFill(digPlan, x, y-1)
		floodFill(digPlan, x, y+1)
		floodFill(digPlan, x-1, y-1)
		floodFill(digPlan, x+1, y-1)
		floodFill(digPlan, x-1, y+1)
		floodFill(digPlan, x+1, y+1)
	}
}

func fillOutsideOfLine(digPlan [][]bool) {
	floodFill(digPlan, 0, 0)
}

func countIncluded(digPlan [][]bool) int {
	count := 0

	for _, row := range digPlan {
		for _, col := range row {
			if col {
				count++
			}
		}
	}

	return count
}

func initDigPlan(size int) [][]bool {
	digPlan := make([][]bool, size)
	for i := 0; i < size; i++ {
		digPlan[i] = make([]bool, size)
	}
	return digPlan
}

func invertDigPlan(digPlan [][]bool) {
	for i, row := range digPlan {
		for j, col := range row {
			digPlan[i][j] = !col
		}
	}
}

func main() {
	// Open the file
	file, err := os.Open("18.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	size := 1000
	digPlan := initDigPlan(size)
	currX, currY := size/2, size/2
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(digPlan, line, &currX, &currY)
	}

	// printDigPlan(digPlan)
	fmt.Println()
	sum = countIncluded(digPlan)

	fillOutsideOfLine(digPlan)
	invertDigPlan(digPlan)

	sum += countIncluded(digPlan)

	// printDigPlan(digPlan)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
