package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFirst(line string) int {
	for _, c := range line {
		if unicode.IsDigit(c) {
			return int(c - '0')
		}
	}
	return -1
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getLast(line string) int {
	line = Reverse(line)
	return getFirst(line)
}

func getDigit(s string) (first, last int) {
	last = -1
	first = -1

	for i, char := range s {
		// Check if the character is a spelled-out digit
		if unicode.IsDigit(char) {
			// Convert the spelled-out digit to its numerical value
			last = int(char - '0')
			if first == -1 {
				first = last
			}
		}

		// Check if the substring starting from the current position matches any spelled-out digit word
		digitWord := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}

		for word, value := range digitWord {
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				last = value
			}
		}

		if first == -1 {
			first = last
		}
	}

	return first, last
}

func main() {
	// Open the file
	file, err := os.Open("1.in")
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

		// fmt.Printf("-> 0:%d%d\n", getDigit(line))
		i1, i2 := getDigit(line)
		sum += i1*10 + i2
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
