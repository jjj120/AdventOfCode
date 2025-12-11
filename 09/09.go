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

type Color int8

const (
	colWhite = iota
	colRed
	colGreen
)

type TileColorMap map[aoc.Vec2d]Color

func printTileSlice(tiles [][]Color) {
	if !selectExample {
		return
	}

	for _, line := range tiles {
		for _, v := range line {
			if v == colRed {
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

func drawTileSlice(tiles [][]Color) {
	fmt.Println("Start drawing tilemap                                  ")
	width := len(tiles[0])
	height := len(tiles)

	const downscaleFactor = 128

	if width <= downscaleFactor || height <= downscaleFactor {
		fmt.Println("Image not printed, image smaller than downscale factor")
		return
	}

	baseImage := image.NewRGBA(image.Rect(0, 0, width/downscaleFactor, height/downscaleFactor))

	for y := range height / downscaleFactor {
		fmt.Printf("Drawing line %d/%d         \r", y, height/downscaleFactor)
		for x := range width / downscaleFactor {

			hasRed := false
			hasGreen := false
			for xOff := range downscaleFactor {
				for yOff := range downscaleFactor {
					color := tiles[y*downscaleFactor+yOff][x*downscaleFactor+xOff]
					hasRed = hasRed || (color == colRed)
					hasGreen = hasGreen || (color == colGreen)
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

	file, err := os.Create(fmt.Sprintf("output_%d.png", downscaleFactor))
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, baseImage); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
	fmt.Println("Finished drawing tilemap")
}

func makeTileSlice(colorMap TileColorMap, width, height int) [][]Color {
	tileMap := make([][]Color, 0, height+2)
	for len(tileMap) < height+2 {
		tileMap = append(tileMap, make([]Color, width+2))
	}
	for tile, color := range colorMap {
		tileMap[tile.Y][tile.X] = color
	}

	fmt.Println("Start filling inside                                        ")
	fmt.Println(width, height, len(tileMap[0]), len(tileMap))

	// color inside green
	for y := range height {
		fmt.Printf("Filling line %d of %d     \r", y, width)
		outside := true
		for x := range width - 1 {
			col := tileMap[y][x]
			nextCol := tileMap[y][x+1]

			colValid := col == colRed || col == colGreen
			nextColValid := nextCol == colRed || nextCol == colGreen

			if colValid && !nextColValid {
				outside = !outside
			}
			if !colValid {
				tileMap[y][x] = colWhite
			}
			if !outside && !colValid {
				tileMap[y][x] = colGreen
			}
		}
	}
	fmt.Printf("                                                              \r")

	return tileMap
}

func makeTileMap(tiles []aoc.Vec2d) (TileColorMap, map[aoc.Vec2d]bool) {
	fmt.Printf("Start making tilemap")

	colorMap := make(TileColorMap)
	mapRed := make(map[aoc.Vec2d]bool)
	prevTile := tiles[0]

	// color edges
	for _, tile := range tiles[1:] {
		mapRed[tile] = true

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

	return colorMap, mapRed
}

func getBiggestRect(tiles [][]Color, corners map[aoc.Vec2d]bool) int {
	fmt.Println("Start checking rects                             ")

	maxSize := 0
	maxSizeP1 := aoc.Vec2d{X: 0, Y: 0}
	maxSizeP2 := aoc.Vec2d{X: 0, Y: 0}
	for p1, _ := range corners {
		for p2, _ := range corners {
			if p1.Equals(p2) {
				continue
			}

			fmt.Printf("Checking %v-%v              \r", p1, p2)

			isValid := true
			// check edges: vertical edges
			for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y) && isValid; y++ {
				tileColor := tiles[y][p1.X]
				if tileColor != colRed && tileColor != colGreen {
					isValid = false
					break
				}

				tileColor = tiles[y][p2.X]
				if tileColor != colRed && tileColor != colGreen {
					isValid = false
					break
				}
			}

			for x := min(p1.X, p2.X); x <= max(p1.X, p2.X) && isValid; x++ {
				tileColor := tiles[p1.Y][x]
				if tileColor != colRed && tileColor != colGreen {
					isValid = false
					break
				}

				tileColor = tiles[p2.Y][x]
				if tileColor != colRed && tileColor != colGreen {
					isValid = false
					break
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
	tileMap, tileMapRed := makeTileMap(tiles)
	tileSlice := makeTileSlice(tileMap, maxX+1, maxY+1)
	printTileSlice(tileSlice)
	drawTileSlice(tileSlice)

	return getBiggestRect(tileSlice, tileMapRed)
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
