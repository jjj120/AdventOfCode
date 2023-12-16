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

func energize(data [][]byte, energized [][]bool, currX, currY, dirX, dirY int, history *[][4]int) {
	// fmt.Printf("I am at %d, %d, going in dir %d %d -> ", currX, currY, dirX, dirY)
	// fmt.Println(history)

	if currX < 0 || currX >= len(data[0]) || currY < 0 || currY >= len(data) {
		return
	}

	for _, elem := range *history {
		if elem[0] == currX && elem[1] == currY && elem[2] == dirX && elem[3] == dirY {
			return
		}
	}

	energized[currY][currX] = true

	*history = append(*history, [4]int{currX, currY, dirX, dirY})

	switch data[currY][currX] {
	case '.':
		energize(data, energized, currX+dirX, currY+dirY, dirX, dirY, history)
	case '/':
		if dirX == 1 && dirY == 0 {
			energize(data, energized, currX, currY-1, 0, -1, history)
		} else if dirX == -1 && dirY == 0 {
			energize(data, energized, currX, currY+1, 0, +1, history)
		} else if dirX == 0 && dirY == 1 {
			energize(data, energized, currX-1, currY, -1, 0, history)
		} else if dirX == 0 && dirY == -1 {
			energize(data, energized, currX+1, currY, +1, 0, history)
		}
	case '\\':
		if dirX == 1 && dirY == 0 {
			energize(data, energized, currX, currY+1, 0, +1, history)
		} else if dirX == -1 && dirY == 0 {
			energize(data, energized, currX, currY-1, 0, -1, history)
		} else if dirX == 0 && dirY == 1 {
			energize(data, energized, currX+1, currY, +1, 0, history)
		} else if dirX == 0 && dirY == -1 {
			energize(data, energized, currX-1, currY, -1, 0, history)
		}
	case '-':
		if dirY == 0 {
			energize(data, energized, currX+dirX, currY+dirY, dirX, dirY, history)
		} else if dirX == 0 {
			energize(data, energized, currX+1, currY, +1, 0, history)
			energize(data, energized, currX-1, currY, -1, 0, history)
		}
	case '|':
		if dirY == 0 {
			energize(data, energized, currX, currY-1, 0, -1, history)
			energize(data, energized, currX, currY+1, 0, +1, history)
		} else if dirX == 0 {
			energize(data, energized, currX+dirX, currY+dirY, dirX, dirY, history)
		}
	}
}

func countEnergized(energized [][]bool) int {
	sum := 0
	for _, row := range energized {
		for _, elem := range row {
			if elem {
				sum++
			}
		}
	}
	return sum
}

func initEnergized(data [][]byte) [][]bool {
	energized := make([][]bool, len(data))
	for i := 0; i < len(data); i++ {
		energized[i] = make([]bool, len(data[i]))
		for j := 0; j < len(data); j++ {
			energized[i][j] = false
		}
	}
	return energized
}

func handleData(data [][]byte) int {

	mostEnergized := 0
	// top and bottom rows
	for i := 0; i < len(data[0]); i++ {
		var history [][4]int
		energized := initEnergized(data)
		energize(data, energized, i, 0, 0, 1, &history)
		mostEnergized = max(mostEnergized, countEnergized(energized))

		history = make([][4]int, 0)
		energized = initEnergized(data)
		energize(data, energized, i, len(data[0])-1, 0, -1, &history)
		mostEnergized = max(mostEnergized, countEnergized(energized))
	}

	// left and right col
	for i := 0; i < len(data); i++ {
		var history [][4]int
		energized := initEnergized(data)
		energize(data, energized, 0, i, 1, 0, &history)
		mostEnergized = max(mostEnergized, countEnergized(energized))

		history = make([][4]int, 0)
		energized = initEnergized(data)
		energize(data, energized, len(data)-1, i, -1, 0, &history)
		mostEnergized = max(mostEnergized, countEnergized(energized))
	}

	// for i := 0; i < len(energized); i++ {
	// 	for j := 0; j < len(energized[i]); j++ {
	// 		if energized[i][j] {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }

	return mostEnergized
}

func main() {
	// Open the file
	file, err := os.Open("16.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var data [][]byte
	// Iterate through each lin
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []byte(line))
	}

	sum := handleData(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
