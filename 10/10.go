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

func findStart(data []string) (int, int) {
	for y, line := range data {
		for x, char := range line {
			if char == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func getDir(data []string, coord []int) byte {
	return data[coord[1]][coord[0]]
}

func checkBounds(data []string, coord []int) bool {
	return coord[0] >= 0 && coord[1] >= 0 && coord[0] < len(data[coord[1]]) && coord[1] < len(data)
}

func getStartContinue(data []string, startCoord []int) []int {
	currCoord := []int{startCoord[0], startCoord[1]}

	currCoord[0] += 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '-' {
		return currCoord
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[0] -= 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '-' {
		return currCoord
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[1] += 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '|' {
		return currCoord
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[1] -= 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '|' {
		return currCoord
	}
	return []int{startCoord[0], startCoord[1]}
}

func nextStep(data []string, currCoord []int, prevCoord []int) []int {
	toTry := []int{currCoord[0], currCoord[1]}

	switch getDir(data, currCoord) {
	case 'S':
		return getStartContinue(data, currCoord)
	case '|':
		toTry[1] += 1
	case '-':
		toTry[0] += 1
	case 'F':
		toTry[0] += 1
	case 'J':
		toTry[0] -= 1
	case 'L':
		toTry[0] += 1
	case '7':
		toTry[0] -= 1
	}

	if checkBounds(data, toTry) && (toTry[0] != prevCoord[0] || toTry[1] != prevCoord[1]) {
		return toTry
	}

	copy(toTry, currCoord)
	switch getDir(data, currCoord) {
	case 'S':
		return getStartContinue(data, currCoord)
	case '|':
		toTry[1] -= 1
	case '-':
		toTry[0] -= 1
	case 'F':
		toTry[1] += 1
	case 'J':
		toTry[1] -= 1
	case 'L':
		toTry[1] -= 1
	case '7':
		toTry[1] += 1
	}

	if checkBounds(data, toTry) && (toTry[0] != prevCoord[0] || toTry[1] != prevCoord[1]) {
		return toTry
	}

	fmt.Printf("Error on %d %d with prev %d %d\n", currCoord[0], currCoord[1], prevCoord[0], prevCoord[1])
	return currCoord
}

func getLoopLength(startX int, startY int, data []string) int {
	prev := []int{startX, startY}
	currCoord := nextStep(data, prev, prev)
	length := 1

	for getDir(data, currCoord) != 'S' {
		newCurrCoord := nextStep(data, currCoord, prev)
		prev = currCoord
		currCoord = newCurrCoord
		length++
		if length%100 == 0 {
			fmt.Printf("%d: curr: %d %d, prev: %d %d\n", length, currCoord[0], currCoord[1], prev[0], prev[1])
		}
	}
	return length
}

func handleData(data []string) int {
	startX, startY := findStart(data)
	fmt.Printf("Start: %d %d\n", startX, startY)

	return (getLoopLength(startX, startY, data)) / 2
}

func main() {
	// Open the file
	file, err := os.Open("10.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var data []string
	// Iterate through each line
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	sum := handleData(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
