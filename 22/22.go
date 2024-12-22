package main

import (
	"bufio"
	"fmt"
	"os"
)

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func assertf(condition bool, message string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(message, args...))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string, secretNumberIndex int) int {
	var secretNum int
	_, err := fmt.Sscanf(line, "%d", &secretNum)
	check(err)

	for i := 0; i < secretNumberIndex; i++ {
		secretNum = calcNextSecret(secretNum)
	}

	fmt.Printf("%s: %d\n", line, secretNum)

	return secretNum
}

func calcNextSecret(number int) int {
	number = pruneSecret(mixSecrets(number, number<<6))
	number = pruneSecret(mixSecrets(number, number>>5))
	number = pruneSecret(mixSecrets(number, number<<11))
	return number
}

func mixSecrets(secretNum, value int) int {
	return secretNum ^ value
}

func pruneSecret(secretNum int) int {
	return secretNum & 0xFFFFFF // mod 16777216
}

func main() {
	// Open the file
	file, err := os.Open("22.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	const secretNumberIndex = 2000
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		sum += handleLine(line, secretNumberIndex)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
