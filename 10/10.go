package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const REPEAT_TIMES = 50

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lookAndSay(lineArr []int) []int {
	currNum := lineArr[0]
	currCount := 1
	var newArr []int

	for i := 1; i < len(lineArr); i++ {
		if currNum == lineArr[i] {
			currCount++
		} else {
			newArr = append(newArr, currCount, currNum)
			currCount = 1
			currNum = lineArr[i]
		}
	}
	newArr = append(newArr, currCount, currNum)
	return newArr
}

func handleLine(line string) int {
	lineArrStr := strings.Split(line, "")
	var lineArr []int
	for _, elem := range lineArrStr {
		num, err := strconv.ParseInt(elem, 10, 0)
		check(err)
		lineArr = append(lineArr, int(num))
	}

	for i := 0; i < REPEAT_TIMES; i++ {
		lineArr = lookAndSay(lineArr)
	}

	return len(lineArr)
}

func main() {
	// Open the file
	file, err := os.Open("10.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		sum += handleLine(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
