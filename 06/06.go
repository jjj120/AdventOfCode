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

func countVisited(obstacles [][]bool, start []int) int {
	currX := start[0]
	currY := start[1]
	currDir := "up"

	fmt.Printf("Start at: %d, %d\n", currX, currY)

	visited := make([][]bool, len(obstacles))
	for i := range visited {
		visited[i] = make([]bool, len(obstacles[0]))
	}

	for checkBounds(obstacles, currX, currY) {
		visited[currY][currX] = true
		currX, currY, currDir = makeStep(obstacles, currX, currY, currDir)
	}

	printGameBoard(obstacles, visited, start)
	return countTrue(visited)
}

func countTrue(visited [][]bool) int {
	count := 0
	for _, row := range visited {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return count
}

func printGameBoard(obstacles, visited [][]bool, start []int) {
	for i, row := range obstacles {
		for j, v := range row {
			if i == start[0] && j == start[1] {
				fmt.Print("^")
			} else if v {
				fmt.Print("#")
			} else if visited[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func checkBounds(obstacles [][]bool, x, y int) bool {
	return x >= 0 && x < len(obstacles[0]) && y >= 0 && y < len(obstacles)
}

func makeStep(obstacles [][]bool, x, y int, dir string) (int, int, string) {
	// dir strings: up, down, left, right
	switch dir {
	case "up":
		if y == 0 {
			return x, y - 1, "up"
		}
		if obstacles[y-1][x] {
			// turn right
			return x, y, "right"
		}
		return x, y - 1, "up"

	case "down":
		if y == len(obstacles)-1 {
			return x, y + 1, "down"
		}
		if obstacles[y+1][x] {
			// turn right
			return x, y, "left"
		}
		return x, y + 1, "down"

	case "left":
		if x == 0 {
			return x - 1, y, "left"
		}
		if obstacles[y][x-1] {
			// turn right
			return x, y, "up"
		}
		return x - 1, y, "left"

	case "right":
		if x == len(obstacles[0])-1 {
			return x + 1, y, "right"
		}
		if obstacles[y][x+1] {
			// turn right
			return x, y, "down"
		}
		return x + 1, y, "right"

	}

	fmt.Printf("Error: invalid direction %s at %d, %d\n", dir, x, y)
	return x, y, dir
}

func main() {
	// Open the file
	file, err := os.Open("06.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var obstacles = make([][]bool, 0)
	var start = make([]int, 2)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		obstacles = append(obstacles, make([]bool, len(line)))
		for i, c := range line {
			if c == '#' {
				obstacles[len(obstacles)-1][i] = true
			} else if c == '^' {
				start[0] = i
				start[1] = len(obstacles) - 1
			}
		}
	}

	sum := countVisited(obstacles, start)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
