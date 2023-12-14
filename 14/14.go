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

const PRINT = true
const NUM_ITERATIONS = 1000000000

func shiftUp(data [][]byte) {
	for colIndex := 0; colIndex < len(data[0]); colIndex++ {
		lastStonePos := -1
		for rowIndex := 0; rowIndex < len(data); rowIndex++ {
			if data[rowIndex][colIndex] == '#' {
				lastStonePos = rowIndex
				continue
			}
			if data[rowIndex][colIndex] == 'O' {
				data[rowIndex][colIndex] = '.'
				lastStonePos++
				data[lastStonePos][colIndex] = 'O'
			}
		}
	}
}

func countLoad(data [][]byte) int {
	load := 0
	for lineIndex, line := range data {
		for _, c := range line {
			if c == 'O' {
				load += len(data) - lineIndex
			}
		}
	}
	return load
}

func spinCircle(data [][]byte) {
	//north
	for colIndex := 0; colIndex < len(data[0]); colIndex++ {
		lastStonePos := -1
		for rowIndex := 0; rowIndex < len(data); rowIndex++ {
			if data[rowIndex][colIndex] == '#' {
				lastStonePos = rowIndex
			}
			if data[rowIndex][colIndex] == 'O' {
				data[rowIndex][colIndex] = '.'
				lastStonePos++
				data[lastStonePos][colIndex] = 'O'
			}
		}
	}

	//west
	for rowIndex := 0; rowIndex < len(data); rowIndex++ {
		lastStonePos := -1
		for colIndex := 0; colIndex < len(data[0]); colIndex++ {
			if data[rowIndex][colIndex] == '#' {
				lastStonePos = colIndex
			}
			if data[rowIndex][colIndex] == 'O' {
				data[rowIndex][colIndex] = '.'
				lastStonePos++
				data[rowIndex][lastStonePos] = 'O'
			}
		}
	}

	//south
	for colIndex := 0; colIndex < len(data[0]); colIndex++ {
		lastStonePos := len(data)
		for rowIndex := len(data) - 1; rowIndex >= 0; rowIndex-- {
			if data[rowIndex][colIndex] == '#' {
				lastStonePos = rowIndex
			}
			if data[rowIndex][colIndex] == 'O' {
				data[rowIndex][colIndex] = '.'
				lastStonePos--
				data[lastStonePos][colIndex] = 'O'
			}
		}
	}

	//east
	for rowIndex := len(data) - 1; rowIndex >= 0; rowIndex-- {
		lastStonePos := len(data)
		for colIndex := len(data[0]) - 1; colIndex >= 0; colIndex-- {
			if data[rowIndex][colIndex] == '#' {
				lastStonePos = colIndex
			}
			if data[rowIndex][colIndex] == 'O' {
				data[rowIndex][colIndex] = '.'
				lastStonePos--
				data[rowIndex][lastStonePos] = 'O'
			}
		}
	}

}

func printData(data [][]byte) {
	if PRINT {
		for _, line := range data {
			for _, c := range line {
				fmt.Printf("%c", c)
			}
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func handleData(data [][]byte) int {
	// fmt.Println(NUM_ITERATIONS)
	// for i := 0; i < NUM_ITERATIONS; i++ {
	// 	if i%100000 == 0 {
	// 		fmt.Printf("\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r%d", i)
	// 		fmt.Printf(" %d%%", i*100/NUM_ITERATIONS)
	// 	}
	// 	spinCircle(data)
	// }

	shiftUp(data)

	printData(data)

	return countLoad(data)
}

func main() {
	// Open the file
	file, err := os.Open("14.in")
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
