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
	queue := []string{towelToCheck}
	alreadyChecked := make(map[string]bool)

	for len(queue) > 0 {
		currTowelToCheck := queue[0]
		queue = queue[1:]

		for _, towel := range towels[currTowelToCheck[0]] {
			if strings.HasPrefix(currTowelToCheck, towel) {
				if len(currTowelToCheck) == len(towel) {
					return 1
				}

				if alreadyChecked[currTowelToCheck[len(towel):]] {
					continue
				}

				alreadyChecked[currTowelToCheck[len(towel):]] = true
				queue = append(queue, currTowelToCheck[len(towel):])
			}
		}
	}
	return 0
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
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		sum += checkTowel(line, towelsMap)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
