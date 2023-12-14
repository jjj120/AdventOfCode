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

func lineShift(channelIn chan []byte, channelOut chan []byte) {
	line := <-channelIn
	if line[0] == byte(0) {
		return
	}
	lastStonePos := -1
	for colIndex := 0; colIndex < len(line); colIndex++ {
		if line[colIndex] == '#' {
			lastStonePos = colIndex
		}
		if line[colIndex] == 'O' {
			line[colIndex] = '.'
			lastStonePos++
			line[lastStonePos] = 'O'
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

func shiftNorth(data [][]byte) {
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
}

func shiftWest(data [][]byte) {
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
}

func shiftSouth(data [][]byte) {
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
}

func shiftEast(data [][]byte) {
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

func spinCircle(data [][]byte) {
	shiftNorth(data)
	shiftWest(data)
	shiftSouth(data)
	shiftEast(data)
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

func deepCopy(original [][]byte) [][]byte {
	copiedSlices := make([][]byte, len(original))

	for i, slice := range original {
		copiedSlice := make([]byte, len(slice))
		copy(copiedSlice, slice)
		copiedSlices[i] = copiedSlice
	}

	return copiedSlices
}

func checkEqual(data1 [][]byte, data2 [][]byte) bool {
	for lineIndex, line := range data1 {
		for charIndex, c := range line {
			if c != data2[lineIndex][charIndex] {
				return false
			}
		}
	}
	return true
}

func contains(prevData [][][]byte, data [][]byte) int {
	for i, element := range prevData {
		if checkEqual(data, element) {
			return i
		}
	}
	return -1
}

func handleData(data [][]byte) int {
	prevData := make([][][]byte, 0)

	firstRec := -1
	secondRec := -1

	for i := 0; i < NUM_ITERATIONS; i++ {
		// if i%100000 == 0 {
		// 	fmt.Printf("\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r%d", i)
		// 	fmt.Printf(" in %s", time.Since(start))
		// 	fmt.Printf(" %d%%          ", i*100/NUM_ITERATIONS)
		// }

		spinCircle(data)

		containsElem := contains(prevData, data)
		if containsElem != -1 {
			firstRec = containsElem
			secondRec = i
			break
		}

		prevData = append(prevData, deepCopy(data))
	}

	period := secondRec - firstRec
	numRuns := (NUM_ITERATIONS - (firstRec + 1)) % period

	fmt.Printf("Found period %d between %d and %d, running %d more\n\n", period, firstRec, secondRec, numRuns)

	for i := 0; i < numRuns; i++ {
		spinCircle(data)
	}

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
