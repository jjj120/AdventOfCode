package main

import (
	"bufio"
	"fmt"
	"os"
)

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func assertf(condition bool, message string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(message, args...))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Lock struct {
	pins   []int
	height int
}
type Key struct {
	pins   []int
	height int
}

func parseLocksKeys(lines []string) ([]Lock, []Key) {
	lines = append(lines, "")
	locks := []Lock{}
	keys := []Key{}

	currToParse := []string{}
	for _, line := range lines {
		if line == "" {
			if len(currToParse) == 0 {
				continue
			}

			currToParse = transposeStr(currToParse)
			counts := []int{}
			for _, row := range currToParse {
				count := 0
				for _, c := range row {
					if c == '#' {
						count++
					}
				}
				counts = append(counts, count-1)
			}

			if currToParse[0][0] == '#' {
				// lock
				lock := Lock{pins: counts, height: len(currToParse[0]) - 2}
				locks = append(locks, lock)
			} else {
				// key
				key := Key{pins: counts, height: len(currToParse[0]) - 2}
				keys = append(keys, key)
			}

			currToParse = []string{}
		} else {
			currToParse = append(currToParse, line)
		}
	}

	return locks, keys
}

func checkOverlap(lock Lock, key Key) bool {
	if len(lock.pins) != len(key.pins) {
		return false
	}
	for i := 0; i < len(lock.pins); i++ {
		if lock.pins[i]+key.pins[i] > lock.height {
			return true
		}
	}
	return false
}

func countNotOverlapping(locks []Lock, keys []Key) int {
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			// fmt.Printf("Checking lock: %v, key: %v --> Overlapping %v\n", lock, key, checkOverlap(lock, key))
			if !checkOverlap(lock, key) {
				count++
			}
		}
	}
	return count
}

func transposeStr(matrix []string) []string {
	transposed := []string{}
	for i := 0; i < len(matrix[0]); i++ {
		transposed = append(transposed, "")
		for j := 0; j < len(matrix); j++ {
			transposed[i] += string(matrix[j][i])
		}
	}
	return transposed
}

func main() {
	// Open the file
	file, err := os.Open("25.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	lines := []string{}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	locks, keys := parseLocksKeys(lines)
	// fmt.Println(locks)
	// fmt.Println(keys)

	sum = countNotOverlapping(locks, keys)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
