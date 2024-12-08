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

func assert(t bool, s string) {
	if !t {
		panic("Assertion failed: " + s)
	}
}

type Coord struct {
	x int
	y int
}

func (a Coord) Add(b Coord) Coord {
	a.x += b.x
	a.y += b.y
	return a
}

func (a Coord) Sub(b Coord) Coord {
	a.x -= b.x
	a.y -= b.y
	return a
}

func parseAntennas(antennas map[rune][]Coord, line string, lineIndex int) map[rune][]Coord {
	for index, char := range line {
		if char != '.' {
			antennas[char] = append(antennas[char], Coord{x: index, y: lineIndex})
		}
	}
	return antennas
}

func countResonantAll(antennas map[rune][]Coord, fieldSizeX, fieldSizeY int) int {
	resonants := make(map[Coord]bool)
	for _, antennaCoords := range antennas {
		if len(antennaCoords) > 1 {
			resonantsNew := getResonants(antennaCoords, fieldSizeX, fieldSizeY)
			for resonant := range resonantsNew {
				resonants[resonant] = resonantsNew[resonant]
			}
		}
	}

	PrintResonants(antennas, resonants, fieldSizeX, fieldSizeY)

	return len(resonants)
}

func getResonants(antennaCoords []Coord, fieldSizeX, fieldSizeY int) map[Coord]bool {
	resonants := make(map[Coord]bool)

	for _, coord1 := range antennaCoords {
		for _, coord2 := range antennaCoords {
			if coord1 != coord2 {
				for _, resonant := range getResonantsSingle(coord1, coord2, fieldSizeX, fieldSizeY) {
					resonants[resonant] = true
				}
			}
		}
	}
	return resonants
}

func getResonantsSingle(coord1, coord2 Coord, fieldSizeX, fieldSizeY int) []Coord {
	diff := coord2.Sub(coord1)

	resonants := []Coord{
		coord1.Sub(diff),
		coord2.Add(diff),
	}

	resonantsClean := make([]Coord, 0)
	for _, resonant := range resonants {
		if resonant.x < 0 || resonant.x >= fieldSizeX || resonant.y < 0 || resonant.y >= fieldSizeY {
			continue
		}
		resonantsClean = append(resonantsClean, resonant)
	}

	return resonantsClean
}

func PrintResonants(antennas map[rune][]Coord, resonants map[Coord]bool, fieldSizeX, fieldSizeY int) {
	for y := 0; y < fieldSizeY; y++ {
		for x := 0; x < fieldSizeX; x++ {
			coord := Coord{x: x, y: y}
			//normal color
			color := "\033[0m"

			if _, ok := resonants[coord]; ok {
				color = "\033[35m"
			}

			antennaRune := '•'
			for antennaName, antennaCoords := range antennas {
				for _, antennaCoord := range antennaCoords {
					if coord == antennaCoord {
						antennaRune = antennaName
						break
					}
				}
			}
			fmt.Print(color + string(antennaRune) + "\033[0m")

		}
		fmt.Println()
	}
}

func PrintRuneMap(antennas map[rune][]Coord) {
	for antennaName, antennaCoords := range antennas {
		fmt.Printf("Antenna %c: %v\n", antennaName, antennaCoords)
	}
}

func main() {
	// Open the file
	file, err := os.Open("08.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	var i = 0
	antennas := make(map[rune][]Coord)
	fieldSizeX := 0
	fieldSizeY := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		antennas = parseAntennas(antennas, line, i)
		i++
		fieldSizeX = max(fieldSizeX, len(line))
		fieldSizeY = i
	}

	fmt.Printf("Field size: %d x %d\n", fieldSizeX, fieldSizeY)

	sum = countResonantAll(antennas, fieldSizeX, fieldSizeY)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	assert(sum < 317, "Test failed, result should not be 317 or bigger")
	assert(sum < 311, "Test failed, result should not be 311 or bigger")
	assert(sum < 306, "Test failed, result should not be 306 or bigger")

	fmt.Printf("Sum: %d\n", sum)
}
