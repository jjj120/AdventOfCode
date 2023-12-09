package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkZeros(numbers []int) bool {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}

func extrapolateValues(matchesNumbers []int) []int {
	// fmt.Println(matchesNumbers)
	if checkZeros(matchesNumbers) {
		return append(matchesNumbers, 0)
	}

	diffs := make([]int, 0)
	prev := matchesNumbers[0]

	for i, match := range matchesNumbers {
		if i == 0 {
			continue
		}
		diffs = append(diffs, match-prev)
		prev = match
	}

	newEx := extrapolateValues(diffs)

	newValue := newEx[len(newEx)-1] + matchesNumbers[len(matchesNumbers)-1]
	return append(matchesNumbers, newValue)
}

func handleLine(line string) int {
	re := regexp.MustCompile("[-]?[0-9]+")
	matches := re.FindAllString(line, -1)
	matchesNumbers := make([]int, 0)
	for _, match := range matches {
		num, err := strconv.ParseInt(match, 10, 0)
		check(err)
		matchesNumbers = append(matchesNumbers, int(num))
	}
	newNumbers := extrapolateValues(matchesNumbers)
	return newNumbers[len(newNumbers)-1]
}

func main() {
	// Open the file
	file, err := os.Open("9.in")
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
}
