package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const RED = 0
const GREEN = 1
const BLUE = 2

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCurrInfo(data string) (int, int) {
	d := strings.Split(data, " ")
	num, err := strconv.Atoi(d[0])
	check(err)

	var col = -1
	switch d[1] {
	case "red":
		col = RED
	case "blue":
		col = BLUE
	case "green":
		col = GREEN
	default:
		fmt.Printf("Error on parsing string %s to color\n", d[1])
	}

	return num, col
}

func getPower(line string) int {
	splitData := strings.Split(line, ": ")
	games := strings.Split(splitData[1], "; ")

	var min [3]int

	for _, game := range games {
		// fmt.Printf("gameData: %s\n", game)
		gameData := strings.Split(game, ", ")
		for _, data := range gameData {
			// fmt.Printf("\tdata: %s\n", data)
			num, color := getCurrInfo(data)

			if num > min[color] {
				min[color] = num
			}

		}
		// fmt.Print("\n")
	}

	return min[0] * min[1] * min[2]
}

func main() {
	// Open the file
	file, err := os.Open("2.in")
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
		sum += getPower(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
