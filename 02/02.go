package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Box struct {
	w, l, h int64
}

func convertToBox(line string) Box {
	dims := strings.Split(line, "x")
	var box Box
	var err error

	box.w, err = strconv.ParseInt(dims[0], 10, 0)
	check(err)
	box.h, err = strconv.ParseInt(dims[1], 10, 0)
	check(err)
	box.l, err = strconv.ParseInt(dims[2], 10, 0)
	check(err)

	return box
}

func handleLine(line string) int64 {
	box := convertToBox(line)
	area := 2*box.l*box.w + 2*box.w*box.h + 2*box.h*box.l

	slack := min(box.h*box.l, box.h*box.w, box.l*box.w)
	return area + slack
}

func main() {
	// Open the file
	file, err := os.Open("02.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum int64 = 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		sum += handleLine(line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
