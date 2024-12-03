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
	report := make([]int, len(strings.Split(line, " ")))

	for i, v := range strings.Split(line, " ") {
		n, err := strconv.Atoi(v)
		check(err)
		report[i] = n
	}

	if checkSafetyWithOneLess(report) {
		return 1
	}

	return 0
}

func checkSafetyWithOneLess(report []int) bool {
	for i := 0; i < len(report); i++ {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])
		if checkSafety(newReport) {
			return true
		}
	}
	return false
}

func checkSafety(report []int) bool {
	// check if all increasing or decreasing
	increasing := true
	decreasing := true

	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			increasing = false
		}
		if report[i] > report[i+1] {
			decreasing = false
		}
		if report[i] == report[i+1] {
			return false
		}
		if !increasing && !decreasing {
			return false
		}

		// have to differ by at most 3
		if report[i]-report[i+1] > 3 || report[i]-report[i+1] < -3 {
			return false
		}
	}
	return increasing || decreasing
}

func main() {
	// Open the file
	file, err := os.Open("02.in")
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
