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

func handleLine(line string) int {
	splitValues := strings.Split(line, ": ")
	testResult, err := strconv.Atoi(splitValues[0])
	check(err)

	testValuesStr := strings.Split(splitValues[1], " ")
	testValues := make([]int, len(testValuesStr))
	for i := range testValuesStr {
		testValues[i], err = strconv.Atoi(testValuesStr[i])
		check(err)
	}

	totals := make([]int, 1, 100)
	totals[0] = testValues[0]
	testValues = testValues[1:]

	for value := range testValues {
		totals = addAndMultiply(totals, testValues[value])
	}

	for _, total := range totals {
		if total == testResult {
			return testResult
		}
	}

	return 0
}

func addAndMultiply(totals []int, value int) []int {
	totalsLen := len(totals)

	for i := 0; i < totalsLen; i++ {
		totals = append(totals, totals[i]+value)
		totals[i] *= value
	}

	return totals
}

func main() {
	// Open the file
	file, err := os.Open("07.in")
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
