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

type Matrix3x3 [3][3]bool

func (m Matrix3x3) AndAll() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !m[i][j] {
				return false
			}
		}
	}
	return true
}

func (m Matrix3x3) At(x, y int) bool {
	assert(x >= 0 && x < 3, "x out of bounds for At")
	assert(y >= 0 && y < 3, "y out of bounds for At")
	return m[y][x]
}

func (m Matrix3x3) AtCentered(x, y int) bool {
	assert(x >= -1 && x <= 1, "x out of bounds for AtCentered")
	assert(y >= -1 && y <= 1, "y out of bounds for AtCentered")
	return m[y+1][x+1]
}

type GardenMap struct {
	width    int
	height   int
	plantMap map[Coord]string
}

func (g GardenMap) CheckBounds(c Coord) bool {
	return c.x >= 0 && c.x < g.width && c.y >= 0 && c.y < g.height
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
			thisArea := map[Coord]bool{}
			a, _ := calcAreaPerimeterFromPosition(currCoord, gardenMap, &alreadyVisited, &thisArea)
			p := calculateAreaPerimeter(thisArea)
			debugPrintf("From %v (%s): Area: %d, Perimeter: %d -> Cost: %d\n", currCoord, string(gardenMap.plantMap[currCoord]), a, p, a*p)
			cost += a * p
		}
	}

	return cost
}

func calculateAreaPerimeter(area map[Coord]bool) int {
	debugPrintf("Area: %v\n", area)

	corners := map[Coord]int{}

	for point := range area {
		currCorners := countCorner(point, area)
		if currCorners > 0 {
			corners[point] = currCorners
		}
	}

	sum := 0
	for _, v := range corners {
		sum += v
	}

	debugPrintf("Corners: %v\n", corners)
	return sum
}

func countCorner(point Coord, area map[Coord]bool) int {
	cornersMatrix := Matrix3x3{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			currCoord := Coord{x: point.x + x, y: point.y + y}
			if _, ok := area[currCoord]; ok {
				cornersMatrix[y+1][x+1] = true
			}
		}
	}

	debugPrintf("Corners matrix for %v: %v\n", point, cornersMatrix)

	if cornersMatrix.AndAll() {
		// Grid full, point in the middle
		debugPrintf("Grid full\n")
		return 0
	}

	corners := 0
	dirs := []Coord{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	// check for outward corners
	for i := 0; i < 4; i++ {
		dir1 := dirs[i]
		dir2 := dirs[(i+1)%4]
		if !cornersMatrix.AtCentered(dir1.x, dir1.y) && !cornersMatrix.AtCentered(dir2.x, dir2.y) {
			corners++
		}
	}

	dirs = []Coord{
		{x: -1, y: -1},
		{x: 1, y: -1},
		{x: 1, y: 1},
		{x: -1, y: 1},
	}

	// check for inward corners
	for _, dir := range dirs {
		dirUpDown := Coord{x: 0, y: dir.y}
		dirLeftRight := Coord{x: dir.x, y: 0}
		if !cornersMatrix.AtCentered(dir.x, dir.y) && cornersMatrix.AtCentered(dirUpDown.x, dirUpDown.y) && cornersMatrix.AtCentered(dirLeftRight.x, dirLeftRight.y) {
			corners++
		}
	}

	return corners
}

func calcAreaPerimeterFromPosition(currCoord Coord, gardenMap GardenMap, alreadyVisited *map[Coord]bool, areaCoords *map[Coord]bool) (int, int) {
	// returns area, perimeter
	if _, ok := (*alreadyVisited)[currCoord]; ok {
		// debugPringf("Already visited %v\n", currCoord)
		fmt.Printf("Already visited %v\n", currCoord)
		return 0, 0
	}

	(*alreadyVisited)[currCoord] = true
	(*areaCoords)[currCoord] = true

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
					a, p := calcAreaPerimeterFromPosition(adj, gardenMap, alreadyVisited, areaCoords)
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
