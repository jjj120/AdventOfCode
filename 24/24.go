package main

import (
	"bufio"
	"fmt"
	"os"
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

	var sum = 0
	gatesOutput := make(map[string]Gate)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parseStartValue(line, gatesOutput)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parseGates(line, gatesOutput)
	}

	for output := range gatesOutput {
		if strings.HasPrefix(output, "z") {
			outVal := gatesOutput[output].compute(gatesOutput)
			outNum, err := strconv.Atoi(output[1:])
			check(err)
			sum |= outVal << outNum
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
