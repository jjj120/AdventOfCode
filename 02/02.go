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
	volume := box.l * box.w * box.h
	ribbon := min(2*(box.h+box.l), 2*(box.h+box.w), 2*(box.l+box.w))
	return volume + ribbon
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
