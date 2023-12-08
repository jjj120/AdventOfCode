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

func goThroughMap(firstLine string, nodes map[string][]string) int {
	currNode := "AAA"
	steps := 0
	for i := 0; strings.Compare(currNode, "ZZZ") != 0; i = ((i + 1) % len(firstLine)) {
		if firstLine[i] == 'L' {
			currNode = nodes[currNode][0]
		} else {
			currNode = nodes[currNode][1]
		}
		// fmt.Println(currNode)
		steps++
	}
	return steps
}

func main() {
	// Open the file
	file, err := os.Open("8.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()
	var nodes map[string][]string
	nodes = make(map[string][]string)

	scanner.Scan()
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		firstSplit := strings.Split(line, " = ")
		firstSplit[1] = strings.ReplaceAll(firstSplit[1], "(", "")
		firstSplit[1] = strings.ReplaceAll(firstSplit[1], ")", "")
		secondSplit := strings.Split(firstSplit[1], ", ")
		nodes[firstSplit[0]] = secondSplit

		// fmt.Println(nodes)
	}

	sum := goThroughMap(firstLine, nodes)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
