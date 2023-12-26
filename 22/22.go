package main

import (
	"bufio"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/tidwall/pinhole"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Brick struct {
	name       int
	x1, y1, z1 int
	x2, y2, z2 int
}

func handleLine(line string, lineIndex int) Brick {
	var b Brick
	b.name = lineIndex + 1
	str := strings.Split(line, "~")
	fmt.Sscanf(str[0], "%d,%d,%d", &b.x1, &b.y1, &b.z1)
	fmt.Sscanf(str[1], "%d,%d,%d", &b.x2, &b.y2, &b.z2)
	return b
}

func toArray(bricks []Brick) [][][]int {
	var arr [][][]int

	maxX, maxY, maxZ := 0, 0, 0
	for _, b := range bricks {
		maxX = max(maxX, b.x1, b.x2)
		maxY = max(maxY, b.y1, b.y2)
		maxZ = max(maxZ, b.z1, b.z2)
	}

	arr = make([][][]int, maxX+1)
	for i := range arr {
		arr[i] = make([][]int, maxY+1)
		for j := range arr[i] {
			arr[i][j] = make([]int, maxZ+1)
		}
	}

	for _, b := range bricks {
		for x := b.x1; x <= b.x2; x++ {
			for y := b.y1; y <= b.y2; y++ {
				for z := b.z1; z <= b.z2; z++ {
					arr[x][y][z] = b.name
				}
			}
		}
	}

	return arr
}

func checkShiftDown(arr *[][][]int, brick Brick) bool {
	if brick.name == 0 {
		fmt.Println("Brick is empty")
		return false
	}

	// fmt.Printf("Checking brick %d at %d %d %d - %d %d %d\n", brick.name, brick.x1, brick.y1, brick.z1, brick.x2, brick.y2, brick.z2)

	for x := brick.x1; x <= brick.x2; x++ {
		for y := brick.y1; y <= brick.y2; y++ {
			for z := brick.z1; z <= brick.z2; z++ {
				if z == 1 {
					// fmt.Println("Found brick at bottom")
					return false
				}
				if !((*arr)[x][y][z-1] == 0 || (*arr)[x][y][z-1] == brick.name) {
					// fmt.Printf("Found brick %d at %d %d %d\n", (*arr)[x][y][z-1], x, y, z-1)
					return false
				}
			}
		}
	}

	// fmt.Println("Brick can be shifted down")
	return true
}

func shiftDown(arr *[][][]int, bricks []Brick) {
	fmt.Println("Sorting bricks")
	sort.Slice(bricks, func(i, j int) bool {
		return min(bricks[i].z1, bricks[i].z2) < min(bricks[j].z1, bricks[j].z2)
	})

	fmt.Println("Shifting down")
	for i, brick := range bricks {
		for checkShiftDown(arr, brick) {
			for x := brick.x1; x <= brick.x2; x++ {
				for y := brick.y1; y <= brick.y2; y++ {
					for z := brick.z1; z <= brick.z2; z++ {
						(*arr)[x][y][z-1] = (*arr)[x][y][z]
						if z == brick.z2 {
							(*arr)[x][y][z] = 0
						}
					}
				}
			}
			bricks[i].z1--
			bricks[i].z2--
			brick = bricks[i]
		}
	}
}

func fillSupports(arr *[][][]int, bricks []Brick, supports map[int][]int, supportedBy map[int][]int) {
	for x := range *arr {
		for y := range (*arr)[x] {
			for z := range (*arr)[x][y] {
				if z+1 < len((*arr)[x][y]) && (*arr)[x][y][z+1] != 0 && (*arr)[x][y][z] != 0 && (*arr)[x][y][z+1] != (*arr)[x][y][z] {
					if !slices.Contains(supports[(*arr)[x][y][z]], (*arr)[x][y][z+1]) {
						supports[(*arr)[x][y][z]] = append(supports[(*arr)[x][y][z]], (*arr)[x][y][z+1])
					}
					if !slices.Contains(supportedBy[(*arr)[x][y][z+1]], (*arr)[x][y][z]) {
						supportedBy[(*arr)[x][y][z+1]] = append(supportedBy[(*arr)[x][y][z+1]], (*arr)[x][y][z])
					}
				}
			}
		}
	}
}

func randomColor() color.RGBA {
	Red := rand.Intn(255)
	Green := rand.Intn(255)
	blue := rand.Intn(255)
	col := color.RGBA{uint8(Red), uint8(Green), uint8(blue), 255}
	return col
}

