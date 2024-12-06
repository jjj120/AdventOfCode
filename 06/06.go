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

type posDir struct {
	x   int
	y   int
	dir string
}

func countLoops(obstacles [][]bool, start posDir) int {
	loopPossibilities := 0

	for obstacleY := 0; obstacleY < len(obstacles[0]); obstacleY++ {
		for obstacleX := 0; obstacleX < len(obstacles); obstacleX++ {
			if obstacleX == start.x && obstacleY == start.y {
				continue
			}
			if obstacles[obstacleY][obstacleX] {
				continue
			}

			fmt.Printf("Checking %d, %d         \r", obstacleX, obstacleY)

			obstacles[obstacleY][obstacleX] = true
			_, loop := countVisited(obstacles, start)
			if loop {
				loopPossibilities++
			}
			obstacles[obstacleY][obstacleX] = false
		}
	}
	fmt.Println("Finished checking")
	return loopPossibilities
}

func countVisited(obstacles [][]bool, start posDir) (int, bool) {
	currPos := start

	visited := make([][]bool, len(obstacles))
	for i := range visited {
		visited[i] = make([]bool, len(obstacles[0]))
	}

	memory := make(map[posDir]bool)

	for checkBounds(obstacles, currPos.x, currPos.y) {
		if memory[currPos] {
			return countTrue(visited), true
		}
		memory[currPos] = true
		visited[currPos.y][currPos.x] = true
		currPos.x, currPos.y, currPos.dir = makeStep(obstacles, currPos.x, currPos.y, currPos.dir)
	}

	// printGameBoard(obstacles, visited, start)
	return countTrue(visited), false
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

func printGameBoard(obstacles, visited [][]bool, start posDir) {
	for i, row := range obstacles {
		for j, v := range row {
			if i == start.x && j == start.y {
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
	var start = posDir{0, 0, "up"}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		obstacles = append(obstacles, make([]bool, len(line)))
		for i, c := range line {
			if c == '#' {
				obstacles[len(obstacles)-1][i] = true
			} else if c == '^' {
				start.x = i
				start.y = len(obstacles) - 1
			}
		}
	}

	fmt.Printf("Start: %d, %d\n", start.x, start.y)
	fmt.Printf("Gameboard size: %d, %d\n", len(obstacles[0]), len(obstacles))

	var sum int
	sum, _ = countVisited(obstacles, start) // Part 1
	fmt.Printf("Sum part 1: %d\n", sum)

	sum = countLoops(obstacles, start) // Part 2
	fmt.Printf("Sum part 2: %d\n", sum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
