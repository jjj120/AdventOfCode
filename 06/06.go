package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func setLight(lights *[1000][1000]int, x1, y1, x2, y2 int, lightStatus bool) {
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		for y := min(y1, y2); y <= max(y1, y2); y++ {
			if lightStatus {
				(*lights)[y][x]++
			} else {
				(*lights)[y][x] = max((*lights)[y][x]-1, 0)
			}
		}
	}
}

func toggleLight(lights *[1000][1000]int, x1, y1, x2, y2 int) {
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		for y := min(y1, y2); y <= max(y1, y2); y++ {
			(*lights)[y][x] += 2
		}
	}
}

func strToInt(str string) int {
	num, err := strconv.ParseInt(str, 10, 0)
	check(err)
	return int(num)
}

func countOn(lights *[1000][1000]int) int {
	lightLevel := 0
	for _, line := range *lights {
		for _, light := range line {
			lightLevel += light
		}
	}
	return lightLevel
}

func printLights(lights *[1000][1000]int) int {
	lit := 0
	for _, line := range *lights {
		for _, light := range line {
			fmt.Printf("%d", light)
		}
	}
	return lit
}

func handleLine(lights *[1000][1000]int, line string) int {
	line = strings.ReplaceAll(line, "turn ", "")
	instructions := strings.Split(line, " ")

	p1 := strings.Split(instructions[1], ",")
	p2 := strings.Split(instructions[3], ",")

	x1 := strToInt(p1[0])
	y1 := strToInt(p1[1])
	x2 := strToInt(p2[0])
	y2 := strToInt(p2[1])

	if strings.Contains(line, "on") {
		setLight(lights, x1, y1, x2, y2, true)
	} else if strings.Contains(line, "off") {
		setLight(lights, x1, y1, x2, y2, false)
	} else {
		toggleLight(lights, x1, y1, x2, y2)
	}

	return 0
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

	var lights [1000][1000]int

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(&lights, line)
	}

	sum := countOn(&lights)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
