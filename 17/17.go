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

func countCombinations(containers []int, target int, numberOfContainersLeft int) int {
	if target == 0 {
		if numberOfContainersLeft == 0 {
			return 1
		}
		return 0
	}

	if len(containers) == 0 {
		return 0
	}

	var sum = 0
	for i := 0; i < len(containers); i++ {
		if containers[i] <= target {
			sum += countCombinations(containers[i+1:], target-containers[i], numberOfContainersLeft-1)
		}
	}

	return sum
}

func findMinNumberOfContainers(containers []int, target int) int {
	if target == 0 {
		return 0
	}

	if len(containers) == 0 {
		return 1000000
	}

	var minimum = 1000000
	for i := 0; i < len(containers); i++ {
		if containers[i] <= target {
			num := findMinNumberOfContainers(containers[i+1:], target-containers[i])
			minimum = min(minimum, num)
		}
	}
	return minimum + 1
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

	minumum := findMinNumberOfContainers(containers, 150)

	fmt.Println(minumum)

	sum = countCombinations(containers, 150, minumum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
