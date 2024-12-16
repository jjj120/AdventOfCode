package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EMPTY = ' '
	WALL  = '#'
	START = 'S'
	END   = 'E'
)

const WALK_COST = 1
const TURN_COST = 1000
const INT_MAX = int(^uint(0) >> 1)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const ColorReset = "\033[0m"

const ColorBlack = "\033[30m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorMagenta = "\033[35m"
const ColorCyan = "\033[36m"
const ColorGray = "\033[37m"
const ColorWhite = "\033[97m"

const BGBlack = "\033[40m"
const BGRed = "\033[41m"
const BGGreen = "\033[42m"
const BGYellow = "\033[43m"
const BGBlue = "\033[44m"
const BGPurple = "\033[45m"
const BGCyan = "\033[46m"
const BGWhite = "\033[47m"

type Coord struct {
	x, y int
}

type RobotPosition struct {
	coord     Coord
	direction int
}

type LabyrinthCost [][][4]int

func (l LabyrinthCost) print() {
	for _, line := range l {
		for _, cell := range line {
			cellMin := min(cell[0], cell[1], cell[2], cell[3])
			cellVal := min(cellMin, 999999)
			fmt.Printf("%6d ", cellVal)
		}
		fmt.Println()
	}
}

func printCost(cost LabyrinthCost, lab Labyrinth) {
	for y, line := range cost {
		for x, costArr := range line {
			minCost := min(costArr[0], costArr[1], costArr[2], costArr[3])
			if lab[y][x] == WALL {
				fmt.Printf("%s%6s%s ", BGBlue, " WALL ", ColorReset)
				continue
			}
			if lab[y][x] == START {
				fmt.Printf("%s%6d%s ", BGGreen, minCost, ColorReset)
				continue
			}
			if lab[y][x] == END {
				fmt.Printf("%s%6d%s ", BGRed, minCost, ColorReset)
				continue
			}
			fmt.Printf("%6d ", minCost)
		}
		fmt.Println()
	}
}

func (l LabyrinthCost) getCost(coord RobotPosition) int {
	return l[coord.coord.y][coord.coord.x][coord.direction]
}

func (l LabyrinthCost) setCost(coord RobotPosition, cost int) {
	l[coord.coord.y][coord.coord.x][coord.direction] = cost
}

func (l LabyrinthCost) getMinCost(coord Coord) int {
	costs := l[coord.y][coord.x]
	return min(costs[0], costs[1], costs[2], costs[3])
}

type Labyrinth []string

func (l Labyrinth) print() {
	for _, line := range l {
		fmt.Println(line)
	}
}

func (l Labyrinth) isWall(coord Coord) bool {
	return l[coord.y][coord.x] == WALL
}

func (l Labyrinth) isEnd(coord Coord) bool {
	return l[coord.y][coord.x] == END
}

func (l Labyrinth) isStart(coord Coord) bool {
	return l[coord.y][coord.x] == START
}

func (l Labyrinth) isInside(coord Coord) bool {
	return coord.y >= 0 && coord.y < len(l) && coord.x >= 0 && coord.x < len(l[0])
}

func (l Labyrinth) getValue(coord Coord) byte {
	return l[coord.y][coord.x]
}

func (l Labyrinth) setValue(coord Coord, value byte) {
	l[coord.y] = l[coord.y][:coord.x] + string(value) + l[coord.y][coord.x+1:]
}

func (l Labyrinth) getNeighbours(coord Coord) []Coord {
	neighbours := make([]Coord, 0, 4)
	neighbours = append(neighbours, Coord{coord.x - 1, coord.y})
	neighbours = append(neighbours, Coord{coord.x + 1, coord.y})
	neighbours = append(neighbours, Coord{coord.x, coord.y - 1})
	neighbours = append(neighbours, Coord{coord.x, coord.y + 1})
	return neighbours
}

func (l Labyrinth) getValidNeighbours(coord Coord) []Coord {
	neighbours := l.getNeighbours(coord)

	validNeighbours := make([]Coord, 0, 4)
	for _, neighbour := range neighbours {
		if l.isInside(neighbour) && !l.isWall(neighbour) {
			validNeighbours = append(validNeighbours, neighbour)
		}
	}

	return validNeighbours
}

func (l Labyrinth) getValidNeighbourPos(coord Coord) []RobotPosition {
	neighbours := make([]RobotPosition, 0, 4)
	neighbours = append(neighbours, RobotPosition{Coord{coord.x - 1, coord.y}, LEFT})
	neighbours = append(neighbours, RobotPosition{Coord{coord.x + 1, coord.y}, RIGHT})
	neighbours = append(neighbours, RobotPosition{Coord{coord.x, coord.y - 1}, UP})
	neighbours = append(neighbours, RobotPosition{Coord{coord.x, coord.y + 1}, DOWN})

	validNeighbours := make([]RobotPosition, 0, 4)
	for _, neighbour := range neighbours {
		if l.isInside(neighbour.coord) && !l.isWall(neighbour.coord) {
			validNeighbours = append(validNeighbours, neighbour)
		}
	}

	return validNeighbours
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findStartEnd(labyrinth Labyrinth) (start, end Coord) {
	for y, line := range labyrinth {
		for x, char := range line {
			if char == START {
				start = Coord{x, y}
			} else if char == END {
				end = Coord{x, y}
			}
		}
	}

	return start, end
}

func findCheapestPath(labyrinth Labyrinth, start, end Coord) int {
	cellCosts := make(LabyrinthCost, len(labyrinth))
	for i := range cellCosts {
		cellCosts[i] = make([][4]int, len(labyrinth[i]))
	}

	// Initialize the costs to max
	for y, line := range labyrinth {
		for x := range line {
			for dir := 0; dir < 4; dir++ {
				cellCosts.setCost(RobotPosition{Coord{x, y}, dir}, INT_MAX)
			}
		}
	}

	cellCosts.setCost(RobotPosition{start, RIGHT}, 0) // the start cell has a cost of 0 to get to itself

	queue := make([]RobotPosition, 0, 200)
	queue = append(queue, RobotPosition{start, RIGHT})

	for len(queue) > 0 {
		current := queue[0]
		currPos := current.coord
		currDir := current.direction
		queue = queue[1:]

		// Get the cost of the current cell
		currentCost := cellCosts.getCost(current)

		neighbours := labyrinth.getValidNeighbourPos(currPos)
		for _, neighbour := range neighbours {
			// Calculate the cost of the neighbour
			neighbourCost := currentCost + WALK_COST

			turns := abs(neighbour.direction - currDir)

			if turns == 1 || turns == 3 {
				// one turn cw or ccw
				neighbourCost += TURN_COST
			}
			if turns == 2 {
				// two turns
				neighbourCost += 2 * TURN_COST
			}

			if neighbourCost < cellCosts.getCost(neighbour) {
				cellCosts.setCost(neighbour, neighbourCost)
				queue = append(queue, RobotPosition{neighbour.coord, neighbour.direction})
			}
		}
	}

	printCost(cellCosts, labyrinth)

	if cellCosts.getMinCost(end) == INT_MAX {
		panic("No path found")
	}

	return cellCosts.getMinCost(end)
}

func main() {
	// Open the file
	file, err := os.Open("16.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	labyrinth := make([]string, 0, 200)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		labyrinth = append(labyrinth, line)
	}

	start, end := findStartEnd(labyrinth)

	var sum = findCheapestPath(labyrinth, start, end)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
