package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string) (int, int) {
	var left, right int
	_, err := fmt.Sscanf(line, "%d   %d", &left, &right)
	if err != nil {
		fmt.Println("Error parsing line:", err)
	}
	return left, right
}

func calcSimScore(leftList, rightList []int) int {
	simScore := 0
	length := len(leftList)
	for i := 0; i < length; i++ {
		number := leftList[i]
		count := countAppearences(rightList, number)
		simScore += count * number
	}
	return simScore
}

func countAppearences(list []int, number int) int {
	count := 0
	for _, value := range list {
		if value == number {
			count++
		}
	}
	return count
}

func main() {
	// Open the file
	file, err := os.Open("01.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var leftList, rightList []int
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		left, right := parseLine(line)
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	// fmt.Printf("leftList: %v\n", leftList)
	// fmt.Printf("rightList: %v\n", rightList)

	sum := calcSimScore(leftList, rightList)


	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
