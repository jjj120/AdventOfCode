package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const numRed = 12
const numGreen = 13
const numBlue = 14

const RED = 0
const GREEN = 1
const BLUE = 2

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getGameNumber(info string) (int, error) {
	return strconv.Atoi(strings.Split(info, " ")[1])
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

func checkPossibility(line string) int {
	splitData := strings.Split(line, ": ")
	gameInfo := splitData[0]
	games := strings.Split(splitData[1], "; ")

	gameNumber, err := getGameNumber(gameInfo)
	check(err)

	for _, game := range games {
		// fmt.Printf("gameData: %s\n", game)
		gameData := strings.Split(game, ", ")
		for _, data := range gameData {
			// fmt.Printf("\tdata: %s\n", data)
			num, color := getCurrInfo(data)
			switch color {
			case RED:
				if num > numRed {
					// fmt.Printf("Game %d impossible: %s\n", gameNumber, splitData[1])
					return 0
				}

			case BLUE:
				if num > numBlue {
					// fmt.Printf("Game %d impossible: %s\n", gameNumber, splitData[1])
					return 0
				}

			case GREEN:
				if num > numGreen {
					// fmt.Printf("Game %d impossible: %s\n", gameNumber, splitData[1])
					return 0
				}
			}
		}
		// fmt.Print("\n")
	}

	// fmt.Printf("Game %d possible: %s\n", gameNumber, splitData[1])

	return gameNumber
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
		sum += checkPossibility(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
