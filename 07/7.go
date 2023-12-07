package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getType(line string) int {
	// var cards map[string]int;
	cards := map[rune]int{
		'A': 0,
		'K': 0,
		'Q': 0,
		'J': 0,
		'T': 0,
		'9': 0,
		'8': 0,
		'7': 0,
		'6': 0,
		'5': 0,
		'4': 0,
		'3': 0,
		'2': 0,
	}

	for _, char := range line {
		cards[char]++
		if char == ' ' {
			break
		}
	}

	for _, num := range cards {
		if num == 5 {
			return 6
		}
		if num == 4 {
			return 5
		}
	}

	// check for five of kind
	for _, num := range cards {
		if num == 5 {
			return 6
		}
		if num == 4 {
			return 5
		}
	}

	three := false
	two := 0
	for _, num := range cards {
		if num == 3 {
			three = true
		}
		if num == 2 {
			two++
		}
	}
	if three && (two == 1) {
		return 4
	}
	if three {
		return 3
	}
	if two == 2 {
		return 2
	}
	if two == 1 {
		return 1
	}
	return 0
}

func getCardValue(card rune) int {
	cards := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}
	return cards[card]
}

func sortArray(array *[]string) {
	sort.Slice(*array, func(i, j int) bool {
		linesI := strings.Split((*array)[i], " ")
		linesJ := strings.Split((*array)[j], " ")

		typeI := getType(linesI[0])
		typeJ := getType(linesJ[0])

		if typeI < typeJ {
			return true
		}
		if typeI > typeJ {
			return false
		}

		iterator := 0
		for ; iterator < 5; iterator++ {
			if linesI[0][iterator] != linesJ[0][iterator] {
				break
			}
		}

		if getCardValue([]rune(linesI[0])[iterator]) < getCardValue([]rune(linesJ[0])[iterator]) {
			return true
		}
		if getCardValue([]rune(linesI[0])[iterator]) > getCardValue([]rune(linesJ[0])[iterator]) {
			return false
		}

		return false
	})
}

func calcBids(array *[]string) int {
	sum := 0
	for i, line := range *array {
		lineSplit := strings.Split(line, " ")
		num, err := strconv.ParseInt(lineSplit[1], 10, 0)
		check(err)
		sum += int(num) * (i + 1)
	}
	return sum
}

func main() {
	// Open the file
	file, err := os.Open("7.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	array := make([]string, 0)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}

	sortArray(&array)

	sum := calcBids(&array)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
