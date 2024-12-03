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
	// get all multiply strings in the format "mul\(\d{0,3},\d{0,3}\)"
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(line, -1)

	res_number := 0
	for match := range matches {
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

		// check if there are two numbers
		if len(numbers) != 2 {
			fmt.Println("Error: invalid number of arguments: ", numbers, " in ", matches[match])
			for i := range numbers {
				fmt.Println(numbers[i])
			}

			os.Exit(1)
		}

		// convert the numbers to integers
		num1, err := strconv.Atoi(numbers[0])
		check(err)
		num2, err := strconv.Atoi(numbers[1])
		check(err)

		// multiply the numbers
		res_number += num1 * num2
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
