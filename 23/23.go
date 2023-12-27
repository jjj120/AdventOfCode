package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/dominikbraun/graph"
)

const (
	PATH = iota
	FOREST
	SLOPE_UP
	SLOPE_DOWN
	SLOPE_RIGHT
	SLOPE_LEFT
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) []int {
	currMap := make([]int, 0)
	for _, char := range line {
		switch char {
		case '.':
			currMap = append(currMap, PATH)
		case '#':
			currMap = append(currMap, FOREST)
		case '^':
			currMap = append(currMap, SLOPE_UP)
		case 'v':
			currMap = append(currMap, SLOPE_DOWN)
		case '>':
			currMap = append(currMap, SLOPE_RIGHT)
		case '<':
			currMap = append(currMap, SLOPE_LEFT)
		}
	}
	return currMap
}

// [y, x]
func findStartPoint(hikingMap [][]int) []int {
	for j, col := range hikingMap[0] {
		if col == PATH {
			return []int{0, j}
		}
	}
	return []int{-1, -1}
}

// [y, x]
func findEndPoint(hikingMap [][]int) []int {
	for j, col := range hikingMap[len(hikingMap)-1] {
		if col == PATH {
			return []int{len(hikingMap) - 1, j}
		}
	}
	return []int{-1, -1}
}

// point: [y, x]
func isNode(hikingMap [][]int, point []int) bool {
	if point[0] < 0 || point[0] >= len(hikingMap) || point[1] < 0 || point[1] >= len(hikingMap[0]) {
		return false
	}

	startPoint := findStartPoint(hikingMap)
	endPoint := findEndPoint(hikingMap)
	if (point[0] == startPoint[0] && point[1] == startPoint[1]) || (point[0] == endPoint[0] && point[1] == endPoint[1]) {
		return true
	}

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	counter := 0
	for _, dir := range dirs {
		newPoint := []int{point[0] + dir[0], point[1] + dir[1]}
		if newPoint[0] < 0 || newPoint[0] >= len(hikingMap) || newPoint[1] < 0 || newPoint[1] >= len(hikingMap[0]) {
			continue
		}

		slopes := []int{SLOPE_UP, SLOPE_DOWN, SLOPE_RIGHT, SLOPE_LEFT}
		if slices.Contains(slopes, hikingMap[newPoint[0]][newPoint[1]]) {
			counter++
		}
	}

	return counter > 1
}

// point: [y, x]
func nextDirections(hikingMap [][]int, point []int) [][]int {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	if hikingMap[point[0]][point[1]] == SLOPE_UP {
		return [][]int{{-1, 0}}
	} else if hikingMap[point[0]][point[1]] == SLOPE_DOWN {
		return [][]int{{1, 0}}
	} else if hikingMap[point[0]][point[1]] == SLOPE_RIGHT {
		return [][]int{{0, 1}}
	} else if hikingMap[point[0]][point[1]] == SLOPE_LEFT {
		return [][]int{{0, -1}}
	}

	possibleDirs := make([][]int, 0)
	for _, dir := range dirs {
		newPoint := []int{point[0] + dir[0], point[1] + dir[1]}
		if newPoint[0] < 0 || newPoint[0] >= len(hikingMap) || newPoint[1] < 0 || newPoint[1] >= len(hikingMap[0]) {
			continue
		}

		if hikingMap[newPoint[0]][newPoint[1]] == FOREST {
			continue
		}

		if (dir[0] == 1) && (dir[1] == 0) && (hikingMap[newPoint[0]][newPoint[1]] == SLOPE_UP) {
			continue
		} else if (dir[0] == -1) && (dir[1] == 0) && (hikingMap[newPoint[0]][newPoint[1]] == SLOPE_DOWN) {
			continue
		} else if (dir[0] == 0) && (dir[1] == -1) && (hikingMap[newPoint[0]][newPoint[1]] == SLOPE_RIGHT) {
			continue
		} else if (dir[0] == 0) && (dir[1] == 1) && (hikingMap[newPoint[0]][newPoint[1]] == SLOPE_LEFT) {
			continue
		}

		possibleDirs = append(possibleDirs, dir)
	}

	return possibleDirs
}

