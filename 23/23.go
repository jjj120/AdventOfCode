package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

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

	possibleDirs := make([][]int, 0)
	for _, dir := range dirs {
		newPoint := []int{point[0] + dir[0], point[1] + dir[1]}
		if newPoint[0] < 0 || newPoint[0] >= len(hikingMap) || newPoint[1] < 0 || newPoint[1] >= len(hikingMap[0]) {
			continue
		}

		if hikingMap[newPoint[0]][newPoint[1]] == FOREST {
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
	queue := make([][]int, 0) // [lastX, lastY, currX, currY, prevX, prevY, cost]
	startPoint := findStartPoint(hikingMap)
	queue = append(queue, []int{startPoint[1], startPoint[0], startPoint[1], startPoint[0], startPoint[1], startPoint[0], 0})
	(*g).AddVertex(coordToNode(startPoint))
	history := make([][]int, 0)

	for len(queue) > 0 {
		currEntry := queue[0]
		queue = queue[1:]

		lastX := currEntry[0]
		lastY := currEntry[1]
		currX := currEntry[2]
		currY := currEntry[3]
		prevX := currEntry[4]
		prevY := currEntry[5]
		cost := currEntry[6]

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

		if contains(history, []int{currY, currX, prevY, prevX}) {
			continue
		}

		history = append(history, []int{currY, currX, prevY, prevX})

		for _, dir := range nextDirections(hikingMap, []int{currY, currX}) {
			newX := currX + dir[1]
			newY := currY + dir[0]
			if newX == prevX && newY == prevY {
				continue
			}
			queue = append(queue, []int{lastX, lastY, newX, newY, currX, currY, cost + 1})
		}
	}
	edges, err := (*g).Edges()
	check(err)
	for _, edge := range edges {
		if edge.Properties.Weight == 0 {
			(*g).RemoveEdge(edge.Source, edge.Target)
		}
		if strings.Compare(edge.Source, edge.Target) == 0 {
			(*g).RemoveEdge(edge.Source, edge.Target)
		}
	}
}

func createGraph(hikingMap [][]int) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Weighted())

	addEdges(&g, hikingMap)

	return g
}

func contains(history [][]int, point []int) bool {
	for _, p := range history {
		if p[0] == point[0] && p[1] == point[1] {
			if len(point) == 2 {
				return true
			} else if p[2] == point[2] && p[3] == point[3] {
				return true
			}
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
	fmt.Println(paths)

	maxLen := 0
	for _, path := range paths {
		length := getPathLength(g, path)
		maxLen = max(maxLen, length)
	}

	return maxLen
}

type QueueEntry struct {
	node string
	len  map[string]int
}

func getMaxLenUndirected(g graph.Graph[string, string], startPoint string, endPoint string) int {
	queue := make([]QueueEntry, 0)
	maxLen := 0
	queue = append(queue, QueueEntry{node: startPoint, len: make(map[string]int)})

	adj, err := g.AdjacencyMap()
	check(err)

	for len(queue) > 0 {
		// if len(queue) >= 100000 && len(queue)%100000 == 0 {
		// 	fmt.Print("\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\rQueue:", len(queue))
		// }
		currEntry := queue[0]
		queue = queue[1:]

		for _, edge := range adj[currEntry.node] {
			if currLen, ok := currEntry.len[edge.Target]; ok || currLen > 142*142 {
				if currLen > 142*142 {
					fmt.Println("Too long:", currLen, edge.Source, edge.Target)
					panic("Too long")
				}
				continue
			}

			nextNode := edge.Target
			nextLen := currEntry.len[edge.Source] + edge.Properties.Weight

			len := map[string]int{}
			for k, v := range currEntry.len {
				len[k] = v
			}

			len[nextNode] = nextLen
			queue = append(queue, QueueEntry{node: nextNode, len: len})
		}
		if _, ok := currEntry.len[endPoint]; ok {
			if currEntry.len[endPoint] > maxLen {
				fmt.Print("\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\rFound: ", currEntry.len[endPoint], " with queue ", len(queue), " and maxLen ", maxLen)
			}
			maxLen = max(maxLen, currEntry.len[endPoint])
		}
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

	nodes := make([]string, 0)
	graph.DFS(g, coordToNode(findStartPoint(hikingMap)), func(value string) bool {
		nodes = append(nodes, value)
		return false
	})

	fmt.Println("Nodes:", len(nodes))
	printMaze(hikingMap, nodes)

	// sum = getMaxLength(g, coordToNode(findStartPoint(hikingMap)), coordToNode(findEndPoint(hikingMap)))
	sum = getMaxLenUndirected(g, coordToNode(findStartPoint(hikingMap)), coordToNode(findEndPoint(hikingMap)))

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("\nSum: %d\n", sum)
}
