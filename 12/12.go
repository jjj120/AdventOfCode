package main

import (
	"bufio"
	"fmt"
	"os"
)

const DEBUG = false

type Coord struct {
	x int
	y int
}

func (c Coord) getAdjacent() []Coord {
	return []Coord{
		{x: c.x - 1, y: c.y},
		{x: c.x + 1, y: c.y},
		{x: c.x, y: c.y - 1},
		{x: c.x, y: c.y + 1},
	}
}

type GardenMap struct {
	width    int
	height   int
	plantMap map[Coord]string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func debugPrintf(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

func handleLine(line string, lineIndex int) map[Coord]string {
	toReturn := map[Coord]string{}
	for i, c := range line {
		toReturn[Coord{x: i, y: lineIndex}] = string(c)
	}

	return toReturn
}

func calcFenceCost(gardenMap GardenMap) int {
	alreadyVisited := map[Coord]bool{}
	cost := 0

	for x := 0; x < gardenMap.width; x++ {
		for y := 0; y < gardenMap.height; y++ {
			currCoord := Coord{x: x, y: y}
			if alreadyVisited[currCoord] {
				continue
			}
			a, p := calcAreaPerimeterFromPosition(currCoord, gardenMap, &alreadyVisited)
			debugPrintf("From %v (%s): Area: %d, Perimeter: %d\n", currCoord, string(gardenMap.plantMap[currCoord]), a, p)
			cost += a * p
		}
	}

	return cost
}

func calcAreaPerimeterFromPosition(currCoord Coord, gardenMap GardenMap, alreadyVisited *map[Coord]bool) (int, int) {
	// returns area, perimeter
	if _, ok := (*alreadyVisited)[currCoord]; ok {
		// debugPringf("Already visited %v\n", currCoord)
		fmt.Printf("Already visited %v\n", currCoord)
		return 0, 0
	}

	(*alreadyVisited)[currCoord] = true

	if _, ok := gardenMap.plantMap[currCoord]; !ok {
		debugPrintf("No plant at %v\n", currCoord)
		return 0, 0
	}

	adjacent := currCoord.getAdjacent()
	currPlant := gardenMap.plantMap[currCoord]
	area := 1
	perimeter := 0

	debugPrintf("Checking %v\n", currCoord)
	for _, adj := range adjacent {
		if plant, ok := gardenMap.plantMap[adj]; ok {
			if _, ok := (*alreadyVisited)[adj]; !ok {
				// New plant to check
				if plant != currPlant {
					debugPrintf("Different plant at %v than at %v\n", adj, currCoord)
					perimeter++
				} else {
					debugPrintf("Same plant at %v as at %v\n", adj, currCoord)
					a, p := calcAreaPerimeterFromPosition(adj, gardenMap, alreadyVisited)
					area += a
					perimeter += p
				}
			} else {
				// Already visited
				debugPrintf("Already visited %v but could be from prev field\n", adj)
				if plant != currPlant {
					debugPrintf("Different plant at %v than at %v\n", adj, currCoord)
					perimeter++
				}
			}
		} else {
			// No plant (perimeter)
			debugPrintf("No plant at %v\n", adj)
			perimeter++
		}
	}

	return area, perimeter
}

func printGardenMap(gardenMap GardenMap) {
	for y := 0; y < gardenMap.height; y++ {
		for x := 0; x < gardenMap.width; x++ {
			fmt.Printf("%s", gardenMap.plantMap[Coord{x: x, y: y}])
		}
		fmt.Println()
	}
}

func main() {
	// Open the file
	file, err := os.Open("12.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	gardenMap := GardenMap{
		width:    0,
		height:   0,
		plantMap: map[Coord]string{},
	}
	lineIndex := 0

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		gardenMap.width = max(gardenMap.width, len(line))
		for k, v := range handleLine(line, lineIndex) {
			gardenMap.plantMap[k] = v
		}
		lineIndex++
		gardenMap.height = lineIndex
	}

	// printGardenMap(gardenMap)

	var sum = calcFenceCost(gardenMap)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	assert(sum != 1370096, "Sum should not be 1370096")
	assert(sum != 6072738, "Sum should not be 6072738")
	fmt.Printf("Sum: %d\n", sum)
}
