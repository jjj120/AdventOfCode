package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Hailstone struct {
	x, y, z    int
	vx, vy, vz int
}

const TESTAREA_MIN = 200000000000000
const TESTAREA_MAX = 400000000000000

// const TESTAREA_MIN = 7
// const TESTAREA_MAX = 27

const PRINT = false

func (h *Hailstone) move() {
	h.x += h.vx
	h.y += h.vy
	h.z += h.vz
}

func (h1 *Hailstone) intersects(h2 *Hailstone) bool {
	if math.Abs((float64(h1.vx)*float64(h2.vy) - float64(h1.vy)*float64(h2.vx))) < 0.000000000001 {
		return false
	}

	// t := ((h2.x-h1.x)*h2.vy - (h2.y-h1.y)*h2.vx) / (h1.vx*h2.vy - h1.vy*h2.vx)
	t := (float64(h2.x-h1.x)*float64(h2.vy) - float64(h2.y-h1.y)*float64(h2.vx)) / (float64(h1.vx)*float64(h2.vy) - float64(h1.vy)*float64(h2.vx))
	// s := ((h2.x-h1.x)*h1.vy - (h2.y-h1.y)*h1.vx) / (h1.vx*h2.vy - h1.vy*h2.vx)
	s := (float64(h2.x-h1.x)*float64(h1.vy) - float64(h2.y-h1.y)*float64(h1.vx)) / (float64(h1.vx)*float64(h2.vy) - float64(h1.vy)*float64(h2.vx))

	intersectionX := float64(h1.x) + float64(h1.vx)*t
	intersectionY := float64(h1.y) + float64(h1.vy)*t

	if intersectionX < TESTAREA_MIN || intersectionX > TESTAREA_MAX {
		// out of bounds
		return false
	}
	if intersectionY < TESTAREA_MIN || intersectionY > TESTAREA_MAX {
		// out of bounds
		return false
	}
	if t < 0 {
		// intersection is in the past for h1
		return false
	}
	if s < 0 {
		// intersection is in the past for h2
		return false
	}

	if PRINT {
		fmt.Println("Hailstone 1:", h1.x, h1.y, h1.z, "@", h1.vx, h1.vy, h1.vz)
		fmt.Println("Hailstone 2:", h2.x, h2.y, h2.z, "@", h2.vx, h2.vy, h2.vz)
		fmt.Println("Intersection:", intersectionX, intersectionY, t)
		fmt.Println()
	}
	return true
}

func handleLine(line string) Hailstone {
	var hailstone Hailstone

	_, err := fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &hailstone.x, &hailstone.y, &hailstone.z, &hailstone.vx, &hailstone.vy, &hailstone.vz)
	check(err)

	return hailstone
}

func checkCombinations(hails []Hailstone) int {
	var sum = 0

	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			if hails[i].intersects(&hails[j]) {
				sum++
			}
		}
	}

	return sum
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
	hails := make([]Hailstone, 0)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		hails = append(hails, handleLine(line))
	}

	sum = checkCombinations(hails)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
