package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(g *graph.Graph[string, string], line string) {
	str := strings.Split(line, ": ")
	from := str[0]
	to := strings.Split(str[1], " ")

	(*g).AddVertex(from)
	for _, v := range to {
		(*g).AddVertex(v)
		(*g).AddEdge(from, v, graph.EdgeWeight(1))
	}
}

func initGraph() graph.Graph[string, string] {
	g := graph.New(graph.StringHash)
	return g
}

func getSize(g graph.Graph[string, string], removedEdges []graph.Edge[string]) int {

	for _, e := range removedEdges {
		// fmt.Printf("Removing edge %s -> %s\n", e.Source, e.Target)
		g.RemoveEdge(e.Source, e.Target)

	}

	cluster1 := make([]string, 0)
	edges, err := g.Edges()
	check(err)

	graph.DFS(g, edges[0].Source, func(vert string) bool {
		cluster1 = append(cluster1, vert)
		return false
	})

	adj, err := g.AdjacencyMap()
	check(err)
	size := len(adj)
	fmt.Printf("Cluster 1: %d\n", len(cluster1))
	fmt.Printf("Cluster 2: %d\n", size-len(cluster1))
	return len(cluster1) * (size - len(cluster1))
}

func main() {
	// Open the file
	file, err := os.Open("25.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	g := initGraph()

	var sum = 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(&g, line)
	}

	// fmt.Println(g.Size())

	// check if file exists
	_, err = os.Stat("25i.dot")
	if err != nil {
		fmt.Println("Dot file does not exist, creating...")
		file, err = os.Create("25i.dot")
		check(err)
		defer file.Close()

		_ = draw.DOT(g, file)

		os.Exit(0)
	}

	// let the graph get plotted with Gephi,
	// Make ForceAtlas layout,
	// then remove the edges that connect the two clusters
	removedEdges := []graph.Edge[string]{
		{Source: "qns", Target: "jxm"},
		{Source: "mgb", Target: "plt"},
		{Source: "dbt", Target: "tjd"},
	}
	sum = getSize(g, removedEdges)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
