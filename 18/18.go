package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string, lights *[100][100]bool, y int) int {
	for x, c := range line {
		if c == '#' {
			lights[x][y] = true
		} else {
			lights[x][y] = false
		}
	}
	return 0
}

func setCorners(lights *[100][100]bool) {
	(*lights)[0][0] = true
	(*lights)[0][99] = true
	(*lights)[99][0] = true
	(*lights)[99][99] = true
}

func countNeighbors(lights *[100][100]bool, x int, y int) int {
	dirs := [8][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	var count = 0
	for _, dir := range dirs {
		newX := x + dir[0]
		newY := y + dir[1]
		if newX >= 0 && newX < 100 && newY >= 0 && newY < 100 {
			if (*lights)[newY][newX] {
				count++
			}
		}
	}

	return count
}

func makeStep(lights *[100][100]bool) {
	newField := [100][100]bool{}
	copy(newField[:], lights[:])

	for y, line := range *lights {
		for x, light := range line {
			neigh := countNeighbors(lights, x, y)
			if light {
				if neigh < 2 || neigh > 3 {
					(newField)[y][x] = false
				}
			} else {
				if neigh == 3 {
					(newField)[y][x] = true
				}
			}
		}
	}

	*lights = newField
}

func printLights(lights *[100][100]bool) {
	if !PRINT {
		return
	}
	tm.MoveCursor(1, 1)
	for _, line := range lights {
		for _, light := range line {
			if light {
				fmt.Print("\033[93m██\033[0m")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	tm.Flush()
}

func countLights(lights *[100][100]bool) int {
	var count = 0
	for _, line := range lights {
		for _, light := range line {
			if light {
				count++
			}
		}
	}
	return count
}

const PRINT = true

func main() {
	if PRINT {
		tm.Clear()
	}
	// Open the file
	file, err := os.Open("18.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	var lights [100][100]bool
	y := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line, &lights, y)
		y++
	}

	for i := 0; i < 100; i++ {
		makeStep(&lights)
		setCorners(&lights)
		printLights(&lights)
		time.Sleep(10 * time.Millisecond)
	}
	// printLights(&lights)

	sum = countLights(&lights)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
