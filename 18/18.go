package main

import (
	"bufio"
	"fmt"
	"os"
)

const FIELD_SIZE = 71
const INT_MAX = 1<<31 - 1

type Coord struct {
	x, y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) Coord {
	coord := Coord{0, 0}
	fmt.Sscanf(line, "%d,%d", &coord.x, &coord.y)
	return coord
}

func findShortestPathLength(obstaclesLst []Coord, start, end Coord) int {
	obstacles := make(map[Coord]bool)
	for _, obstacle := range obstaclesLst {
		obstacles[obstacle] = true
	}

	// fmt.Printf("Length: %d\n", len(obstaclesLst))
	// printObstacles(obstacles)

	queue := make([]Coord, 0, 50)
	queue = append(queue, start)
	cost := make(map[Coord]int)

	for i := 0; i < FIELD_SIZE; i++ {
		for j := 0; j < FIELD_SIZE; j++ {
			cost[Coord{i, j}] = INT_MAX
		}
	}
	cost[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direction := range []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Coord{current.x + direction.x, current.y + direction.y}
			if next.x < 0 || next.x >= FIELD_SIZE || next.y < 0 || next.y >= FIELD_SIZE {
				continue
			}
			if _, ok := obstacles[next]; ok {
				continue
			}
			if _, ok := cost[next]; ok {
				if cost[next] <= cost[current]+1 {
					continue
				}
			}
			cost[next] = cost[current] + 1
			queue = append(queue, next)
		}
	}

	return cost[end]
}

func printObstacles(obstacles map[Coord]bool) {
	fmt.Printf("Obstacles:\n")
	for x := 0; x < FIELD_SIZE+2; x++ {
		ColorPrint(RGBtoAnsiEscapeString(0, 0, 127, true), UNICODE_BLOCK)
	}
	fmt.Println()
	for y := 0; y < FIELD_SIZE; y++ {
		ColorPrint(RGBtoAnsiEscapeString(0, 0, 127, true), UNICODE_BLOCK)
		for x := 0; x < FIELD_SIZE; x++ {
			if _, ok := obstacles[Coord{x, y}]; ok {
				ColorPrint(RGBtoAnsiEscapeString(0x66, 0x33, 0x99, true), UNICODE_BLOCK)
			} else {
				fmt.Print(".")
			}
		}
		ColorPrint(RGBtoAnsiEscapeString(0, 0, 127, true), UNICODE_BLOCK)
		fmt.Println()
	}
	for x := 0; x < FIELD_SIZE+2; x++ {
		ColorPrint(RGBtoAnsiEscapeString(0, 0, 127, true), UNICODE_BLOCK)
	}
	fmt.Println()
}

func main() {
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
	coords := make([]Coord, 0, 50)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		coords = append(coords, handleLine(line))
	}

	for i := 0; i < len(coords); i++ {
		if i%100 == 0 {
			fmt.Printf("Processing %d-th of %d obstacle\r", i, len(coords))
		}
		sum = findShortestPathLength(coords[:i+1], Coord{0, 0}, Coord{FIELD_SIZE - 1, FIELD_SIZE - 1})
		if sum == INT_MAX {
			fmt.Printf("No path found for %v obstacle                            \n", coords[i])
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
