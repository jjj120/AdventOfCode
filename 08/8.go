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

func singleStep(dir byte, currNode string, nodes map[string][]string) string {
	if dir == 'L' {
		return nodes[currNode][0]
	} else {
		return nodes[currNode][1]
	}
}

func currNodesEndNodes(currNode []string) bool {
	for _, node := range currNode {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func ggt(num1 int, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}

func kgv(num1 int, num2 int) int {
	return (num1 * num2) / ggt(num1, num2)
}

func kgvSteps(steps []int) int {
	if len(steps) == 0 {
		return 0
	}
	res := steps[0]
	for i := 1; i < len(steps); i++ {
		res = kgv(res, steps[i])
	}
	return res
}

func goThroughMap(firstLine string, nodes map[string][]string, startNodes []string) int {
	steps := make([]int, len(startNodes))

	for i, start := range startNodes {
		for start[2] != 'Z' {
			nextStep := firstLine[steps[i]%len(firstLine)]
			if nextStep == 'L' {
				start = nodes[start][0]
			} else {
				start = nodes[start][1]
			}
			steps[i]++
		}
	}

	res := kgvSteps(steps)

	return res
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
	nodes := make(map[string][]string)

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

	startNodes := make([]string, 0)
	for node := range nodes {
		if node[2] == 'A' {
			startNodes = append(startNodes, node)
		}
	}
	sum := goThroughMap(firstLine, nodes, startNodes)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
