package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	AND = iota
	OR
	LSHIFT
	RSHIFT
	NOT
	ASSIGNMENT
)

type Operation struct {
	assignedTo  string
	value1      string
	value2      string
	operation   int
	solved      bool
	solvedValue uint16
}

func handleLine(assignments *map[string]Operation, line string) int {
	var operaton Operation
	operaton.solved = false
	operaton.solvedValue = 0
	if strings.Contains(line, "NOT") {
		str := strings.Split(line, " ")
		operaton.operation = NOT
		operaton.assignedTo = str[3]
		operaton.value1 = str[1]
		operaton.value2 = str[1]
		(*assignments)[operaton.assignedTo] = operaton
		return 0
	}

	if strings.Contains(line, "AND") {
		operaton.operation = AND
		str := strings.Split(line, " ")
		operaton.assignedTo = str[4]
		operaton.value1 = str[0]
		operaton.value2 = str[2]
		(*assignments)[operaton.assignedTo] = operaton
		return 0
	}

	if strings.Contains(line, "OR") {
		operaton.operation = OR
		str := strings.Split(line, " ")
		operaton.assignedTo = str[4]
		operaton.value1 = str[0]
		operaton.value2 = str[2]
		(*assignments)[operaton.assignedTo] = operaton
		return 0
	}

	if strings.Contains(line, "LSHIFT") {
		operaton.operation = LSHIFT
		str := strings.Split(line, " ")
		operaton.assignedTo = str[4]
		operaton.value1 = str[0]
		operaton.value2 = str[2]
		(*assignments)[operaton.assignedTo] = operaton
		return 0
	}

	if strings.Contains(line, "RSHIFT") {
		operaton.operation = RSHIFT
		str := strings.Split(line, " ")
		operaton.assignedTo = str[4]
		operaton.value1 = str[0]
		operaton.value2 = str[2]
		(*assignments)[operaton.assignedTo] = operaton
		return 0
	}

	operaton.operation = ASSIGNMENT
	str := strings.Split(line, " ")
	operaton.assignedTo = str[2]
	operaton.value1 = str[0]
	operaton.value2 = str[0]
	(*assignments)[operaton.assignedTo] = operaton

	fmt.Printf("%s:\t%s,%s -> %s\n", line, operaton.value1, operaton.value2, operaton.assignedTo)

	return 0
}

func evaluate(assignments *map[string]Operation, solveFor string) uint16 {
	if len(solveFor) == 0 {
		panic("THERE SHOULD BE NO EMPTY STRING IN THE MAP!!")
	}

	fmt.Printf("Solving for %s\n", solveFor)
	currOperation := (*assignments)[solveFor]

	if currOperation.solved {
		return currOperation.solvedValue
	}

	if currOperation.operation == ASSIGNMENT {
		num, err := strconv.ParseUint(currOperation.value1, 10, 16)

		if err != nil {
			// assigning variable to variable
			currOperation.solved = true
			currOperation.solvedValue = evaluate(assignments, currOperation.value1)
			(*assignments)[solveFor] = currOperation
			return currOperation.solvedValue
		}

		// assigning number to variable, basecase
		return uint16(num)
	}

	var v1 uint16
	num1, err1 := strconv.ParseUint(currOperation.value1, 10, 16)
	if err1 != nil {
		// assigning variable to variable
		currOperation.solved = true
		currOperation.solvedValue = evaluate(assignments, currOperation.value1)
		(*assignments)[solveFor] = currOperation
		v1 = currOperation.solvedValue
	} else {
		// assigning number to variable, basecase
		v1 = uint16(num1)
	}

	var v2 uint16
	num2, err2 := strconv.ParseUint(currOperation.value2, 10, 16)
	if err2 != nil {
		// assigning variable to variable
		currOperation.solved = true
		currOperation.solvedValue = evaluate(assignments, currOperation.value2)
		(*assignments)[solveFor] = currOperation
		v2 = currOperation.solvedValue
	} else {
		// assigning number to variable, basecase
		v2 = uint16(num2)
	}

	// fmt.Printf("%d %d %d\n", v1, currOperation.operation, v2)

	switch currOperation.operation {
	case AND:
		currOperation.solvedValue = v1 & v2
	case OR:
		currOperation.solvedValue = v1 | v2
	case LSHIFT:
		currOperation.solvedValue = v1 << v2
	case RSHIFT:
		currOperation.solvedValue = v1 >> v2
	case NOT:
		currOperation.solvedValue = ^v1
	}

	currOperation.solved = true
	(*assignments)[solveFor] = currOperation
	return currOperation.solvedValue
}

func main() {
	// Open the file
	file, err := os.Open("07.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	assignments := make(map[string]Operation)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(&assignments, line)
	}

	newAssignments := make(map[string]Operation)
	for k, v := range assignments {
		newAssignments[k] = v
	}

	var op Operation
	op.assignedTo = "b"
	op.operation = ASSIGNMENT
	op.solved = true
	op.solvedValue = evaluate(&assignments, "a")

	fmt.Printf("b: %d\n", op.solvedValue)

	newAssignments[op.assignedTo] = op

	sum := evaluate(&newAssignments, "a")

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
