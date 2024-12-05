package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseOrderDuos(line string) (int, int) {
	// format: X|Y

	before := 0
	after := 0
	_, err := fmt.Sscanf(line, "%d|%d", &before, &after)
	check(err)
	return before, after
}

func parsePrintedPageNumbers(line string) []int {
	// separate by comma
	var pages []int
	for _, page := range strings.Split(line, ",") {
		p, err := strconv.Atoi(page)
		check(err)
		pages = append(pages, p)
	}
	return pages
}

func checkPages(pages []int, beforeMap map[int][]int) bool {
	for i := 0; i < len(pages); i++ {
		for _, before := range beforeMap[pages[i]] {
			if slices.Contains(pages[i+1:], before) {
				// check if some page that should be before is after
				return false
			}
		}
	}
	return true
}

func getMiddlePage(pages []int) int {
	middleIndex := len(pages) / 2
	return pages[middleIndex]
}

func printMap(m map[int][]int) {
	for key, value := range m {
		fmt.Printf("%d -> %v\n", key, value)
	}
}

func main() {
	// Open the file
	file, err := os.Open("05.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	beforeMap := make(map[int][]int)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		before, after := parseOrderDuos(line)
		beforeMap[after] = append(beforeMap[after], before)
	}

	// printMap(beforeMap)
	// fmt.Println("")

	var sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		pages := parsePrintedPageNumbers(line)

		if checkPages(pages, beforeMap) {
			// fmt.Printf("Valid: %v\n", pages)
			sum += getMiddlePage(pages)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
