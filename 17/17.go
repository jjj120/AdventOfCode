package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) int {
	num, err := strconv.Atoi(line)
	check(err)
	return num
}

func countCombinations(containers []int, target int) int {
	if target == 0 {
		return 1
	}

	if len(containers) == 0 {
		return 0
	}

	var sum = 0
	for i := 0; i < len(containers); i++ {
		if containers[i] <= target {
			sum += countCombinations(containers[i+1:], target-containers[i])
		}
	}
	return sum
}

func main() {
	// Open the file
	file, err := os.Open("17.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	var containers []int
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		containers = append(containers, handleLine(line))
	}

	sum = countCombinations(containers, 150)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
