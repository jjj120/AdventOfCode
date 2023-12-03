package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// returns if char is a symbol
func checkSymbol(char rune) bool {
	return !(unicode.IsDigit(char)) && (char != '.')
}

// returns if in the surroundings of a point is a symbol
func checkSurroundingsOfPoint(data []string, lineIndex int, index int) bool {
	dirs := [3]int{-1, 0, 1}

	for _, dirY := range dirs {
		newLineIndex := lineIndex + dirY
		if newLineIndex < 0 || newLineIndex >= len(data) {
			continue
		}

		for _, dirX := range dirs {
			newIndex := index + dirX
			// fmt.Printf("Checking %d %d\n", newLineIndex, newIndex)
			if newIndex < 0 || newIndex >= len(data[newLineIndex]) {
				continue
			}
			if checkSymbol([]rune(data[newLineIndex])[newIndex]) {
				return true
			}
		}
	}
	return false
}

// returns number if it is a valid part and 0 if not
func checkSurroundings(data []string, lineIndex int, firstIndex int, lastIndex int, numberStr string) int {
	for i := firstIndex; i < lastIndex; i++ {
		if checkSurroundingsOfPoint(data, lineIndex, i) {
			number, err := strconv.ParseInt(numberStr, 10, 0)
			check(err)
			return int(number)
		}
	}
	return 0
}

func handleLine(lineIndex int, line string, data []string) int {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)
	matches := re.FindAllStringIndex(line, -1)

	sum := 0

	fmt.Printf("Line %d\n", lineIndex)
	for i, num := range numbers {
		fmt.Printf("Found %s at %d:%d\n", num, matches[i][0], matches[i][1])
		sum += checkSurroundings(data, lineIndex, matches[i][0], matches[i][1], num)
	}
	return sum
}

func sumPartNumbers(data []string) int {
	sum := 0
	for lineIndex, line := range data {
		// loop through all strings and search for all numbers
		sum += handleLine(lineIndex, line, data)
	}
	return sum
}

func main() {
	// Open the file
	file, err := os.Open("3.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	data := make([]string, 0)
	// Iterate through each line
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	sum := sumPartNumbers(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
