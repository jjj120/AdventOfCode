package main

import (
	"bufio"
	"fmt"
	"math"
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

	blinkNumber := 75

	numbersAmount := map[int]int{}

	for _, n := range numbers {
		numbersAmount[n]++
	}

	for range blinkNumber {
		numbersAmount = blink(numbersAmount)
	}
	sum := 0
	for _, amount := range numbersAmount {
		sum += amount
	}
	return sum
}

func blink(numbers map[int]int) map[int]int {
	numbersNew := make(map[int]int)

	for num, amount := range numbers {
		if num == 0 {
			numbersNew[1] += amount
		} else if even, dig := evenDigits(num); even {
			left, right := splitNumber(num, dig)
			numbersNew[left] += amount
			numbersNew[right] += amount
		} else {
			numbersNew[num*2024] += amount
		}
	}
	return numbersNew
}

func evenDigits(n int) (bool, int) {
	digits := 0
	for n > 0 {
		n /= 10
		digits++
	}
	return digits%2 == 0, digits
}

func splitNumber(n int, digits int) (int, int) {
	half := digits / 2
	divisor := int(math.Pow(10, float64(half)))
	right := n % divisor
	left := n / divisor
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
