package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 9
const selectExample = false

func printTileMap(tiles map[aoc.Vec2d]bool, width int, height int) {
	if !selectExample {
		return
	}
	for y := range height {
		for x := range width {
			pos := aoc.Vec2d{X: x, Y: y}

			if v, ok := tiles[pos]; ok && v {
				aoc.ColorPrint(aoc.ConstantToAnsiEscapeString(aoc.ANSI_BRIGHT_RED_FG), aoc.UNICODE_BLOCK)
			} else {
				fmt.Print(aoc.UNICODE_BLOCK)
			}
		}
		fmt.Println("")
	}
}

func getBiggestRect(tiles map[aoc.Vec2d]bool) int {
	maxSize := 0
	maxSizeP1 := aoc.Vec2d{X: 0, Y: 0}
	maxSizeP2 := aoc.Vec2d{X: 0, Y: 0}
	for p1 := range tiles {
		for p2 := range tiles {
			if p1.Equals(p2) {
				continue
			}
			size := (max(p1.X, p2.X) - min(p1.X, p2.X) + 1) * (max(p1.Y, p2.Y) - min(p1.Y, p2.Y) + 1)

			if size > maxSize {
				maxSizeP1 = p1
				maxSizeP2 = p2
				maxSize = size
			}
		}
	}
	fmt.Printf("MaxSize P1 is %v\n", maxSizeP1)
	fmt.Printf("MaxSize P2 is %v\n", maxSizeP2)
	fmt.Printf("with size %d\n", maxSize)
	return maxSize
}

func handleLines(lines []string) int {
	tiles := make(map[aoc.Vec2d]bool)
	maxX := 0
	maxY := 0
	for _, line := range lines {
		var tile aoc.Vec2d
		fmt.Sscanf(line, "%d,%d", &tile.X, &tile.Y)
		tiles[tile] = true

		maxX = max(maxX, tile.X)
		maxY = max(maxY, tile.Y)
	}
	printTileMap(tiles, maxX+2, maxY+2)
	return getBiggestRect(tiles)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 50, "Example is wrong!")
	}
}
