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

func checkVowels(line string) bool {
	vowels := 0
	for _, char := range line {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowels++
		}
	}
	return (vowels >= 3)
}

func checkLettersContaining(line string) bool {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range letters {
		str := "" + string(char) + string(char)
		if strings.Contains(line, str) {
			return true
		}
	}
	return false
}

func checkLettersNotContaining(line string) bool {
	return !strings.Contains(line, "ab") && !strings.Contains(line, "cd") && !strings.Contains(line, "pq") && !strings.Contains(line, "xy")
}

func handleLine(line string) int {
	if checkVowels(line) && checkLettersContaining(line) && checkLettersNotContaining(line) {
		return 1
	}
	return 0
}

func main() {
	// Open the file
	file, err := os.Open("05.in")
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
