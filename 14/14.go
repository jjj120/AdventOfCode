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

type Reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
}

func parseReindeer(line string) Reindeer {
	var rein Reindeer
	fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &rein.name, &rein.speed, &rein.duration, &rein.rest)
	return rein
}

func handleLine(reindeers *[]Reindeer, line string) {
	rein := parseReindeer(line)
	*reindeers = append(*reindeers, rein)
}

func givePoints(dist []int, points []int) {
	max := 0
	for _, val := range dist {
		if val > max {
			max = val
		}
	}

	for i, val := range dist {
		if val == max {
			points[i]++
		}
	}
}

func simReindeers(reindeers []Reindeer, time int) int {
	deersDist := make([]int, len(reindeers))
	deersPoints := make([]int, len(reindeers))

	for t := 0; t < time; t++ {
		for i, rein := range reindeers {
			if t%(rein.duration+rein.rest) < rein.duration {
				deersDist[i] += rein.speed
			}
		}
		givePoints(deersDist, deersPoints)
	}

	max := 0
	for _, points := range deersPoints {
		if points > max {
			max = points
		}
	}

	return max
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

	var sum = 0
	var reindeers []Reindeer
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(&reindeers, line)
	}

	sum = simReindeers(reindeers, 2503)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
