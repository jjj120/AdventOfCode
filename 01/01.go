package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"math"
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

func combineListsAndSum(leftList, rightList []int) int {
	sum := 0
	length := len(leftList)
	for i := 0; i < length; i++ {
		min_left := slices.Min(leftList)
		min_right := slices.Min(rightList)

		index_left := slices.Index(leftList, min_left)
		index_right := slices.Index(rightList, min_right)

		sum += int(math.Abs(float64(min_left - min_right)))

		leftList = append(leftList[:index_left], leftList[index_left+1:]...)
		rightList = append(rightList[:index_right], rightList[index_right+1:]...)

		// fmt.Printf("min_left: %d, min_right: %d\n", min_left, min_right)
	}
	return sum
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

	sum := combineListsAndSum(leftList, rightList)


	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
