package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const maxMatches = 10

func countMatches(line string) int {
	info := strings.Split(line, ": ")
	numbers := strings.Split(info[1], " | ")
	re := regexp.MustCompile("[0-9]+")

	winningNumbersStr := re.FindAllString(numbers[0], -1)
	myNumbersStr := re.FindAllString(numbers[1], -1)

	winningNumbers := make([]int, 0)
	myNumbers := make([]int, 0)

	for _, numStr := range winningNumbersStr {
		num, err := strconv.ParseInt(numStr, 10, 0)
		check(err)
		winningNumbers = append(winningNumbers, int(num))
	}

	for _, numStr := range myNumbersStr {
		num, err := strconv.ParseInt(numStr, 10, 0)
		check(err)
		myNumbers = append(myNumbers, int(num))
	}

	matches := 0

	for _, myNumber := range myNumbers {
		for _, winningNumber := range winningNumbers {
			if myNumber == winningNumber {
				matches++
			}
		}
	}
	return matches
}

func sumMatches(lines []string) int {
	copies := make([]int, len(lines)+maxMatches)
	for j := 0; j < len(copies); j++ {
		copies[j] = 1
	}

	for lineIndex, line := range lines {
		matches := countMatches(line)
		for i := lineIndex + 1; i < lineIndex+matches+1; i++ {
			copies[i] += copies[lineIndex]
		}
	}

	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += copies[i]
	}

	return sum
}

func main() {
	// Open the file
	file, err := os.Open("4.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	cards := make([]string, 0)

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, line)
	}

	sum := sumMatches(cards)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
