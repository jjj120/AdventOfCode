package main

import (
	"bufio"
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
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

func handleLine(lines []string) int {
	machine := parseMachine(lines)

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

	return machine
}

func calcPrizeCost(machine Machine) int {
	btnA := machine.moveA
	btnB := machine.moveB
	prize := machine.prize
	var (
		A = mat.NewDense(2, 2, []float64{float64(btnA.x), float64(btnB.x), float64(btnA.y), float64(btnB.y)})
		b = mat.NewDense(2, 1, []float64{float64(prize.x), float64(prize.y)})
		x = mat.NewDense(2, 1, nil)
	)

	// Solve for x such that Ax = b
	var qr mat.QR
	qr.Factorize(A)

	err := qr.SolveTo(x, false, b)
	if err != nil {
		fmt.Printf("could not solve QR: %+v", err)
	}

	// check if the solution is integer (round and check if it is equal to the original)
	roundedX := mat.NewDense(2, 1, nil)
	roundedX.Apply(func(i, j int, v float64) float64 {
		return float64(int(v + 0.5))
	}, x)

	if !mat.EqualApprox(roundedX, x, 1e-6) {
		// fmt.Printf("Solution is not integer: \n%v\n", mat.Formatted(x))
		return 0
	}

	cost := int(roundedX.At(0, 0))*TOKEN_BUTTON_A + int(roundedX.At(1, 0))*TOKEN_BUTTON_B
	// fmt.Printf("%.3f\n", mat.Formatted(x))
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
