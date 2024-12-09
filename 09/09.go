package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) int {
	expanded := expandLine(line)
	// printFormatted(expanded)

	shifted := shiftDown(expanded)
	// printFormatted(shifted)
	saveToFile(shifted)
	return calcChecksum(shifted)
}

func expandLine(line string) []int {
	expanded := []int{}
	for i := 0; i < len(line); i++ {
		length, err := strconv.Atoi(string(line[i]))
		check(err)

		if i%2 == 0 {
			// file
			for j := 0; j < length; j++ {
				expanded = append(expanded, i/2)
			}
		} else {
			// free
			for j := 0; j < length; j++ {
				expanded = append(expanded, -1)
			}
		}
	}
	return expanded
}

func shiftDown(line []int) []int {
	shifted := make([]int, 0, len(line))

	for i := 0; i < len(line); i++ {
		if line[i] == -1 {
			for j := len(line) - 1; j > 0; j-- {
				// delete the last element if it is -1
				if j <= i {
					return shifted
				}
				if line[j] == -1 {
					line = line[:j]
				} else {
					break
				}
			}

			shifted = append(shifted, line[len(line)-1])
			line = line[:len(line)-1]
		} else {
			shifted = append(shifted, line[i])
		}
	}
	return shifted
}

func calcChecksum(line []int) int {
	sum := 0
	for i := 0; i < len(line); i++ {
		sum += i * line[i]
	}
	return sum
}

func printFormatted(line []int) {
	for i := 0; i < len(line); i++ {
		if line[i] == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(line[i])
		}
	}
	fmt.Println()
}

func saveToFile(line []int) {
	f, err := os.Create("output.txt")
	check(err)
	defer f.Close()

	for i := 0; i < len(line); i++ {
		if line[i] == -1 {
			f.WriteString(".")
		} else {
			f.WriteString(strconv.Itoa(line[i]))
		}
	}
	f.WriteString("\n")
}

func main() {
	// Open the file
	file, err := os.Open("09.in")
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
