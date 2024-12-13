package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Vec2d struct {
	x int
	y int
}

func isInteger(f float64) bool {
	return f == float64(int(f))
}

type Machine struct {
	moveA Vec2d
	moveB Vec2d
	prize Vec2d
}

func (m Machine) Print() {
	fmt.Printf("A: X=%d, Y=%d\n", m.moveA.x, m.moveA.y)
	fmt.Printf("B: X=%d, Y=%d\n", m.moveB.x, m.moveB.y)
	fmt.Printf("Prize: X=%d, Y=%d\n", m.prize.x, m.prize.y)
}

const TOKEN_BUTTON_A = 3
const TOKEN_BUTTON_B = 1
const CONVERSION_ADDITION = 10000000000000

func handleLine(lines []string) int {
	machine := parseMachine(lines)
	// fmt.Println("")
	// machine.Print()

	return calcPrizeCost(machine)
}

func parseMachine(lines []string) Machine {
	var machine Machine

	machine.moveA = Vec2d{0, 0}
	machine.moveB = Vec2d{0, 0}
	machine.prize = Vec2d{0, 0}

	n, err := fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &machine.moveA.x, &machine.moveA.y)
	check(err)
	if n != 2 {
		fmt.Printf("Error parsing Button A of line %s\n", lines[0])
	}

	n, err = fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &machine.moveB.x, &machine.moveB.y)
	check(err)
	if n != 2 {
		fmt.Printf("Error parsing Button B of line %s\n", lines[1])
	}

	n, err = fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &machine.prize.x, &machine.prize.y)
	check(err)
	if n != 2 {
		fmt.Printf("Error parsing Price of line %s\n", lines[2])
	}

	// add conversion factor to the prize
	machine.prize.x += CONVERSION_ADDITION
	machine.prize.y += CONVERSION_ADDITION

	return machine
}

func calcPrizeCost(machine Machine) int {
	btnA := machine.moveA
	btnB := machine.moveB
	prize := machine.prize

	a1 := btnA.x
	a2 := btnA.y
	b1 := btnB.x
	b2 := btnB.y
	p1 := prize.x
	p2 := prize.y

	x1 := float64(b2*p1-p2*b1) / float64(a1*b2-a2*b1)
	x2 := float64(a1*p2-p1*a2) / float64(a1*b2-a2*b1)

	// check if the solution is integer (round and check if it is equal to the original)
	if !isInteger(x1) || !isInteger(x2) {
		// fmt.Printf("Solution is not integer: \n%.3f, %.3f\n", x1, x2)
		return 0
	}

	cost := int(x1*TOKEN_BUTTON_A) + int(x2*TOKEN_BUTTON_B)
	// fmt.Printf("%.3f, %.3f\n", x1, x2)
	// fmt.Printf("Cost: %d\n", cost)

	return cost
}

func main() {
	// Open the file
	file, err := os.Open("13.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	// Iterate through each line
	for scanner.Scan() {
		lineButtonA := scanner.Text()
		if !scanner.Scan() {
			panic("Expected line for Button B")
		}
		lineButtonB := scanner.Text()
		if !scanner.Scan() {
			panic("Expected line for Price")
		}
		linePrice := scanner.Text()
		scanner.Scan() // Skip empty line

		lines := []string{lineButtonA, lineButtonB, linePrice}
		sum += handleLine(lines)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
