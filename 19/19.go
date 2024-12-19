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

func parseTowels(line string) ([]string, map[byte][]string) {
	towelsList := strings.Split(line, ", ")
	towels := make(map[byte][]string)
	for _, towel := range towelsList {
		towels[towel[0]] = append(towels[towel[0]], towel)
	}
	return towelsList, towels
}

func checkTowel(towelToCheck string, towels map[byte][]string) int {
	possArray := make([]int, len(towelToCheck)+1)
	// save the number of possibilities for each index
	possArray[0] = 1

	for idxToCheck := 0; idxToCheck < len(towelToCheck); idxToCheck++ {
		towelsLst := towels[towelToCheck[idxToCheck]]
		for _, towel := range towelsLst {
			if strings.HasPrefix(towelToCheck[idxToCheck:], towel) {
				possArray[idxToCheck+len(towel)] += possArray[idxToCheck]
			}
		}
	}
	return possArray[len(possArray)-1]
}

func main() {
	// Open the file
	file, err := os.Open("19.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()
	_, towelsMap := parseTowels(firstLine)

	scanner.Scan() // Skip empty line

	var sum = 0
	numLines := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		numLines++
		sum += checkTowel(line, towelsMap)

		// fmt.Printf("Line: %s, Possibilities: %d\n", line, checkTowel(line, towelsMap))
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	if sum <= 29271194035771 {
		fmt.Println("Incorrect sum 29271194035771")
		return
	}

	fmt.Printf("Sum: %d\n", sum)
}
