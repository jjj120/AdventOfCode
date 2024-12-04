package main

import (
	"bufio"
	"fmt"
	"os"
)

const DEBUG = false

var SEARCH_DIRECTIONS = [][]int{
	{1, 0},
	{0, 1},
	{1, 1},
	{1, -1},
	{-1, 0},
	{0, -1},
	{-1, -1},
	{-1, 1},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printDebug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

func countOccurrences(searchArray []string) int {
	count := 0
	for y := 0; y < len(searchArray); y++ {
		for x := 0; x < len(searchArray[y]); x++ {
			// always search to the right, down, and diagonally to the bottom and top right
			// always search for both search strings
			count += countAtPosition(searchArray, x, y)
		}
	}
	return count
}

func countAtPosition(searchArray []string, x, y int) int {
	// always search to the right, down, and diagonally to the bottom and top right
	// always search for both search strings
	if searchArray[y][x] != 'A' {
		return 0
	}

	if y-1 < 0 || y+1 >= len(searchArray) || x-1 < 0 || x+1 >= len(searchArray[y]) {
		return 0
	}

	// M.M
	// .A.
	// S.S
	if searchArray[y-1][x-1] == 'M' && searchArray[y-1][x+1] == 'M' && searchArray[y+1][x-1] == 'S' && searchArray[y+1][x+1] == 'S' {
		return 1
	}

	// M.S
	// .A.
	// M.S
	if searchArray[y-1][x-1] == 'M' && searchArray[y-1][x+1] == 'S' && searchArray[y+1][x-1] == 'M' && searchArray[y+1][x+1] == 'S' {
		return 1
	}

	// S.M
	// .A.
	// S.M
	if searchArray[y-1][x-1] == 'S' && searchArray[y-1][x+1] == 'M' && searchArray[y+1][x-1] == 'S' && searchArray[y+1][x+1] == 'M' {
		return 1
	}

	// S.S
	// .A.
	// M.M
	if searchArray[y-1][x-1] == 'S' && searchArray[y-1][x+1] == 'S' && searchArray[y+1][x-1] == 'M' && searchArray[y+1][x+1] == 'M' {
		return 1
	}

	return 0
}

func main() {
	// Open the file
	file, err := os.Open("04.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var searchArray []string
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		searchArray = append(searchArray, line)
	}

	sum := countOccurrences(searchArray)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
