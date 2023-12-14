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

func checkLettersContaining(line string) bool {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, char1 := range letters {
		for _, char2 := range letters {
			str := string(char1) + string(char2)
			if strings.Count(line, str) >= 2 {
				return true
			}
		}
	}
	// fmt.Printf("Found no two doubles in %s\n", line)
	return false
}

func checkLettersContaining2(line string) bool {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for _, char1 := range letters {
		for _, char2 := range letters {
			str := string(char1) + string(char2) + string(char1)
			if strings.Contains(line, str) {
				return true
			}
		}
	}
	// fmt.Printf("Found no mirror in %s\n", line)
	return false
}

func handleLine(line string) int {
	if checkLettersContaining(line) && checkLettersContaining2(line) {
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
