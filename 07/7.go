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

func checkCards(cards map[rune]int, number int) bool {
	// check for 5 same
	for _, num := range cards {
		if num == number {
			return true
		}
	}
	return false
}

func checkFullHouse(cards map[rune]int, jokers int) bool {
	three := false
	threeCard := '0'
	// use jokers for first and no jokers for second
	for card, num := range cards {
		if num == 3 {
			three = true
			threeCard = card
		}
	}

	for card, num := range cards {
		if (num-jokers) == 2 && card != threeCard {
			return three
		}
	}
	return false
}

func checkDoublePair(cards map[rune]int, jokers int) bool {
	two := false
	twoCard := '0'
	// use jokers for first and no jokers for second
	for card, num := range cards {
		if num == 2 {
			two = true
			twoCard = card
		}
	}

	for card, num := range cards {
		if (num-jokers) == 2 && card != twoCard {
			return two
		}
	}
	return false
}

func getType(line string) int {
	// var cards map[string]int;
	cards := map[rune]int{
		'A': 0,
		'K': 0,
		'Q': 0,
		'T': 0,
		'9': 0,
		'8': 0,
		'7': 0,
		'6': 0,
		'5': 0,
		'4': 0,
		'3': 0,
		'2': 0,
		'J': 0,
	}

	// count occuring numbers
	for _, char := range line {
		if char == ' ' {
			break
		}
		cards[char]++
	}

	// increase cards by joker amount
	jokers := cards['J']
	cards['J'] = 0

	for card := range cards {
		if card != 'J' {
			cards[card] += jokers
		}
	}

	if checkCards(cards, 5) {
		return 6
	}
	if checkCards(cards, 4) {
		return 5
	}
	if checkFullHouse(cards, jokers) {
		return 4
	}
	if checkCards(cards, 3) {
		return 3
	}
	if checkDoublePair(cards, jokers) {
		return 2
	}
	if checkCards(cards, 2) {
		return 1
	}
	return 0
}

func getCardValue(card rune) int {
	cards := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}
	return cards[card]
}

func sortArray(array *[]string) {
	sort.Slice((*array), func(i, j int) bool {
		if i >= 1000 || j >= 1000 {
			fmt.Printf("Invalid read at %d or %d", i, j)
		}
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

		if len(linesI[0]) != 5 {
			fmt.Printf("%s has length %d\n", linesI[0], len(linesI[0]))
		}

		if len(linesJ[0]) != 5 {
			fmt.Printf("%s has length %d\n", linesJ[0], len(linesJ[0]))
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
		// fmt.Printf("%d: %s\n", i, line)
		lineSplit := strings.Split(line, " ")
		if len(lineSplit[0]) != 5 {
			fmt.Printf("%s has length %d\n", lineSplit[0], len(lineSplit[0]))
		}
		num, err := strconv.ParseInt(lineSplit[1], 10, 0)
		check(err)
		// fmt.Printf("Type: %d %s -> %d\n", getType(lineSplit[0]), line, i+1)
		sum += (int(num) * (i + 1))
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

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sortArray(&array)

	// for _, cards := range array {
	// 	fmt.Println(cards)
	// }

	sum := calcBids(&array)

	fmt.Printf("Sum: %d\n", sum)
}
