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

func assert(cond bool, msg string) {
	if !cond {
		panic("Assertion failed: " + msg)
	}
}

type Coord struct {
	x, y int
}

const WIDTH = 101
const HEIGHT = 103
const SIM_TIME = 100 // seconds

type Robot struct {
	px, py int
	vx, vy int
}

func parseRobot(line string) Robot {
	var r Robot
	_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)
	check(err)
	return r
}

func simRobots(robots []Robot, t int) {
	for i := range robots {
		robots[i].px += robots[i].vx * t
		robots[i].py += robots[i].vy * t

		robots[i].px = (robots[i].px + t*WIDTH) % WIDTH   // get the robots back into the field
		robots[i].py = (robots[i].py + t*HEIGHT) % HEIGHT // get the robots back into the field

		assert(robots[i].px >= 0 && robots[i].px < WIDTH, "robot out of bounds")
		assert(robots[i].py >= 0 && robots[i].py < HEIGHT, "robot out of bounds")
	}
}

func countScore(robots []Robot) int {
	numQuadrants := new([4]int)
	fmt.Printf("Width: %d, Height: %d, Width/2: %d, Height/2: %d\n", WIDTH, HEIGHT, WIDTH/2, HEIGHT/2)
	for _, r := range robots {
		if r.px < WIDTH/2 && r.py < HEIGHT/2 {
			// first quadrant
			numQuadrants[0]++
		} else if r.px > WIDTH/2 && r.py < HEIGHT/2 {
			// second quadrant
			numQuadrants[1]++
		} else if r.px < WIDTH/2 && r.py > HEIGHT/2 {
			// third quadrant
			numQuadrants[2]++
		} else if r.px > WIDTH/2 && r.py > HEIGHT/2 {
			// fourth quadrant
			numQuadrants[3]++
		} else {
			// fmt.Printf("Error: robot at (%d, %d) is not in any quadrant\n", r.px, r.py)
		}
	}

	fmt.Printf("Quadrants: %v\n", numQuadrants)
	score := numQuadrants[0] * numQuadrants[1] * numQuadrants[2] * numQuadrants[3]
	return score
}

func printRobotsField(robots []Robot) {
	positionCount := make(map[Coord]int)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			for _, r := range robots {
				if r.px == x && r.py == y {
					positionCount[Coord{x, y}]++
				}
			}
		}
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if positionCount[Coord{x, y}] > 0 {
				fmt.Printf("%d", positionCount[Coord{x, y}])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printRobots(robots []Robot) {
	for _, r := range robots {
		fmt.Printf("p=(%d, %d) v=(%d, %d)\n", r.px, r.py, r.vx, r.vy)
	}
}

func main() {
	// Open the file
	file, err := os.Open("14.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0, 20)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		robots = append(robots, parseRobot(line))
	}

	printRobots(robots)
	fmt.Println("")
	printRobotsField(robots)
	simRobots(robots, SIM_TIME)
	printRobotsField(robots)

	var sum = 0
	sum = countScore(robots)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	assert(sum != 216012712, "Wrong answer 216012712")
	assert(sum != 88046400, "Wrong answer 88046400")
	fmt.Printf("Sum: %d\n", sum)
}
