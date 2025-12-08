package main

import (
	"cmp"
	"fmt"
	"slices"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 8
const selectExample = false

type VecsWithDist struct {
	dist float64
	vec1 aoc.Vec3d
	vec2 aoc.Vec3d
}

func getSortedPairs(boxes []aoc.Vec3d) []VecsWithDist {
	pairs := make([]VecsWithDist, 0)

	for iA, vecA := range boxes {
		for _, vecB := range boxes[iA+1:] {
			dist := vecA.Distance(vecB)
			pairs = append(pairs, VecsWithDist{dist: dist, vec1: vecA, vec2: vecB})
		}
	}
	slices.SortFunc(pairs, func(a, b VecsWithDist) int { return cmp.Compare(a.dist, b.dist) })

	return pairs
}

type VecIntMap map[aoc.Vec3d]int

func (m VecIntMap) includes(v aoc.Vec3d) bool {
	_, ok := m[v]
	return ok
}

func getCircuitLen(boxes []aoc.Vec3d, limit int) int {
	circuits := make(VecIntMap)

	nextCircIndex := 1
	circCount := 0
	includedBoxes := 0

	sortedDists := getSortedPairs(boxes)

	for _, pair := range sortedDists {
		if circuits.includes(pair.vec1) && circuits.includes(pair.vec2) {
			// merge the two circuits
			circ1 := circuits[pair.vec1]
			circ2 := circuits[pair.vec2]
			if circ1 == circ2 {
				continue
			}
			for v, c := range circuits {
				if c == circ1 {
					circuits[v] = circ2
				}
			}
			circCount--
		} else if circuits.includes(pair.vec1) {
			// does not include vec2
			circuits[pair.vec2] = circuits[pair.vec1]
			includedBoxes++
		} else if circuits.includes(pair.vec2) {
			// does not include vec1
			circuits[pair.vec1] = circuits[pair.vec2]
			includedBoxes++
		} else {
			// new circuit!
			circuits[pair.vec1] = nextCircIndex
			circuits[pair.vec2] = nextCircIndex
			nextCircIndex++
			circCount++
			includedBoxes += 2
		}

		if includedBoxes == len(boxes) && circCount == 1 {
			// all connected
			return pair.vec1.X * pair.vec2.X
		}
	}

	return -1
}

func handleLines(lines []string) int {
	boxes := make([]aoc.Vec3d, 0, len(lines))
	for _, line := range lines {
		var newBox aoc.Vec3d
		_, err := fmt.Sscanf(line, "%d,%d,%d", &newBox.X, &newBox.Y, &newBox.Z)
		boxes = append(boxes, newBox)
		aoc.Check(err)
	}

	if selectExample {
		return getCircuitLen(boxes, 10)
	}
	return getCircuitLen(boxes, 1000)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 25272, "Example wrong!")
	} else {
		aoc.Assert(sum < 8537956038, "Too high!")
	}
}
