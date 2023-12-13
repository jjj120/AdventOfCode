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

// from https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func lineEquals(block [][]string, line1 int, line2 int) int {
	// if (line1 < 0 || line1 >= len(block)) || (line2 < 0 || line2 >= len(block)) {
	// 	return false
	// }
	diffs := 0
	for j, char := range block[line1] {
		if strings.Compare(char, block[line2][j]) != 0 {
			diffs++
		}
	}
	return diffs
}

func checkHMirror(block [][]string, mirrorAfter int) bool {
	if mirrorAfter >= len(block)-1 {
		return false
	}

	if mirrorAfter == 0 {
		return lineEquals(block, 0, 1) == 1
	}

	diffs := 0
	for offset := 0; offset <= min(mirrorAfter, len(block)-mirrorAfter-2); offset++ {
		diffs += lineEquals(block, mirrorAfter-offset, mirrorAfter+offset+1)
		if diffs > 1 {
			return false
		}
	}

	return diffs == 1
}

func getHorizontalMirror(block [][]string) int {
	for mirrorAfter := 0; mirrorAfter < len(block); mirrorAfter++ {
		if checkHMirror(block, mirrorAfter) {
			// fmt.Printf("Found mirror after %d\n", mirrorAfter)
			return mirrorAfter
		}
	}
	return -1
}

func getVerticalMirror(block [][]string) int {
	return getHorizontalMirror(transpose(block))
}

func handleBlock(block []string) int {
	var blockExpanded [][]string
	for _, line := range block {
		blockExpanded = append(blockExpanded, strings.Split(line, ""))
	}

	sum := getVerticalMirror(blockExpanded) + 1
	return sum + 100*(getHorizontalMirror(blockExpanded)+1)
}

func handleData(data [][]string) int {
	sum := 0
	for _, block := range data {
		sum += handleBlock(block)
	}
	return sum
}

func main() {
	// Open the file
	file, err := os.Open("13.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var data [][]string
	// Iterate through each line
	var block []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			data = append(data, block)
			block = make([]string, 0)
		} else {
			block = append(block, line)
		}
	}
	data = append(data, block)

	sum := handleData(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
