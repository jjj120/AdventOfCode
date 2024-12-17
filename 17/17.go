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

func parseIntToProgram(line []int) []Instruction {
	var program []Instruction

	if len(line)%2 != 0 {
		return program // invalid program
	}

	for i := 0; i < len(line); i += 2 {
		program = append(program, Instruction{line[i], line[i+1]})
	}

	return program
}

func executeProgram(program []Instruction, registers Registers) []int {
	debugPrintf("Registers: %v\n", registers)
	debugPrintf("Program: %v\n", program)
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
			output = append(output, combo%8)
		case opcode_bdv:
			debugPrintf("BDV: %d -- %d / %d = %d -> regB\n", combo, registers.a, 1<<combo, registers.a/(1<<combo))
			registers.b = registers.a / (1 << combo)
		case opcode_cdv:
			debugPrintf("CDV: %d -- %d / %d = %d -> regC\n", combo, registers.a, 1<<combo, registers.a/(1<<combo))
			registers.c = registers.a / (1 << combo)
		}
	}
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

func findAinput(registers Registers, program []Instruction) int {
	matchFound := 1
	for a := 1; ; a++ {
		registers.a = a
		output := executeProgram(program, registers)

		matching, matchingNum := checkProgramEquality(program, parseIntToProgram(output))
		if matching {
			fmt.Printf("%16d %16o -- matching: %2d -- %s\n", a, a, matchingNum, programToString(parseIntToProgram(output)))
			return a
		}
		if matchingNum > matchFound {
			fmt.Printf("%16d %16o -- matching: %2d -- %s\n", a, a, matchingNum, programToString(parseIntToProgram(output)))
			matchFound = matchingNum
			a *= 64 // shift up by two numbers
		}
	}
}

func programToString(program []Instruction) string {
	str := ""
	for _, instr := range program {
		str += fmt.Sprintf("%d,%d,", instr.opcode, instr.arg)
	}
	return str
}

func checkProgramEquality(program1, program2 []Instruction) (bool, int) {
	matchingNum := 0
	matching := true
	for i := 1; i <= len(program1) && i <= len(program2); i++ {
		if program1[len(program1)-i].opcode != program2[len(program2)-i].opcode || program1[len(program1)-i].arg != program2[len(program2)-i].arg {
			return false, matchingNum
		} else {
			matchingNum++
		}
	}

	return matching && matchingNum == len(program1) && matchingNum == len(program2), matchingNum
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

	sum = findAinput(registers, program)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