// point: [y, x]
func coordToNode(point []int) string {
	return fmt.Sprint(point[1]) + "," + fmt.Sprint(point[0])
}

func addEdges(g *graph.Graph[string, string], hikingMap [][]int) {
	queue := make([][]int, 0) // [lastX, lastY, currX, currY, cost]
	startPoint := findStartPoint(hikingMap)
	queue = append(queue, []int{startPoint[1], startPoint[0], startPoint[1], startPoint[0], 0})
	(*g).AddVertex(coordToNode(startPoint))
	history := make([][]int, 0)

	lastX := startPoint[1]
	lastY := startPoint[0]
	cost := 0
	for len(queue) > 0 {
		currEntry := queue[0]
		queue = queue[1:]
		lastX = currEntry[0]
		lastY = currEntry[1]
		currX := currEntry[2]
		currY := currEntry[3]
		cost = currEntry[4]

		if currX < 0 || currX >= len(hikingMap) || currY < 0 || currY >= len(hikingMap[0]) {
			continue
		}

		if isNode(hikingMap, []int{currY, currX}) && !(startPoint[1] == currX && startPoint[0] == currY) {
			(*g).AddVertex(coordToNode([]int{currY, currX}))
			(*g).AddEdge(coordToNode([]int{lastY, lastX}), coordToNode([]int{currY, currX}), graph.EdgeWeight(cost))
			lastX = currX
			lastY = currY
			cost = 0
		}

		if contains(history, []int{currY, currX}) {
			continue
		}

		history = append(history, []int{currY, currX})

		for _, dir := range nextDirections(hikingMap, []int{currY, currX}) {
			newX := currX + dir[1]
			newY := currY + dir[0]
			queue = append(queue, []int{lastX, lastY, newX, newY, cost + 1})
		}
	}

	(*g).AddVertex(coordToNode(findEndPoint(hikingMap)))
	(*g).AddEdge(coordToNode([]int{lastY, lastX}), coordToNode(findEndPoint(hikingMap)), graph.EdgeWeight(cost))
}

func createGraph(hikingMap [][]int) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Weighted(), graph.Acyclic(), graph.Directed())

	addEdges(&g, hikingMap)

	return g
}

func contains(history [][]int, point []int) bool {
	for _, p := range history {
		if p[0] == point[0] && p[1] == point[1] {
			return true
		}
	}
	return false
}

func getPathLength(g graph.Graph[string, string], path []string) int {
	length := 0
	for i := 0; i < len(path)-1; i++ {
		edge, err := g.Edge(path[i], path[i+1])
		check(err)
		length += edge.Properties.Weight
	}
	return length
}

func getMaxLength(g graph.Graph[string, string], startPoint string, endPoint string) int {
	paths, err := graph.AllPathsBetween(g, startPoint, endPoint)
	check(err)

	fmt.Println("Paths:", len(paths))

	maxLen := 0
	for _, path := range paths {
		length := getPathLength(g, path)
		maxLen = max(maxLen, length)
	}

	return maxLen
}

func printMaze(hikingMap [][]int, nodes []string) {
	for i, row := range hikingMap {
		for j, col := range row {
			if slices.Contains(nodes, coordToNode([]int{i, j})) {
				fmt.Print("\033[31mX\033[0m")
			} else if col == PATH {
				fmt.Print(".")
			} else if col == FOREST {
				fmt.Print("#")
			} else if col == SLOPE_UP {
				fmt.Print("^")
			} else if col == SLOPE_DOWN {
				fmt.Print("v")
			} else if col == SLOPE_RIGHT {
				fmt.Print(">")
			} else if col == SLOPE_LEFT {
				fmt.Print("<")
			}
		}
		fmt.Println()
	}
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
	hikingMap := make([][]int, 0)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		hikingMap = append(hikingMap, handleLine(line))
	}

	g := createGraph(hikingMap)

	numEdges, err := g.Size()
	check(err)
	fmt.Println("Edges:", numEdges)

	nodes, err := graph.TopologicalSort(g)
	check(err)

	fmt.Println("Nodes:", len(nodes))
	// printMaze(hikingMap, nodes)

	sum = getMaxLength(g, coordToNode(findStartPoint(hikingMap)), coordToNode(findEndPoint(hikingMap)))

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
