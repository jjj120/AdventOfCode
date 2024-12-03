package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) int {
	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	re_do := regexp.MustCompile(`do\(\)`)
	re_dont := regexp.MustCompile(`don't\(\)`)

	matches := re.FindAllString(line, -1)

	res_number := 0
	do_multiply := true

	for match := range matches {
		if re_do.MatchString(matches[match]) {
			do_multiply = true
			continue
		}
		if re_dont.MatchString(matches[match]) {
			do_multiply = false
			continue
		}

		// get the two numbers
		re_digits := regexp.MustCompile(`\d{1,3}`)
		numbers := re_digits.FindAllString(matches[match], -1)

		// strip out all empty strings
		for i := 0; i < len(numbers); i++ {
			if numbers[i] == "" {
				numbers = append(numbers[:i], numbers[i+1:]...)
				i--
			}
		}

		// convert the numbers to integers
		num1, err := strconv.Atoi(numbers[0])
		check(err)
		num2, err := strconv.Atoi(numbers[1])
		check(err)

		// multiply the numbers
		if do_multiply {
			res_number += num1 * num2
		}
	}

	return res_number
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
	combinedLines := ""
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		combinedLines += line
	}

	sum = handleLine(combinedLines)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
