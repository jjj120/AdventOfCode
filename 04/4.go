package main

import (
	"bufio"
	"fmt"
	"math"
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

func handleLine(line string) int {
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

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
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
