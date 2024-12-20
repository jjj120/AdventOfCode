package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type Coord struct {
	x, y int
}

func getStartEnd(labyrinth []string) (startpoint Coord, endpoint Coord) {
	for y, line := range labyrinth {
		for x, char := range line {
			if char == 'S' {
				startpoint = Coord{x, y}
			} else if char == 'E' {
				endpoint = Coord{x, y}
			}
		}
	}
	return
}

func calcCosts(labyrinth []string) map[Coord]int {
	start, end := getStartEnd(labyrinth)

	costs := make(map[Coord]int)
	costs[start] = 0

	curr := start

	for {
		// Check if we reached the end
		if curr == end {
			break
		}

		for _, dir := range []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Coord{curr.x + dir.x, curr.y + dir.y}
			if labyrinth[next.y][next.x] == '#' {
				continue
			}

			_, ok := costs[next]
			if ok {
				continue
			}
			costs[next] = costs[curr] + 1
			curr = next
		}
	}

	return costs
}

func checkCheats(labyrinth []string, costs map[Coord]int, minSave int) int {
	cheats := 0
	costsMap := map[int]int{}

	for y, line := range labyrinth {
		for x, char := range line {
			if char != '#' {
				cheatCosts := checkCheatsFromPoint(costs, Coord{x, y})
				for save, count := range cheatCosts {
					costsMap[save] += count
				}
			}
		}
	}

	saves := make([]int, len(costsMap))
	i := 0
	for k := range costsMap {
		saves[i] = k
		i++
	}
	slices.Sort(saves)

	for _, save := range saves {
		count := costsMap[save]
		// fmt.Printf("There are %d cheats that save %d picoseconds\n", count, save)
		if save >= minSave {
			cheats += count
		}
	}

	return cheats
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkCheatsFromPoint(costs map[Coord]int, point Coord) map[int]int {
	saveMap := map[int]int{}

	maxCheatCost := 20

	for dx := -maxCheatCost; dx <= maxCheatCost; dx++ {
		for dy := -maxCheatCost; dy <= maxCheatCost; dy++ {
			shortcutCost := abs(dx) + abs(dy)

			if shortcutCost > maxCheatCost || shortcutCost == 0 {
				continue
			}

			neighbour := Coord{point.x + dx, point.y + dy}

			_, ok := costs[neighbour]
			if !ok {
				continue
			}

			if costs[neighbour] > costs[point]+shortcutCost {
				// fmt.Printf("Cheat found at %d,%d to %d, %d with cost saving %d\n", point.x, point.y, neighbour.x, neighbour.y, costs[neighbour]-costs[point]-shortcutCost)
				saveMap[costs[neighbour]-costs[point]-shortcutCost]++
			}
		}
	}
	return saveMap
}

func printLabyrinthWithCosts(labyrinth []string, costs map[Coord]int) {
	maxCost := 0
	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}

	fmt.Print("   ")
	for x := 0; x < len(labyrinth[0]); x++ {
		fmt.Printf("%d", x%10)
	}
	fmt.Println()

	for y, line := range labyrinth {
		fmt.Printf("%2d ", y)
		for x, char := range line {
			if char == '#' {
				if _, ok := costs[Coord{x, y}]; ok {
					panic("Costs should not be calculated for walls")
				}
				ColorPrint("\033[30m", "█")
			} else {
				cost, ok := costs[Coord{x, y}]
				if ok {
					r, g, b := HSLtoRGB(float64(cost)/float64(maxCost)*360, 1, 0.5)
					ColorPrint(RGBtoAnsiEscapeString(r, g, b, true), "█")
				} else {
					fmt.Print(" ") // should never happen
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	// Open the file
	file, err := os.Open("20.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	labyrinth := []string{}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		labyrinth = append(labyrinth, line)
	}

	costs := calcCosts(labyrinth)
	// printLabyrinthWithCosts(labyrinth, costs)

	var sum = 0
	const minSave = 100
	sum = checkCheats(labyrinth, costs, minSave)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
