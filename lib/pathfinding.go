package lib

type Path struct {
	NodeStart Vec2d
	Nodes     []Vec2d
}

func FindShortestPath(grid [][]bool, start, end Vec2d) []Vec2d {
	// Create a queue of paths
	queue := []Path{
		{start, []Vec2d{start}},
	}

	// Create a set of visited nodes
	visited := make(map[Vec2d]bool)

	// While the queue is not empty
	for len(queue) > 0 {
		// Get the next path
		path := queue[0]
		queue = queue[1:]

		// Get the last node in the path
		node := path.Nodes[len(path.Nodes)-1]

		// If the node is the end, return the path
		if node == end {
			return path.Nodes
		}

		// If the node has been visited, skip it
		if visited[node] {
			continue
		}

		// Mark the node as visited
		visited[node] = true

		// Get the neighbors of the node
		neighbors := getValidNeighbors(grid, node)

		// Add the neighbors to the queue
		for _, neighbor := range neighbors {
			newPath := append(path.Nodes, neighbor)
			queue = append(queue, Path{neighbor, newPath})
		}
	}

	// If no path was found, return an empty path
	return []Vec2d{}
}

func getValidNeighbors(grid [][]bool, node Vec2d) []Vec2d {
	neighbors := []Vec2d{}

	// Check the four cardinal directions
	for _, neighbor := range getNeighbors(node) {
		if isValidNeighbor(grid, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func getNeighbors(node Vec2d) []Vec2d {
	return []Vec2d{
		node.Add(Vec2d{1, 0}),
		node.Add(Vec2d{0, 1}),
		node.Add(Vec2d{-1, 0}),
		node.Add(Vec2d{0, -1}),
	}
}

func isValidNeighbor(grid [][]bool, node Vec2d) bool {
	// Check if the node is within the bounds of the grid
	if node.X < 0 || node.X >= len(grid) || node.Y < 0 || node.Y >= len(grid[0]) {
		return false
	}

	// Check if the node is not blocked
	return !grid[node.X][node.Y]
}
