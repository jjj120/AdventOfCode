package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	opcode int
	arg    int
}

type Registers struct {
	a int
	b int
	c int
}

const (
	opcode_adv = 0
	opcode_bxl = 1
	opcode_bst = 2
	opcode_jnz = 3
	opcode_bxc = 4
	opcode_out = 5
	opcode_bdv = 6
	opcode_cdv = 7
)

const DEBUG = false

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func debugPrintf(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

func parseRegister(line string, registers Registers) Registers {
	var regA, regB, regC int
	_, errA := fmt.Sscanf(line, "Register A: %d", &regA)
	_, errB := fmt.Sscanf(line, "Register B: %d", &regB)
	_, errC := fmt.Sscanf(line, "Register C: %d", &regC)

	if errA == nil {
		registers.a = regA
	}
	if errB == nil {
		registers.b = regB
	}
	if errC == nil {
		registers.c = regC
	}
	return registers
}

func parseProgram(line string) []Instruction {
	var program []Instruction

	if !strings.HasPrefix(line, "Program: ") {
		fmt.Printf("Line: %s\n", line)
		panic("Invalid program")
	}

	line = line[len("Program: "):]

	for i := 0; i < len(line); i += 4 {
		var opcode, arg int
		_, err := fmt.Sscanf(line[i:i+3], "%d,%d", &opcode, &arg)
		check(err)
		program = append(program, Instruction{opcode, arg})
	}

	return program
}

func executeProgram(program []Instruction, registers Registers) []int {
	fmt.Printf("Registers: %v\n", registers)
	fmt.Printf("Program: %v\n", program)
	output := []int{}

	for pc := 0; pc < len(program); pc++ { // pc points to the instruction index, not the memory address
		instr := program[pc]
		combo := getComboOperand(instr.arg, registers)

		switch instr.opcode {
		case opcode_adv:
			debugPrintf("ADV: %d -- %d/%d = %d -> regA\n", combo, registers.a, 1<<combo, registers.a/(1<<combo))
			registers.a = registers.a / (1 << combo)
		case opcode_bxl:
			debugPrintf("BXL: %d -- %d ^ %d = %d -> regB\n", instr.arg, registers.b, instr.arg, registers.b^instr.arg)
			registers.b = registers.b ^ instr.arg
		case opcode_bst:
			debugPrintf("BST: %d -- %d %% 8 = %d -> regB\n", combo, combo, combo%8)
			registers.b = combo % 8
		case opcode_jnz:
			debugPrintf("JNZ: %d -- %d != 0 -> jumping to: %d\n", instr.arg/2, registers.a, instr.arg/2-1)
			if registers.a != 0 {
				pc = instr.arg/2 - 1
			}
		case opcode_bxc:
			debugPrintf("BXC: %d -- %d ^ %d = %d -> regB\n", combo, registers.b, registers.c, registers.b^registers.c)
			registers.b = registers.b ^ registers.c
		case opcode_out:
			debugPrintf("OUT: %d -- %d\n", combo, combo%8)
			fmt.Printf("%d,", combo%8)
			output = append(output, combo%8)
		case opcode_bdv:
			debugPrintf("BDV: %d -- %d / %d = %d -> regB\n", combo, registers.a, 1<<combo, registers.a/(1<<combo))
			registers.b = registers.a / (1 << combo)
		case opcode_cdv:
			debugPrintf("CDV: %d -- %d / %d = %d -> regC\n", combo, registers.a, 1<<combo, registers.a/(1<<combo))
			registers.c = registers.a / (1 << combo)
		}
	}
	fmt.Println("")
	// join output
	outString := ""
	for _, o := range output {
		outString += fmt.Sprintf("%d,", o)
	}
	outString = outString[:len(outString)-1]
	fmt.Println(outString)
	return output
}

func getComboOperand(combo int, registers Registers) int {
	var operand int
	switch combo {
	case 0, 1, 2, 3:
		operand = combo
	case 4:
		operand = registers.a
	case 5:
		operand = registers.b
	case 6:
		operand = registers.c
	default:
		operand = combo
	}
	return operand
}

func main() {
	// Open the file
	file, err := os.Open("17.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	registers := Registers{}
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		registers = parseRegister(line, registers)
	}

	scanner.Scan()
	program := parseProgram(scanner.Text())

	executeProgram(program, registers)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
