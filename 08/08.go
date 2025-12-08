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

func minDist(boxes []aoc.Vec3d) VecsWithDist {
	minimum := VecsWithDist{dist: boxes[0].Distance(boxes[1]), vec1: boxes[0], vec2: boxes[1]}
	for iA, vecA := range boxes {
		for _, vecB := range boxes[iA+1:] {
			dist := vecA.Distance(vecB)
			if dist < minimum.dist {
				minimum.dist = dist
				minimum.vec1 = vecA
				minimum.vec2 = vecB
			}
		}
	}
	return minimum
}

func minDistMap(boxes map[aoc.Vec3d]bool) VecsWithDist {
	boxesSlice := make([]aoc.Vec3d, 0)
	for v, b := range boxes {
		if b {
			boxesSlice = append(boxesSlice, v)
		}
	}
	return minDist(boxesSlice)
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
	// boxesMap := make(map[aoc.Vec3d]bool)
	// for _, box := range boxes {
	// 	boxesMap[box] = true
	// }

	circuits := make(VecIntMap)

	nextCircIndex := 1
	circCount := 0

	sortedDists := getSortedPairs(boxes)
	for _, v := range sortedDists[:limit] {
		fmt.Println(v)
	}

	for _, pair := range sortedDists[:limit] {
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
			continue
		}
		if circuits.includes(pair.vec1) {
			// does not include vec2
			circuits[pair.vec2] = circuits[pair.vec1]
			continue
		}
		if circuits.includes(pair.vec2) {
			// does not include vec1
			circuits[pair.vec1] = circuits[pair.vec2]
			continue
		}

		// new circuit!
		circuits[pair.vec1] = nextCircIndex
		circuits[pair.vec2] = nextCircIndex
		nextCircIndex++
		circCount++
	}

	circSizes := make(map[int]int)
	actualCircuits := make([][]aoc.Vec3d, nextCircIndex)

	for v, c := range circuits {
		fmt.Printf("%v: %d\n", v, c)
		if _, ok := circSizes[c]; ok {
			circSizes[c]++
			actualCircuits[c] = append(actualCircuits[c], v)
		} else {
			circSizes[c] = 1
			actualCircuits[c] = []aoc.Vec3d{v}
		}
	}

	for c, cs := range actualCircuits {
		fmt.Printf("Circ %d with size %d: %v\n", c, len(cs), cs)
	}

	slices.SortFunc(actualCircuits, func(a, b []aoc.Vec3d) int { return cmp.Compare(len(a), len(b)) })
	slices.Reverse(actualCircuits)

	return len(actualCircuits[0]) * len(actualCircuits[1]) * len(actualCircuits[2])
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
		aoc.Assert(sum == 40, "Example wrong!")
	}
}
