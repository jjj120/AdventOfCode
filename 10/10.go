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

func findStart(data []string) (int, int) {
	for y, line := range data {
		for x, char := range line {
			if char == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func getDir(data []string, coord []int) byte {
	return data[coord[1]][coord[0]]
}

func checkBounds(data []string, coord []int) bool {
	return coord[0] >= 0 && coord[1] >= 0 && coord[0] < len(data[coord[1]]) && coord[1] < len(data)
}

func getStartContinue(data []string, startCoord []int) []int {
	currCoord := []int{startCoord[0], startCoord[1]}

	currCoord[0] += 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '-' {
		return currCoord
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[0] -= 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '-' {
		return currCoord
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[1] += 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '|' {
		return currCoord
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[1] -= 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '|' {
		return currCoord
	}
	return []int{startCoord[0], startCoord[1]}
}

func nextStep(data []string, currCoord []int, prevCoord []int) []int {
	toTry := []int{currCoord[0], currCoord[1]}

	switch getDir(data, currCoord) {
	case 'S':
		return getStartContinue(data, currCoord)
	case '|':
		toTry[1] += 1
	case '-':
		toTry[0] += 1
	case 'F':
		toTry[0] += 1
	case 'J':
		toTry[0] -= 1
	case 'L':
		toTry[0] += 1
	case '7':
		toTry[0] -= 1
	}

	if checkBounds(data, toTry) && (toTry[0] != prevCoord[0] || toTry[1] != prevCoord[1]) {
		return toTry
	}

	copy(toTry, currCoord)
	switch getDir(data, currCoord) {
	case 'S':
		return getStartContinue(data, currCoord)
	case '|':
		toTry[1] -= 1
	case '-':
		toTry[0] -= 1
	case 'F':
		toTry[1] += 1
	case 'J':
		toTry[1] -= 1
	case 'L':
		toTry[1] -= 1
	case '7':
		toTry[1] += 1
	}

	if checkBounds(data, toTry) && (toTry[0] != prevCoord[0] || toTry[1] != prevCoord[1]) {
		return toTry
	}

	fmt.Printf("Error on %d %d with prev %d %d\n", currCoord[0], currCoord[1], prevCoord[0], prevCoord[1])
	return currCoord
}

func findStartPipeType(data []string) byte {
	startCoordX, startCoordY := findStart(data)
	startCoord := []int{startCoordX, startCoordY}
	currCoord := []int{startCoord[0], startCoord[1]}

	top := false
	left := false
	right := false
	bottom := false

	currCoord[0] += 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '-' {
		right = true
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[0] -= 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '-' {
		left = true
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[1] += 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '|' {
		bottom = true
	}

	currCoord = []int{startCoord[0], startCoord[1]}
	currCoord[1] -= 1
	if checkBounds(data, currCoord) && getDir(data, currCoord) == '|' {
		top = true
	}

	if top && bottom {
		return '|'
	}
	if left && right {
		return '-'
	}
	if top && left {
		return 'J'
	}
	if top && right {
		return 'L'
	}
	if left && bottom {
		return '7'
	}
	if right && bottom {
		return 'F'
	}

	return '.'
}

func sumIncluded(line []byte) int {
	sum := 0
	add := false
	uTop := false
	uBot := false

	for i, char := range line {
		// toggle only on cross with line
		// cross with line is '|', '7', 'J'
		// only sum cells with '.'

		if uTop || uBot {
			if char == 'J' {
				if uBot {
					add = !add
				}
				uTop = false
				uBot = false
			} else if char == '7' {
				if uTop {
					add = !add
				}
				uTop = false
				uBot = false
			}
			continue
		}

		if char == '|' {
			add = !add
		}

		if char == 'F' {
			uBot = true
		}

		if char == 'L' {
			uTop = true
		}

		if add && char == '.' {
			sum++
			line[i] = 'O'
		}
	}
	return sum
}

func printData(data [][]byte) {
	for _, line := range data {
		for _, char := range line {
			switch char {
			case 'O':
				fmt.Printf("\033[1m\033[36m■\033[0m\033[0m")
			case 'S':
				fmt.Printf("\033[1m\033[31mS\033[0m\033[0m")
			case 'F':
				fmt.Printf("╔")
			case 'L':
				fmt.Printf("╚")
			case '7':
				fmt.Printf("╗")
			case 'J':
				fmt.Printf("╝")
			case '|':
				fmt.Printf("║")
			case '-':
				fmt.Printf("═")
			case '.':
				fmt.Printf(" ")
			default:
				fmt.Printf("%c", char)
			}
		}
		fmt.Printf("\n")
	}
}

func getEnclosed(data []string, loop [][]int) int {
	dataNew := make([][]string, 0)
	dataSanitized := make([][]byte, 0)

	for _, line := range data {
		dataNew = append(dataNew, strings.Split(line, ""))
	}

	for _, coord := range loop {
		dataNew[coord[1]][coord[0]] = "x"
	}

	for i := range data {
		var currLine []byte
		for j := range data[i] {
			if strings.Compare(dataNew[i][j], "x") == 0 {
				currLine = append(currLine, data[i][j])
			} else {
				currLine = append(currLine, '.')
			}
		}
		dataSanitized = append(dataSanitized, currLine)
	}
	startX, startY := findStart(data)
	dataSanitized[startY][startX] = findStartPipeType(data)

	// printData(dataSanitized)
	sum := 0
	for _, line := range dataSanitized {
		sum += sumIncluded(line)
	}
	printData(dataSanitized)

	return sum
}

func getLoop(startX int, startY int, data []string) int {
	var loop [][]int
	prev := []int{startX, startY}
	loop = append(loop, []int{prev[0], prev[1]})

	currCoord := nextStep(data, prev, prev)
	loop = append(loop, []int{currCoord[0], currCoord[1]})

	length := 1

	for getDir(data, currCoord) != 'S' {
		newCurrCoord := nextStep(data, currCoord, prev)
		prev = currCoord
		currCoord = newCurrCoord
		length++
		loop = append(loop, []int{currCoord[0], currCoord[1]})
		// if length%100 == 0 {
		// 	fmt.Printf("%d: curr: %d %d, prev: %d %d\n", length, currCoord[0], currCoord[1], prev[0], prev[1])
		// }
	}

	return getEnclosed(data, loop)
}

func handleData(data []string) int {
	startX, startY := findStart(data)
	// fmt.Printf("Start: %d %d\n", startX, startY)

	return getLoop(startX, startY, data)
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

	var data []string
	// Iterate through each line
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	sum := handleData(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