func plotCoords(arr *[][][]int, bricks []Brick, fileName string, colors map[int]color.RGBA) {
	const scale = 0.1
	p := pinhole.New()

	p.Begin()
	p.DrawLine(0, 0, 0, 0, 0, 2)
	p.Colorize(color.RGBA{255, 0, 0, 255})
	p.End()

	p.Begin()
	p.DrawLine(0, 0, 0, 0, 1, 0)
	p.Colorize(color.RGBA{0, 255, 0, 255})
	p.End()

	p.Begin()
	p.DrawLine(0, 0, 0, 1, 0, 0)
	p.Colorize(color.RGBA{0, 0, 255, 255})
	p.End()

	for _, brick := range bricks {
		p.Begin()

		p.DrawCube(float64(brick.x1)*scale, float64(brick.y1)*scale, float64(brick.z1-1)*scale, float64(brick.x2+1)*scale, float64(brick.y2+1)*scale, float64(brick.z2)*scale)
		p.Colorize(colors[brick.name])

		p.End()
		// fmt.Printf("%d: %d %d %d %d %d %d\n", brick.name, brick.x1, brick.y1, brick.z1, brick.x2, brick.y2, brick.z2)
	}

	p.Translate(-1*scale, -1*scale, -5*scale)
	p.Rotate(-4.01*math.Pi/8, math.Pi/4, 0)
	p.SavePNG(fileName, 1500, 1500, nil)

	// -----------------------------

	p2 := pinhole.New()
	p2.Begin()
	p2.DrawLine(0, 0, 0, 0, 0, 2)
	p2.Colorize(color.RGBA{255, 0, 0, 255})
	p2.End()

	p2.Begin()
	p2.DrawLine(0, 0, 0, 0, 1, 0)
	p2.Colorize(color.RGBA{0, 255, 0, 255})
	p2.End()

	p2.Begin()
	p2.DrawLine(0, 0, 0, 1, 0, 0)
	p2.Colorize(color.RGBA{0, 0, 255, 255})
	p2.End()

	for _, brick := range bricks {
		p2.Begin()

		p2.DrawCube(float64(brick.x1)*scale, float64(brick.y1)*scale, float64(brick.z1-1)*scale, float64(brick.x2+1)*scale, float64(brick.y2+1)*scale, float64(brick.z2)*scale)
		p2.Colorize(colors[brick.name])

		p2.End()
		// fmt.Printf("%d: %d %d %d %d %d %d\n", brick.name, brick.x1, brick.y1, brick.z1, brick.x2, brick.y2, brick.z2)
	}

	fileName = strings.Replace(fileName, ".png", "_top.png", 1)
	p2.SavePNG(fileName, 1500, 1500, nil)

	p2.Translate(-5*scale, 0, -3*scale)
	p2.Rotate(-math.Pi/2, math.Pi/2, 0)
	fileName = strings.Replace(fileName, "_top.png", "_side.png", 1)
	p2.SavePNG(fileName, 1500, 1500, nil)
}

func countDisintegrated(bricks []Brick, supports map[int][]int, supportedBy map[int][]int) int {
	var sum = 0
	for _, brick := range bricks {
		if len(supports[brick.name]) == 0 {
			sum++
			continue
		}
		canBeRemoved := true
		for _, supported := range supports[brick.name] {
			if len(supportedBy[supported]) == 1 {
				canBeRemoved = false
				break
			}
		}
		if canBeRemoved {
			sum++
		}
	}

	return sum
}

func countFalling(bricks []Brick, supports map[int][]int, supportedBy map[int][]int, removed int) int {
	var sum = 0

	queue := []int{removed}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		sum++
		for _, supported := range supports[current] {
			if len(supportedBy[supported]) == 1 {
				queue = append(queue, supported)
			}

			letFall := true
			for _, support := range supportedBy[supported] {
				if support != current && !visited[support] {
					letFall = false
					break
				}
			}
			if letFall {
				queue = append(queue, supported)
			}
		}
	}

	return sum - 1 // -1 because the removed brick does not count
}

func sumFalling(bricks []Brick, supports map[int][]int, supportedBy map[int][]int) int {
	var sum = 0

	for _, brick := range bricks {
		sum += countFalling(bricks, supports, supportedBy, brick.name)
		// fmt.Printf("Brick %d: %d\n", brick.name, countFalling(bricks, supports, supportedBy, brick.name))
	}

	return sum
}

func main() {
	// Open the file
	file, err := os.Open("22.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	var bricks []Brick
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		bricks = append(bricks, handleLine(line, len(bricks)))
	}

	// fmt.Println(bricks)

	fmt.Println("Start processing to array")
	arr := toArray(bricks)

	// colors := make(map[int]color.RGBA)
	// for i := range bricks {
	// 	colors[i+1] = randomColor()
	// }

	// plotCoords(&arr, bricks, "viz_before.png", colors)

	fmt.Println("Start shifting down")
	shiftDown(&arr, bricks)

	fmt.Println("Start filling supports")
	supports := make(map[int][]int)
	supportedBy := make(map[int][]int)
	fillSupports(&arr, bricks, supports, supportedBy)

	// fmt.Println(supportedBy)
	// fmt.Println(supports)

	// plotCoords(&arr, bricks, "viz.png", colors)

	sum = sumFalling(bricks, supports, supportedBy)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
