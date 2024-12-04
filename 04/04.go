package main

import (
	"bufio"
	"fmt"
	"os"
)

const XMAS = "XMAS"
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

func countAtPosition(searchArray []string, x, y int, searchTerms []string) int {
	// always search to the right, down, and diagonally to the bottom and top right
	// always search for both search strings
	count := 0

	for _, searchTerm := range searchTerms {
		if searchArray[y][x] != searchTerm[0] {
			continue
		}

		for _, direction := range SEARCH_DIRECTIONS {
			dx := direction[0]
			dy := direction[1]
			if searchInDirection(searchArray, x, y, dx, dy, searchTerm) {
				count++
			}
		}
	}

	return count
}

func searchInDirection(searchArray []string, x, y, dx, dy int, searchTerm string) bool {
	if x+dx*len(searchTerm) > len(searchArray[y]) || y+dy*len(searchTerm) > len(searchArray) || y+dy*len(searchTerm) < -1 || x+dx*len(searchTerm) < -1 {
		return false
	}
	searchLen := len(searchTerm)
	for i := 0; i < searchLen; i++ {
		if searchArray[y+dy*i][x+dx*i] != searchTerm[i] {
			return false
		}
	}
	return true
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
