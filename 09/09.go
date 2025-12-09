package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

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

type TileColorMap map[aoc.Vec2d]Color

func (m TileColorMap) Inside(p aoc.Vec2d) bool {
	inside := false
	c, ok := m[p]
	y := p.Y

	for x := range p.X {
		pos := aoc.Vec2d{X: x, Y: y}
		_, exists := m[pos]
		_, exists_next := m[pos.Add(aoc.Vec2d{X: 1, Y: 0})]

		if exists && !exists_next {
			inside = !inside
		}
	}
	return inside || (ok && c != colWhite)
}

func (m TileColorMap) Outside(p aoc.Vec2d) bool {
	return !m.Inside(p)
}

func printTileMap(tiles TileColorMap, width int, height int) {
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

func drawTileMap(tiles TileColorMap, width, height int) {
	fmt.Println("Start drawing tilemap                                  ")
	const downscaleFactor = 64
	baseImage := image.NewRGBA(image.Rect(0, 0, width/downscaleFactor+1, height/downscaleFactor+1))

	for y := range height / downscaleFactor {
		fmt.Printf("Drawing line %d/%d         \r", y, height/downscaleFactor)
		for x := range width / downscaleFactor {

			hasRed := false
			hasGreen := false
			for xOff := range downscaleFactor {
				for yOff := range downscaleFactor {
					pos := aoc.Vec2d{X: x*downscaleFactor + xOff, Y: y*downscaleFactor + yOff}

					v, ok := tiles[pos]
					if ok && v == colRed {
						hasRed = true
					}
					if ok && v == colGreen {
						hasGreen = true
					}
				}
			}
			if hasRed {
				baseImage.SetRGBA(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			} else if hasGreen {
				baseImage.SetRGBA(x, y, color.RGBA{R: 0, G: 255, B: 0, A: 255})
			} else {
				baseImage.Set(x, y, color.Black)
			}
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, baseImage); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
	fmt.Println("Finished drawing tilemap")
}

func makeTileMap(tiles []aoc.Vec2d, width, height int) TileColorMap {
	fmt.Printf("Start making tilemap")

	colorMap := make(TileColorMap)
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

	return colorMap
}

func getBiggestRect(tiles TileColorMap) int {
	fmt.Println("Start checking rects                             ")

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
			for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y) && isValid; y++ {
				for x := min(p1.X, p2.X); x <= max(p1.X, p2.X) && isValid; x++ {
					fmt.Printf("Checking %v-%v at %d,%d                        \r", p1, p2, x, y)

					if tiles.Outside(aoc.Vec2d{X: x, Y: y}) {
						isValid = false
						break
					}
				}
			}
			if !isValid {
				continue
			}

			size := (max(p1.X, p2.X) - min(p1.X, p2.X) + 1) * (max(p1.Y, p2.Y) - min(p1.Y, p2.Y) + 1)
			fmt.Printf("Got size %d at %v-%v                                     \r", size, p1, p2)
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
	minX := 1 << 31
	minY := 1 << 31
	for _, line := range lines {
		var tile aoc.Vec2d
		fmt.Sscanf(line, "%d,%d", &tile.X, &tile.Y)
		tiles = append(tiles, tile)

		maxX = max(maxX, tile.X)
		maxY = max(maxY, tile.Y)
		minX = min(minX, tile.X)
		minY = min(minY, tile.Y)
	}
	fmt.Printf("X: min: %d, max: %d\n", minX, maxX)
	fmt.Printf("Y: min: %d, max: %d\n", minY, maxY)
	tileMap := makeTileMap(tiles, maxX+1, maxY+1)
	printTileMap(tileMap, maxX+2, maxY+2)
	drawTileMap(tileMap, maxX+2, maxY+2)

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
