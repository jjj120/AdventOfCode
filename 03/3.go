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

// // returns if char is a symbol
// func checkSymbol(char rune) bool {
// 	return !(unicode.IsDigit(char)) && (char != '.')
// }

// // returns if in the surroundings of a point is a symbol
// func checkSurroundingsOfPoint(data []string, lineIndex int, index int) bool {
// 	dirs := [3]int{-1, 0, 1}

// 	for _, dirY := range dirs {
// 		newLineIndex := lineIndex + dirY
// 		if newLineIndex < 0 || newLineIndex >= len(data) {
// 			continue
// 		}

// 		for _, dirX := range dirs {
// 			newIndex := index + dirX
// 			// fmt.Printf("Checking %d %d\n", newLineIndex, newIndex)
// 			if newIndex < 0 || newIndex >= len(data[newLineIndex]) {
// 				continue
// 			}
// 			if checkSymbol([]rune(data[newLineIndex])[newIndex]) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // returns number if it is a valid part and 0 if not
// func checkSurroundings(data []string, lineIndex int, firstIndex int, lastIndex int, numberStr string) int {
// 	for i := firstIndex; i < lastIndex; i++ {
// 		if checkSurroundingsOfPoint(data, lineIndex, i) {
// 			number, err := strconv.ParseInt(numberStr, 10, 0)
// 			check(err)
// 			return int(number)
// 		}
// 	}
// 	return 0
// }

func checkBounds(index int, dataLine string) bool {
	return (index >= 0 && index < len(dataLine))
}

func findNumber(data []string, lineIndex int, index int) int {
	startIndex := index
	endIndex := index

	for checkBounds(startIndex, data[lineIndex]) && unicode.IsDigit([]rune(data[lineIndex])[startIndex]) {
		startIndex--
	}
	startIndex++

	for checkBounds(endIndex, data[lineIndex]) && unicode.IsDigit([]rune(data[lineIndex])[endIndex]) {
		endIndex++
	}

	num, err := strconv.ParseInt(data[lineIndex][startIndex:endIndex], 10, 0)
	check(err)

	return int(num)
}

func findAdjacentNumbers(data []string, lineIndex int, star []int) []int {
	numbers := make([]int, 0)
	dirs := [3]int{-1, 0, 1}

	for _, dirY := range dirs {
		newLineIndex := lineIndex + dirY
		if newLineIndex < 0 || newLineIndex >= len(data) {
			continue
		}

		if unicode.IsDigit([]rune(data[newLineIndex])[star[0]]) {
			numbers = append(numbers, findNumber(data, newLineIndex, star[0]))
			continue
		}

		if unicode.IsDigit([]rune(data[newLineIndex])[star[0]-1]) {
			numbers = append(numbers, findNumber(data, newLineIndex, star[0]-1))
		}

		if unicode.IsDigit([]rune(data[newLineIndex])[star[0]+1]) {
			numbers = append(numbers, findNumber(data, newLineIndex, star[0]+1))
		}
	}

	return numbers
}

func checkIfGear(data []string, lineIndex int, stars [][]int) int {
	sum := 0
	for _, star := range stars {
		numbers := findAdjacentNumbers(data, lineIndex, star)
		// fmt.Print(numbers)
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}
	return sum
}

func findStars(data []string) [][][]int {
	matches := make([][][]int, 0)
	re := regexp.MustCompile(`(\*)+`)

	for _, line := range data {
		// loop through all strings and search for all numbers
		matches = append(matches, re.FindAllStringIndex(line, -1))
	}
	return matches
}

func sumGearRatios(data []string) int {
	sum := 0
	allStars := findStars(data)

	for lineIndex, stars := range allStars {
		sum += checkIfGear(data, lineIndex, stars)
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

	sum := sumGearRatios(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
