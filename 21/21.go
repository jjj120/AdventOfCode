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

func parseData(data []string) (map[complex128]bool, complex128) {
	garden := make(map[complex128]bool)
	var start complex128 = -1
	for y, line := range data {
		for x, c := range line {
			if c != '#' {
				garden[complex(float64(x), float64(y))] = true
			}
			if c == 'S' {
				start = complex(float64(x), float64(y))
			}
		}
	}

	if start == -1 {
		panic("No start found!")
	}

	return garden, start
}

func complexMod(num complex128, mod int) complex128 {
	if float64(int(real(num))) != real(num) && float64(int(imag(num))) != imag(num) {
		fmt.Println(num, real(num), imag(num))
		panic("Complex number not integer!")
	}
	return complex(float64((int(real(num))+10*mod)%mod), float64((int(imag(num))+10*mod)%mod))
}

func calculateNumEnds(garden map[complex128]bool, start complex128, numIterations int, maxSize int) int {
	queue := make(map[complex128]bool)
	queue[start] = true

	done := make([]int, 0)

	for i := 0; i < 3*maxSize; i++ {
		if (i % maxSize) == (maxSize-1)/2 {
			fmt.Println(i, len(queue))
			done = append(done, len(queue))
		}
		if len(done) == 3 {
			break
		}

		newQueue := make(map[complex128]bool)

		for _, dir := range []complex128{1, -1, 1i, -1i} {
			for point := range queue {
				if _, ok := garden[complexMod(point+dir, maxSize)]; ok {
					newQueue[point+dir] = true
				}
			}
		}
		queue = newQueue
	}

	quadraticFunction := func(n, a, b, c int) int {
		return a + n*(b-a+((n-1)*(c-2*b+a)/2))
	}

	fmt.Println(numIterations/maxSize, done[0], done[1], done[2])
	return quadraticFunction(numIterations/maxSize, done[0], done[1], done[2])
}

func assert(a bool, msg string) {
	if !a {
		panic("Assertion failed: " + msg)
		// fmt.Println("Assertion failed: " + msg)
	} else {
		fmt.Println("Assertion succeeded: " + msg)
	}
}

const PRINT = true

func main() {
	// Open the file
	file, err := os.Open("21.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	var gardenInput = []string{}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		gardenInput = append(gardenInput, line)
	}

	garden, start := parseData(gardenInput)
	maxSize := len(gardenInput)
	fmt.Println("Max size:", maxSize)
	fmt.Println("Start:", start)
	fmt.Println("Garden size:", len(garden))

	// // Tests for 21.ex
	// assert(calculateNumEnds(garden, start, 6, maxSize) == 16, "Test 1: "+fmt.Sprint((calculateNumEnds(garden, start, 6, maxSize)))+" = "+fmt.Sprint(16))
	// assert(calculateNumEnds(garden, start, 10, maxSize) == 50, "Test 2: "+fmt.Sprint((calculateNumEnds(garden, start, 10, maxSize)))+" = "+fmt.Sprint(50))
	// assert(calculateNumEnds(garden, start, 50, maxSize) == 1594, "Test 3: "+fmt.Sprint((calculateNumEnds(garden, start, 50, maxSize)))+" = "+fmt.Sprint(1594))
	// assert(calculateNumEnds(garden, start, 100, maxSize) == 6536, "Test 4: "+fmt.Sprint((calculateNumEnds(garden, start, 100, maxSize)))+" = "+fmt.Sprint(6536))
	// assert(calculateNumEnds(garden, start, 500, maxSize) == 167004, "Test 5: "+fmt.Sprint((calculateNumEnds(garden, start, 500, maxSize)))+" = "+fmt.Sprint(167004))
	// assert(calculateNumEnds(garden, start, 1000, maxSize) == 668697, "Test 6: "+fmt.Sprint((calculateNumEnds(garden, start, 1000, maxSize)))+" = "+fmt.Sprint(668697))
	// assert(calculateNumEnds(garden, start, 5000, maxSize) == 16733044, "Test 7: "+fmt.Sprint((calculateNumEnds(garden, start, 5000, maxSize)))+" = "+fmt.Sprint(16733044))

	sum = calculateNumEnds(garden, start, 26501365, maxSize)

	assert(sum == 599763113936220, "Right solution: 599763113936220, got "+fmt.Sprint(sum))

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
