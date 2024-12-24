package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func parseStartValue(line string, gates map[string]Gate) {
	split := strings.Split(line, ": ")
	assertf(len(split) == 2, "Invalid line format: %s", line)
	startVal, err := strconv.Atoi(split[1])
	check(err)
	gates[split[0]] = ConstGate{startVal}
}

func parseGates(line string, gates map[string]Gate) {
	split := strings.Split(line, " ")
	assertf(len(split) == 5, "Invalid line format: %s", line)
	input1 := split[0]
	input2 := split[2]
	output := split[4]
	switch split[1] {
	case "AND":
		gates[output] = &AndGate{input1, input2, -1, false}
	case "OR":
		gates[output] = &OrGate{input1, input2, -1, false}
	case "XOR":
		gates[output] = &XOrGate{input1, input2, -1, false}
	default:
		assertf(false, "Invalid gate type: %s", split[1])
	}
}

func generateDot(gates map[string]Gate) string {
	// go run *.go && dot -Tpng 24.dot -o out.png
	// go run *.go && dot -Tsvg 24.dot -o out.svg

	usedOutputs := make(map[string]bool)
	connections := make([]string, 0)
	nodes := make(map[string]string)
	outputSet := make(map[string]bool)

	// Collect all used outputs
	for _, gate := range gates {
		for _, input := range gate.getInputs() {
			usedOutputs[input] = true
		}
	}

	// Generate DOT nodes and connections
	for output, gate := range gates {
		if usedOutputs[output] || strings.HasPrefix(output, "x") || strings.HasPrefix(output, "y") || strings.HasPrefix(output, "z") {
			label := gate.Type() + "\n" + output
			settings := ""
			if strings.HasPrefix(output, "x") || strings.HasPrefix(output, "y") || strings.HasPrefix(output, "z") {
				settings = "style=filled color=red fillcolor=red"
			}
			nodes[output] = fmt.Sprintf(`"%s" [label="%s" %s];`, output, label, settings)
			for _, input := range gate.getInputs() {
				connections = append(connections, fmt.Sprintf(`"%s" -> "%s";`, input, output))
				if strings.HasPrefix(input, "x") || strings.HasPrefix(input, "y") || strings.HasPrefix(input, "z") {
					if !outputSet[input] {
						nodes[input] = fmt.Sprintf(`"%s" [label="%s", shape=ellipse];`, input, input)
						outputSet[input] = true
					}
				}
			}
		}
	}

	// Assemble the DOT graph
	var dotBuilder strings.Builder
	dotBuilder.WriteString("digraph G {\n")
	for _, node := range nodes {
		dotBuilder.WriteString("\t" + node + "\n")
	}
	for _, connection := range connections {
		dotBuilder.WriteString("\t" + connection + "\n")
	}
	dotBuilder.WriteString("}")

	return dotBuilder.String()
}

func makeSwaps(swaps map[string]string, gatesOutput map[string]Gate) map[string]Gate {
	gatesCopy := make(map[string]Gate)
	for k, v := range gatesOutput {
		gatesCopy[k] = v
	}

	for out, in := range swaps {
		gatesCopy[out], gatesCopy[in] = gatesCopy[in], gatesCopy[out]

		gatesCopy[in].invalidateCache()
		gatesCopy[out].invalidateCache()
	}
	return gatesCopy
}

func main() {
	// Open the file
	file, err := os.Open("24.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	numLines := 0
	// Iterate through each line
	for scanner.Scan() {
		// ignore the first part of the input
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		numLines++
	}

	gatesOutput := make(map[string]Gate)
	for scanner.Scan() {
		line := scanner.Text()
		parseGates(line, gatesOutput)
	}

	fmt.Printf("Number of lines: %d\n", numLines)
	// solution := bruteForceSwaps(gatesOutput, numLines/2)
	// fmt.Printf("Solution: %s\n", solution)

	// Generate the dot file
	dot := generateDot(gatesOutput)
	dotFile, err := os.Create("24.dot")
	check(err)
	defer dotFile.Close()
	_, err = dotFile.WriteString(dot)
	check(err)

	// Get solution by hand, check here:
	swaps := map[string]string{
		"z06": "dhg",
		"z23": "bhd",
		"z38": "nbf",
		"brk": "dpd",
	}

	// there is some error with the node that has output z06, z23, z38 and ??

	dot = generateDot(makeSwaps(swaps, gatesOutput))
	dotFile, err = os.Create("24.swapped.dot")
	check(err)
	defer dotFile.Close()
	_, err = dotFile.WriteString(dot)
	check(err)

	swapSlice := []string{}
	for k, v := range swaps {
		swapSlice = append(swapSlice, k, v)
	}
	sortedSlice := sort.StringSlice(swapSlice)
	sortedSlice.Sort()
	fmt.Printf("Solution: %s\n", strings.Join(sortedSlice, ","))

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
