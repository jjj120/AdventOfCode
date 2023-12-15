package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hashPart(part string) int {
	currValue := 0
	for _, char := range part {
		currValue += int(char)
		currValue *= 17
		currValue %= 256
	}
	return currValue
}

func handleLine(line string) int {
	lineParts := strings.Split(line, ",")
	sum := 0
	for _, part := range lineParts {
		sum += hashPart(part)
	}
	return sum
}

func main() {
	// Open the file
	file, err := os.Open("15.in")
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
