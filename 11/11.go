package main

import (
	"fmt"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 11
const selectExample = false

type CacheKey struct {
	from, to         string
	vis_dac, vis_fft bool
}

func countPaths(devices map[string][]string, from, to string, vis_dac, vis_fft bool, cache *map[CacheKey]int) int {
	sum := 0
	if from == "dac" {
		vis_dac = true
	}
	if from == "fft" {
		vis_fft = true
	}
	if v, ok := (*cache)[CacheKey{from: from, to: to, vis_dac: vis_dac, vis_fft: vis_fft}]; ok {
		return v
	}
	for _, d := range devices[from] {
		if d == to {
			if vis_dac && vis_fft {
				sum++
			}
		} else {
			sum += countPaths(devices, d, to, vis_dac, vis_fft, cache)
		}
	}
	(*cache)[CacheKey{from: from, to: to, vis_dac: vis_dac, vis_fft: vis_fft}] = sum
	return sum
}

func handleLines(lines []string) int {
	devices := make(map[string][]string)
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		deviceName := splitLine[0]
		splitLine = strings.Split(splitLine[1], " ")
		devices[deviceName] = splitLine
	}

	cache := make(map[CacheKey]int)
	return countPaths(devices, "svr", "out", false, false, &cache)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 2, "Example wrong!")
	}
}
