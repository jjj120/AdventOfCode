package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) int {
	var history map[[2]int]int = make(map[[2]int]int)
	currPos := [2]int{0, 0}
	history[currPos]++

	for _, dir := range line {
		switch dir {
		case '^':
			currPos[1]++
		case '>':
			currPos[0]++
		case 'v':
			currPos[1]--
		case '<':
			currPos[0]--
		}

		history[currPos]++
	}

	return len(history)
}

func main() {
	// Open the file
	file, err := os.Open("03.in")
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
