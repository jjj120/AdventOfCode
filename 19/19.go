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

func parseReplacement(line string) (string, string) {
	// Split the line into two parts at " => "
	split := strings.Split(line, " => ")
	if len(split) == 2 {
		return split[0], split[1]
	}
	panic("Invalid replacement: " + line)
}

func countMolecules(molecule string, replacements map[string][]string) int {
	molecules := make(map[string]bool)

	for from, tos := range replacements {
		for i := 0; i <= len(molecule)-len(from); i++ {
			if molecule[i:i+len(from)] == from {
				for _, to := range tos {
					newMolecule := molecule[:i] + to + molecule[i+len(from):]
					molecules[newMolecule] = true
				}
			}
		}
	}

	return len(molecules)
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

	replacements := make(map[string][]string)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			break
		}
		from, to := parseReplacement(line)
		replacements[from] = append(replacements[from], to)
	}

	scanner.Scan()
	baseMolecule := scanner.Text()

	var sum = 0
	sum = countMolecules(baseMolecule, replacements)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
