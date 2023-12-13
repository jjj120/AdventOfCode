package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const UNUSED = -2
const PRINT = false
const ENLARGEMENT = 5

func canPlace(springs string, startIndex int, groupSize int) bool {
	for i := startIndex; i < startIndex+groupSize; i++ {
		if i >= len(springs) || springs[i] == '.' {
			return false
		}
	}

	if startIndex > 0 && springs[startIndex-1] == '#' {
		return false
	}

	if startIndex+groupSize < len(springs) && springs[startIndex+groupSize] == '#' {
		return false
	}

	return true

	// if len(springs)-startIndex < groupSize {
	// 	// not enough space anymore
	// 	// fmt.Println("Exit because not enough space")
	// 	return false
	// }

	// if startIndex > 0 && springs[startIndex-1] == '#' {
	// 	return false
	// }

	// if !(startIndex <= len(springs)-groupSize) {
	// 	return false
	// }

	// for i := 0; i < groupSize; i++ {
	// 	if springs[startIndex+i] == '.' {
	// 		return false
	// 	}
	// }

	// //foundSpace = foundSpace && (startIndex+groupSize == len(springs) || (springs[startIndex+groupSize] != '#')) // check if string has an end where it should have
	// return (startIndex == len(springs)-groupSize) || (springs[startIndex+groupSize] != '#') // check if string has an end where it should have

}

func countPoss(memory *[][]int, springs string, groups []int, startIndex int, nextGroupInd int) int {
	if nextGroupInd >= len(groups) {
		for i := startIndex; i < len(springs); i++ {
			if springs[i] == '#' {
				return 0
			}
		}
		return 1
	}
	if startIndex >= len(springs) {
		return 0
	}

	if (*memory)[startIndex][nextGroupInd] != UNUSED {
		// fmt.Println("Already Set data")
		return (*memory)[startIndex][nextGroupInd]
	}

	currPoss := 0

	if canPlace(springs, startIndex, groups[nextGroupInd]) {

		if PRINT {
			fmt.Printf("%d, %d: \t", nextGroupInd, groups[nextGroupInd])
			for i, char := range springs {
				if i >= startIndex && i < startIndex+groups[nextGroupInd] {
					fmt.Printf("\033[1m\033[36m%c\033[0m\033[0m", char)
				} else {
					fmt.Printf("%c", char)
				}
			}
			fmt.Println("")
		}
		newStart := startIndex + groups[nextGroupInd] + 1
		currPoss += countPoss(memory, springs, groups, newStart, nextGroupInd+1)
	}

	if !(startIndex > 0 && springs[startIndex-1] == '#') {
		currPoss += countPoss(memory, springs, groups, startIndex+1, nextGroupInd)
	}

	(*memory)[startIndex][nextGroupInd] = currPoss
	return currPoss
}

func splitData(line string) (string, []int) {
	stringParts := strings.Split(line, " ")

	regNum := regexp.MustCompile("[0-9]+")

	damagedSpringsNumbers := regNum.FindAllString(stringParts[1], -1)

	var numbers []int
	for i := 0; i < len(damagedSpringsNumbers); i++ {
		num, err := strconv.ParseInt(damagedSpringsNumbers[i], 10, 0)
		check(err)
		numbers = append(numbers, int(num))
	}
	return stringParts[0], numbers
}

func handleLine(line string) int {
	if ENLARGEMENT != 1 {
		line = (strings.Repeat(strings.Split(line, " ")[0]+"?", ENLARGEMENT))[:len(strings.Split(line, " ")[0])*ENLARGEMENT+ENLARGEMENT-1] + " " + strings.TrimRight(strings.Repeat(strings.Split(line, " ")[1]+",", ENLARGEMENT), ",")
	}

	data, groups := splitData(line)

	var memory [][]int
	for i := 0; i < len(data); i++ {
		memory = append(memory, make([]int, len(groups)))
		for j := 0; j < len(groups); j++ {
			(memory)[i][j] = UNUSED
		}
	}

	res := countPoss(&memory, data, groups, 0, 0)

	if PRINT {
		for i, lineMem := range memory {
			fmt.Printf("%d: ", i)
			fmt.Println(lineMem)
		}

		fmt.Println("")

		fmt.Println(line)
		fmt.Printf("found solution %d\n\n", res)
		fmt.Println(": ", res)
	}

	return res
}

func main() {
	// Open the file
	file, err := os.Open("12.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	i := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\r\r\r\r\r%d", i)
		sum += handleLine(line)
		i++
	}
	fmt.Print("\n")

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
