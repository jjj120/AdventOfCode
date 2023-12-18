package main

import (
	"bufio"
	"fmt"
	"math"
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
	dir int
	len int
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func parseInstruction(line string) digInstruction {
	splitString := strings.Split(line, " ")
	splitString[2] = strings.ReplaceAll(splitString[2], "(", "")
	splitString[2] = strings.ReplaceAll(splitString[2], ")", "")
	splitString[2] = strings.ReplaceAll(splitString[2], "#", "")

	var instr digInstruction
	switch splitString[2][5] - '0' {
	case 0:
		instr.dir = RIGHT
	case 1:
		instr.dir = DOWN
	case 2:
		instr.dir = LEFT
	case 3:
		instr.dir = UP
	}

	num, err := strconv.ParseInt(splitString[2][0:5], 16, 0)
	check(err)
	instr.len = int(num)

	return instr
}

func addToVertices(line string, vertices *[][2]int, x, y *int) {
	instr := parseInstruction(line)

	switch instr.dir {
	case UP:
		*y -= instr.len
	case RIGHT:
		*x += instr.len
	case DOWN:
		*y += instr.len
	case LEFT:
		*x -= instr.len
	}

	p2 := [2]int{*x, *y}
	*vertices = append(*vertices, p2)
}

func calcArea(vertices [][2]int) int {
	if len(vertices)%2 != 0 {
		vertices = append(vertices, vertices[0])
	}
	var area int
	for i := 0; i < len(vertices); i += 2 {
		area += vertices[i][0]*vertices[i+1][1] - vertices[i+1][0]*vertices[i][1]
	}
	return area
}

func calcPerimeter(vertices [][2]int) int {
	var perimeter int
	for i := 0; i < len(vertices)-1; i++ {
		perimeter += int(math.Sqrt(math.Pow(float64(vertices[i][0]-vertices[i+1][0]), 2) + math.Pow(float64(vertices[i][1]-vertices[i+1][1]), 2)))
	}
	return perimeter
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
	currX, currY := 0, 0
	// Iterate through each line

	vertices := make([][2]int, 0)
	vertices = append(vertices, [2]int{0, 0})
	for scanner.Scan() {
		line := scanner.Text()
		addToVertices(line, &vertices, &currX, &currY)
	}

	sum = calcArea(vertices)           // area of polygon
	sum += calcPerimeter(vertices) / 2 // add half of perimeter for the fields on the side, that the edge runs through in the middle
	sum++                              // add 1 for the one field outside the perimeter

	fmt.Printf("Area: %d, Perimeter: %d\n", calcArea(vertices), calcPerimeter(vertices))

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
