package main

import (
	"bufio"
	"fmt"
	"os"
)

const EMPTY = 0
const WALL = 1
const BOX = 2

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
		case '#':
			row = append(row, WALL)
		case 'O':
			row = append(row, BOX)
		case '@':
			row = append(row, EMPTY)
			start = Coord{i, index}
		default:
			panic("Invalid character in input file")
		}
	}
	return row
}

func doMoves(gamefield [][]int, moves []int, robotPos Coord) ([][]int, Coord) {
	for _, move := range moves {
		gamefield, robotPos = doMove(gamefield, move, robotPos)
		// printMove(move)
		// printGamefield(gamefield, robotPos)
		// fmt.Println()
	}
	return gamefield, robotPos
}

func doMove(gamefield [][]int, move int, robotPos Coord) ([][]int, Coord) {
	if move == UP {
		switch gamefield[robotPos.y-1][robotPos.x] {
		case EMPTY:
			return gamefield, Coord{robotPos.x, robotPos.y - 1}
		case BOX:
			for i := robotPos.y - 1; i >= 0; i-- {
				if gamefield[i][robotPos.x] == BOX {
					continue
				}
				if gamefield[i][robotPos.x] == EMPTY {
					gamefield[i][robotPos.x] = BOX
					gamefield[robotPos.y-1][robotPos.x] = EMPTY
					return gamefield, Coord{robotPos.x, robotPos.y - 1}
				}
				if gamefield[i][robotPos.x] == WALL {
					return gamefield, robotPos
				}
			}
		case WALL:
			return gamefield, robotPos
		}
	}

	if move == DOWN {
		switch gamefield[robotPos.y+1][robotPos.x] {
		case EMPTY:
			return gamefield, Coord{robotPos.x, robotPos.y + 1}
		case BOX:
			for i := robotPos.y + 1; i < len(gamefield); i++ {
				if gamefield[i][robotPos.x] == BOX {
					continue
				}
				if gamefield[i][robotPos.x] == EMPTY {
					gamefield[i][robotPos.x] = BOX
					gamefield[robotPos.y+1][robotPos.x] = EMPTY
					return gamefield, Coord{robotPos.x, robotPos.y + 1}
				}
				if gamefield[i][robotPos.x] == WALL {
					return gamefield, robotPos
				}
			}
		case WALL:
			return gamefield, robotPos
		}
	}

	if move == LEFT {
		switch gamefield[robotPos.y][robotPos.x-1] {
		case EMPTY:
			return gamefield, Coord{robotPos.x - 1, robotPos.y}
		case BOX:
			for i := robotPos.x - 1; i >= 0; i-- {
				if gamefield[robotPos.y][i] == BOX {
					continue
				}
				if gamefield[robotPos.y][i] == EMPTY {
					gamefield[robotPos.y][i] = BOX
					gamefield[robotPos.y][robotPos.x-1] = EMPTY
					return gamefield, Coord{robotPos.x - 1, robotPos.y}
				}
				if gamefield[robotPos.y][i] == WALL {
					return gamefield, robotPos
				}
			}
		case WALL:
			return gamefield, robotPos
		}
	}

	if move == RIGHT {
		switch gamefield[robotPos.y][robotPos.x+1] {
		case EMPTY:
			return gamefield, Coord{robotPos.x + 1, robotPos.y}
		case BOX:
			for i := robotPos.x + 1; i < len(gamefield[robotPos.y]); i++ {
				if gamefield[robotPos.y][i] == BOX {
					continue
				}
				if gamefield[robotPos.y][i] == EMPTY {
					gamefield[robotPos.y][i] = BOX
					gamefield[robotPos.y][robotPos.x+1] = EMPTY
					return gamefield, Coord{robotPos.x + 1, robotPos.y}
				}
				if gamefield[robotPos.y][i] == WALL {
					return gamefield, robotPos
				}
			}
		case WALL:
			return gamefield, robotPos
		}
	}

	panic("Invalid move: " + string(move))
}

func calcGPS(gamefield [][]int) int {
	gps := 0
	for y, row := range gamefield {
		for x, cell := range row {
			if cell == BOX {
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
			case BOX:
				fmt.Print(BGYellow, "O", ColorReset)
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
