package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type connection struct {
	from string
	to   string
	cost int
}

func handleLine(connections map[string][]connection, line string) {
	str := strings.Split(line, " = ")
	str2 := strings.Split(str[0], " to ")

	var conn connection
	conn.from = str2[0]
	conn.to = str2[1]
	num, err := strconv.ParseInt(str[1], 10, 0)
	check(err)
	conn.cost = int(num)

	connections[conn.from] = append(connections[conn.from], conn)
	connections[conn.to] = append(connections[conn.to], conn)
}

func connects(connections map[string][]connection, from, to string) int {
	for _, conns := range connections {
		for _, conn := range conns {
			if (strings.Compare(from, conn.from) == 0 && strings.Compare(to, conn.to) == 0) || (strings.Compare(from, conn.to) == 0 && strings.Compare(to, conn.from) == 0) {
				return conn.cost
			}
		}
	}
	return -1
}

func findLen(connections map[string][]connection, history []string) int {
	if len(connections) == len(history) {
		// fmt.Printf("%d Finished with history: ", len(history))
		// fmt.Println(history)
		return 0
	}

	shortest := 0
	for key := range connections {
		if !slices.Contains(history, key) {
			history2 := make([]string, len(history))
			copy(history2, history)
			history2 = append(history2, key)

			costs := 0
			if !(len(history) == 0) {
				costs = connects(connections, history[len(history)-1], key)
				if costs == -1 {
					continue
				}
			}
			shortest = max(shortest, findLen(connections, history2)+costs)
		}
	}

	return shortest
}

func findShortest(connections map[string][]connection) int {
	history := make([]string, 0)
	shortest := findLen(connections, history)
	return shortest
}

func main() {
	// Open the file
	file, err := os.Open("09.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	connections := make(map[string][]connection)
	var sum = 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(connections, line)
	}

	sum = findShortest(connections)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
