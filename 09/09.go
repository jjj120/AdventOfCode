package main

import (
	"fmt"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 9
const selectExample = false

type Color int

const (
	colWhite = iota
	colRed
	colGreen
)

func printTileMap(tiles map[aoc.Vec2d]Color, width int, height int) {
	if !selectExample {
		return
	}

	for y := range height {
		for x := range width {
			pos := aoc.Vec2d{X: x, Y: y}

			v, ok := tiles[pos]

			if !ok {
				fmt.Print(aoc.UNICODE_BLOCK)
			} else if v == colRed {
				aoc.ColorPrint(aoc.ConstantToAnsiEscapeString(aoc.ANSI_BRIGHT_RED_FG), aoc.UNICODE_BLOCK)
			} else if v == colGreen {
				aoc.ColorPrint(aoc.ConstantToAnsiEscapeString(aoc.ANSI_BRIGHT_GREEN_FG), aoc.UNICODE_BLOCK)
			} else {
				fmt.Print(aoc.UNICODE_BLOCK)
			}
		}
		fmt.Println("")
	}
}

func makeTileMap(tiles []aoc.Vec2d, width, height int) map[aoc.Vec2d]Color {
	fmt.Printf("Start making tilemap with size %dx%d = %d\n", width, height, width*height)
	// 2147483647 is max
	// 9672230450
	colorMap := make(map[aoc.Vec2d]Color, width*height/16)
	prevTile := tiles[0]

	// color edges
	for _, tile := range tiles[1:] {
		dir := tile.Sub(prevTile)
		// fmt.Printf("%v with len %f (%d) and norm %v\n", dir, dir.Length(), int(dir.Length()), dir.Normalize())
		dir = dir.Normalize()
		aoc.Assert(dir.X != 0 || dir.Y != 0, "Got direction 0!")
		// fmt.Printf("Tile %v to %v with dir %v            \r", prevTile, tile, dir)

		for !prevTile.Equals(tile) {
			// fmt.Printf("%v + %v\n", prevTile, dir)
			prevTile = prevTile.Add(dir)
			fmt.Printf("Tile %v to %v with dir %v            \r", prevTile, tile, dir)

			colorMap[prevTile] = colGreen
		}

		fmt.Printf("Completed Tile %v                                             \r", tile)
		colorMap[tile] = colRed
	}

	dir := tiles[0].Sub(prevTile).Normalize()
	for !prevTile.Equals(tiles[0]) {
		// fmt.Printf("%v + %v\n", prevTile, dir)
		prevTile = prevTile.Add(dir)
		colorMap[prevTile] = colGreen
	}
	colorMap[tiles[0]] = colRed
	fmt.Println("Start coloring inside                ")

	// color inside of all
	for y := range height {
		fmt.Printf("Currently at line %d/%d                                             \r", y, height)
		inside := false

		for x := range width {
			pos := aoc.Vec2d{X: x, Y: y}
			_, exists := colorMap[pos]
			_, exists_next := colorMap[pos.Add(aoc.Vec2d{X: 1, Y: 0})]

			if inside && !exists {
				colorMap[pos] = colGreen
			}

			if exists && !exists_next {
				inside = !inside
			}
		}
	}

	return colorMap
}

func getBiggestRect(tiles map[aoc.Vec2d]Color) int {
	fmt.Println("Start checking rects")

	maxSize := 0
	maxSizeP1 := aoc.Vec2d{X: 0, Y: 0}
	maxSizeP2 := aoc.Vec2d{X: 0, Y: 0}
	for p1, c1 := range tiles {
		if c1 != colRed {
			continue
		}
		for p2, c2 := range tiles {
			if c2 != colRed {
				continue
			}
			if p1.Equals(p2) {
				continue
			}
			fmt.Printf("Checking %v-%v              \r", p1, p2)

			isValid := true
			for x := min(p1.X, p2.X); x <= max(p1.X, p2.X) && isValid; x++ {
				for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
					if c, ok := tiles[aoc.Vec2d{X: x, Y: y}]; !ok || c == colWhite {
						isValid = false
						break
					}
				}
			}
			if !isValid {
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
	tiles := make([]aoc.Vec2d, 0)
	maxX := 0
	maxY := 0
	for _, line := range lines {
		var tile aoc.Vec2d
		fmt.Sscanf(line, "%d,%d", &tile.X, &tile.Y)
		tiles = append(tiles, tile)

		maxX = max(maxX, tile.X)
		maxY = max(maxY, tile.Y)
	}
	tileMap := makeTileMap(tiles, maxX+1, maxY+1)
	printTileMap(tileMap, maxX+2, maxY+2)
	return getBiggestRect(tileMap)
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 24, "Example is wrong!")
	}
}
