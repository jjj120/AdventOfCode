package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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

func handleLine(line string) (p1 string, p2 string) {
	parts := strings.Split(line, "-")
	if parts == nil || len(parts) != 2 {
		fmt.Printf("Invalid line: %s\n", line)
		panic("Invalid line")
	}
	p1 = parts[0]
	p2 = parts[1]

	// fmt.Printf("p1: %s, p2: %s\n", p1, p2)
	return
}

func coundThreeWithStart(connections map[string][]string, startingWith string) int {
	visited := make(map[[3]string]bool)

	count := 0

	for p1, conns := range connections {
		for _, p2 := range conns {
			for _, p3 := range connections[p2] {
				if p1 == p2 || p1 == p3 || p2 == p3 {
					continue
				}

				p1Start := strings.HasPrefix(p1, startingWith)
				p2Start := strings.HasPrefix(p2, startingWith)
				p3Start := strings.HasPrefix(p3, startingWith)

				if !(p1Start || p2Start || p3Start) {
					continue
				}

				if !slices.Contains(conns, p3) {
					continue
				}

				// Sort the 3 points
				var points [3]string
				points[0] = p1
				points[1] = p2
				points[2] = p3
				points = sort3Slice(points)

				if visited[points] {
					continue
				}
				visited[points] = true
				count++
				// fmt.Printf("Found %s %s %s\n", p1, p2, p3)
			}
		}
	}
	return count
}

func sort3Slice(a [3]string) [3]string {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	// Open the file
	file, err := os.Open("23.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	connections := make(map[string][]string)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		p1, p2 := handleLine(line)
		connections[p1] = append(connections[p1], p2)
		connections[p2] = append(connections[p2], p1)
	}

	// fmt.Printf("Connections: %v\n", connections)
	sum = coundThreeWithStart(connections, "t")

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
