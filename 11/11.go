package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) int {
	numbersStr := strings.Split(line, " ")
	numbers := make([]int, len(numbersStr))
	var err error
	for i, n := range numbersStr {
		numbers[i], err = strconv.Atoi(n)
		check(err)
	}

	blinkNumber := 25

	for range blinkNumber {
		numbers = blink(numbers)
		// fmt.Printf("Blink %d: %v\n", i+1, numbers)
	}
	return len(numbers)
}

func blink(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			numbers[i] = 1
		} else if evenDigits(numbers[i]) {
			left, right := splitNumber(numbers[i])
			numbers = append(numbers[:i], append([]int{left, right}, numbers[i+1:]...)...)
			i++
		} else {
			numbers[i] *= 2024
		}
	}
	return numbers
}

func evenDigits(n int) bool {
	nStr := strconv.Itoa(n)
	return len(nStr)%2 == 0
}

func splitNumber(n int) (int, int) {
	nStr := strconv.Itoa(n)
	middle := len(nStr) / 2
	left, err := strconv.Atoi(nStr[:middle])
	check(err)
	right, err := strconv.Atoi(nStr[middle:])
	check(err)
	return left, right
}

func main() {
	// Open the file
	file, err := os.Open("11.in")
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
