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

type point struct {
	x, y int
}

func handleLine(line string) []int {
	mapLine := []int{}
	for _, c := range line {
		mapLine = append(mapLine, int(c)-'0')
	}
	return mapLine
}

func countTrailheads(topoMap [][]int) int {
	sum := 0
	for y := 0; y < len(topoMap); y++ {
		for x := 0; x < len(topoMap[0]); x++ {
			if topoMap[y][x] == 0 {
				// found a trail start
				reachablesNew := getReachable(topoMap, x, y)
				sum += len(reachablesNew)
			}
		}
	}

	return sum
}

func getReachable(topoMap [][]int, x, y int) map[point]bool {
	if x < 0 || y < 0 || x >= len(topoMap[0]) || y >= len(topoMap) {
		return map[point]bool{}
	}
	if topoMap[y][x] == 9 {
		return map[point]bool{{x, y}: true}
	}
	toReturn := map[point]bool{}

	if y > 0 && topoMap[y-1][x] == topoMap[y][x]+1 {
		reachables := getReachable(topoMap, x, y-1)
		for k, v := range reachables {
			toReturn[k] = v
		}
	}
	if y < len(topoMap)-1 && topoMap[y+1][x] == topoMap[y][x]+1 {
		reachables := getReachable(topoMap, x, y+1)
		for k, v := range reachables {
			toReturn[k] = v
		}
	}
	if x > 0 && topoMap[y][x-1] == topoMap[y][x]+1 {
		reachables := getReachable(topoMap, x-1, y)
		for k, v := range reachables {
			toReturn[k] = v
		}
	}
	if x < len(topoMap[0])-1 && topoMap[y][x+1] == topoMap[y][x]+1 {
		reachables := getReachable(topoMap, x+1, y)
		for k, v := range reachables {
			toReturn[k] = v
		}
	}

	return toReturn
}

func main() {
	// Open the file
	file, err := os.Open("10.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	topoMap := [][]int{}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		topoMap = append(topoMap, handleLine(line))
	}

	sum = countTrailheads(topoMap)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
