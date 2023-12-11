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

type Galaxy struct {
	x, y       int
	newX, newY int
}

func findGalaxies(data []string) []Galaxy {
	var galaxies []Galaxy
	for y, line := range data {
		for x, char := range line {
			if char == '#' {
				var galaxy Galaxy
				galaxy.x = x
				galaxy.y = y
				galaxy.newX = -1
				galaxy.newY = -1
				galaxies = append(galaxies, galaxy)
			}
		}
	}
	return galaxies
}

func lineEmpty(line string) bool {
	for _, char := range line {
		if char != '.' {
			return false
		}
	}
	return true
}

func expandRows(data []string, galaxies []Galaxy) []Galaxy {
	expandAmount := make([]int, len(data))

	for i, line := range data {
		if i == 0 {
			if lineEmpty(line) {
				expandAmount[i] = 1
			} else {
				expandAmount[i] = 0
			}
			continue
		}
		if lineEmpty(line) {
			expandAmount[i] = expandAmount[i-1] + 1
		} else {
			expandAmount[i] = expandAmount[i-1] + 0
		}
	}

	for i, galaxy := range galaxies {
		galaxies[i].newY = galaxy.y + expandAmount[galaxy.y]
	}
	return galaxies
}

func colEmpty(data []string, colIndex int) bool {
	for _, line := range data {
		if line[colIndex] != '.' {
			return false
		}
	}
	return true
}

func expandCols(data []string, galaxies []Galaxy) []Galaxy {
	expandAmount := make([]int, len(data))

	for i := range data[0] {
		if i == 0 {
			if colEmpty(data, i) {
				expandAmount[i] = 1
			} else {
				expandAmount[i] = 0
			}
			continue
		}
		if colEmpty(data, i) {
			expandAmount[i] = expandAmount[i-1] + 1
		} else {
			expandAmount[i] = expandAmount[i-1] + 0
		}
	}

	for i, galaxy := range galaxies {
		galaxies[i].newX = galaxy.x + expandAmount[galaxy.x]
	}
	return galaxies
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getShortestPath(galaxy1 Galaxy, galaxy2 Galaxy) int {
	return abs(galaxy1.newX-galaxy2.newX) + abs(galaxy1.newY-galaxy2.newY)
}

func getShortestPaths(galaxies []Galaxy) int {
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < i; j++ {
			sum += getShortestPath(galaxies[i], galaxies[j])
		}
	}
	return sum
}

func handleData(data []string) int {
	galaxies := findGalaxies(data)
	galaxies = expandRows(data, galaxies)
	galaxies = expandCols(data, galaxies)

	return getShortestPaths(galaxies)
}

func main() {
	// Open the file
	file, err := os.Open("11.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var data []string
	// Iterate through each line
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	sum := handleData(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
