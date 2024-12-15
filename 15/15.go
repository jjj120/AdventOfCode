package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const EMPTY = 0
const WALL = 1
const BOX_LEFT = 2
const BOX_RIGHT = 3

const UP = 0
const DOWN = 1
const LEFT = 2
const RIGHT = 3

const ColorReset = "\033[0m"

const ColorBlack = "\033[30m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorMagenta = "\033[35m"
const ColorCyan = "\033[36m"
const ColorGray = "\033[37m"
const ColorWhite = "\033[97m"

const BGBlack = "\033[40m"
const BGRed = "\033[41m"
const BGGreen = "\033[42m"
const BGYellow = "\033[43m"
const BGBlue = "\033[44m"
const BGPurple = "\033[45m"
const BGCyan = "\033[46m"
const BGWhite = "\033[47m"

type Coord struct {
	x, y int
}

var start = Coord{0, 0}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string, index int) []int {
	var row []int
	for i, c := range line {
		switch c {
		case '.':
			row = append(row, EMPTY)
			row = append(row, EMPTY)
		case '#':
			row = append(row, WALL)
			row = append(row, WALL)
		case 'O':
			row = append(row, BOX_LEFT)
			row = append(row, BOX_RIGHT)
		case '@':
			row = append(row, EMPTY)
			row = append(row, EMPTY)
			start = Coord{i * 2, index}
		default:
			panic("Invalid character in input file")
		}
	}
	return row
}

func doMoves(gamefield [][]int, moves []int, robotPos Coord) ([][]int, Coord) {
	// printGamefield(gamefield, robotPos)
	for _, move := range moves {
		gamefield, robotPos = doMove(gamefield, move, robotPos)
		// printMove(move)
		// printGamefield(gamefield, robotPos)
		// fmt.Println()
	}
	return gamefield, robotPos
}

func doMove(gamefield [][]int, move int, robotPos Coord) ([][]int, Coord) {
	var offset Coord

	switch move {
	case UP:
		offset = Coord{0, -1}
	case DOWN:
		offset = Coord{0, 1}
	case LEFT:
		offset = Coord{-1, 0}
	case RIGHT:
		offset = Coord{1, 0}
	}

	targets := []Coord{Coord{robotPos.x, robotPos.y}}

	if gamefield[robotPos.y+offset.y][robotPos.x+offset.x] == WALL {
		// cannot move
		return gamefield, robotPos
	}

	for targetIndex := 0; targetIndex < len(targets); targetIndex++ {
		target := targets[targetIndex]

		nextCoord := Coord{target.x + offset.x, target.y + offset.y}

		if slices.Contains(targets, nextCoord) {
			// ignore targets that have already been found
			continue
		}

		if gamefield[nextCoord.y][nextCoord.x] == WALL {
			// cannot move
			return gamefield, robotPos
		}

		if gamefield[nextCoord.y][nextCoord.x] == BOX_LEFT {
			// add box to targets
			targets = append(targets, nextCoord)
			nextCoord.x++
			targets = append(targets, nextCoord)
		}
		if gamefield[nextCoord.y][nextCoord.x] == BOX_RIGHT {
			// add box to targets
			targets = append(targets, nextCoord)
			nextCoord.x--
			targets = append(targets, nextCoord)
		}
	}

	newGamefield := make([][]int, len(gamefield))
	for i, row := range gamefield {
		newGamefield[i] = make([]int, len(row))
		copy(newGamefield[i], row)
	}

	for _, target := range targets {
		newGamefield[target.y][target.x] = EMPTY
	}

	for _, target := range targets {
		newGamefield[target.y+offset.y][target.x+offset.x] = gamefield[target.y][target.x]
	}

	return newGamefield, Coord{robotPos.x + offset.x, robotPos.y + offset.y}
}

func calcGPS(gamefield [][]int) int {
	gps := 0
	for y, row := range gamefield {
		for x, cell := range row {
			if cell == BOX_LEFT {
				gps += y*100 + x
			}
		}
	}
	return gps
}

func printGamefield(gamefield [][]int, robotPos Coord) {
	for y, row := range gamefield {
		for x, cell := range row {
			if x == robotPos.x && y == robotPos.y {
				fmt.Print(BGCyan, "@", ColorReset)
				continue
			}
			switch cell {
			case EMPTY:
				fmt.Print(" ")
			case WALL:
				fmt.Print(ColorRed, "█", ColorReset)
			case BOX_LEFT:
				fmt.Print(BGYellow, "[", ColorReset)
			case BOX_RIGHT:
				fmt.Print(BGYellow, "]", ColorReset)
			}
		}
		fmt.Println()
	}
}

func printMove(move int) {
	switch move {
	case UP:
		fmt.Println("UP")
	case DOWN:
		fmt.Println("DOWN")
	case LEFT:
		fmt.Println("LEFT")
	case RIGHT:
		fmt.Println("RIGHT")
	}
	fmt.Println("")
}

func main() {
	// Open the file
	file, err := os.Open("15.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	gamefield := [][]int{}
	index := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		gamefield = append(gamefield, handleLine(line, index))
		index++
	}

	moves := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			switch c {
			case '^':
				moves = append(moves, UP)
			case 'v':
				moves = append(moves, DOWN)
			case '<':
				moves = append(moves, LEFT)
			case '>':
				moves = append(moves, RIGHT)
			default:
				panic("Invalid character in input file")
			}
		}
	}

	// printGamefield(gamefield, start)
	// fmt.Println()
	gamefieldAfter, _ := doMoves(gamefield, moves, start)
	// printGamefield(gamefieldAfter, robotPos)
	var sum = 0
	sum = calcGPS(gamefieldAfter)

	// fmt.Printf("Start: %d, %d\n", start[0], start[1])

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
