package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
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

func findBestCombinationSingle(lines []string) int {
	maxBananas := 0
	maxSequence := [4]int{0, 0, 0, 0}
	for x1 := -9; x1 <= 9; x1++ {
		for x2 := -9; x2 <= 9; x2++ {
			for x3 := -9; x3 <= 9; x3++ {
				for x4 := -9; x4 <= 9; x4++ {
					combination := [4]int{x1, x2, x3, x4}
					if !checkValidComb(combination) {
						continue
					}
					fmt.Printf("Trying combination: %v       \r", combination)
					sum := 0
					for _, line := range lines {
						sum += handleLine(line, combination)
					}
					if sum > maxBananas {
						maxBananas = sum
						maxSequence = combination
					}
				}
			}
		}
	}
	fmt.Printf("Max sequence: %v with bananas %d\n", maxSequence, maxBananas)
	return maxBananas
}

func findBestCombinationParallel(lines []string) int {
	maxBananas := 0
	var maxSequence [4]int
	var mutex sync.Mutex
	waitgroup := sync.WaitGroup{}

	for x1 := -9; x1 <= 9; x1++ {
		waitgroup.Add(1)
		go bruteForceLastThree(x1, lines, &mutex, &maxBananas, &maxSequence, &waitgroup)
	}

	waitgroup.Wait()

	fmt.Printf("Max sequence: %v with bananas %d\n", maxSequence, maxBananas)

	return maxBananas
}

func bruteForceLastThree(x1 int, lines []string, mutex *sync.Mutex, maxBananas *int, maxSequence *[4]int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for x2 := -9; x2 <= 9; x2++ {
		for x3 := -9; x3 <= 9; x3++ {
			for x4 := -9; x4 <= 9; x4++ {
				combination := [4]int{x1, x2, x3, x4}
				if !checkValidComb(combination) {
					continue
				}
				sum := 0
				for _, line := range lines {
					sum += handleLine(line, combination)
				}
				mutex.Lock()
				if sum > *maxBananas {
					*maxBananas = sum
					*maxSequence = combination
				}
				mutex.Unlock()
			}
		}
	}
}

func checkValidComb(combination [4]int) bool {
	number := 0
	for i := 0; i < 4; i++ {
		number += combination[i]
		if number > 9 || number < -9 {
			return false
		}
	}
	return true
}

func handleLine(line string, combinationToSearch [4]int) int {
	var secretNum int
	_, err := fmt.Sscanf(line, "%d", &secretNum)
	check(err)

	const numIterations = 2000

	previousSequence := [4]int{0, 0, 0, 0}
	prevOnes := getOnes(secretNum)
	for i := 0; i < 4; i++ {
		secretNum = calcNextSecret(secretNum)
		currOnes := getOnes(secretNum)
		previousSequence[i] = (currOnes - prevOnes)
		prevOnes = currOnes
	}

	for i := 0; i < numIterations-4; i++ {
		secretNum = calcNextSecret(secretNum)
		currOnes := getOnes(secretNum)

		previousSequence[0] = previousSequence[1]
		previousSequence[1] = previousSequence[2]
		previousSequence[2] = previousSequence[3]
		previousSequence[3] = (currOnes - prevOnes)

		if previousSequence == combinationToSearch {
			return currOnes
		}

		prevOnes = currOnes
	}

	// fmt.Printf("%s: %d\n", line, secretNum)

	return 0
}

func calcNextSecret(number int) int {
	number = pruneSecret(mixSecrets(number, number<<6))
	number = pruneSecret(mixSecrets(number, number>>5))
	number = pruneSecret(mixSecrets(number, number<<11))
	return number
}

func mixSecrets(secretNum, value int) int {
	return secretNum ^ value
}

func pruneSecret(secretNum int) int {
	return secretNum & 0xFFFFFF // mod 16777216
}

func getOnes(number int) int {
	return number % 10
}

func main() {
	// Open the file
	file, err := os.Open("22.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	lines := make([]string, 0)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	sum = findBestCombinationParallel(lines)
	// sum = findBestCombinationSingle(lines)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
