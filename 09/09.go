package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type file struct {
	index  int
	length int
	empty  bool
}

func (f file) copy() file {
	return file{index: f.index, length: f.length, empty: f.empty}
}

func handleLine(line string) int {
	files := make([]file, 0, len(line))
	for i, c := range line {
		files = append(files, file{index: i / 2, length: int(c - '0'), empty: i%2 == 1})
	}

	shifted := shiftDown(files)

	return calcChecksum(shifted)
}

func shiftDown(files []file) []file {
	for i := len(files) - 1; i > 0; i-- {
		if files[i].empty {
			continue
		}

		emptyIndex := firstEmptyFit(files[:i], files[i])
		if emptyIndex == -1 { // No empty file fits, no need to shift
			continue
		}

		fileToShift := files[i].copy()
		files[i].empty = true // Empty the file

		files[emptyIndex].length -= fileToShift.length
		// Shift the files
		files = append(files[:emptyIndex], append([]file{fileToShift}, files[emptyIndex:]...)...)

		if emptyIndex < i {
			i++
		}
	}
	return files
}

func firstEmptyFit(files []file, fitFile file) int {
	for i, f := range files {
		if f.empty && f.length >= fitFile.length {
			return i
		}
	}
	return -1
}

func calcChecksum(files []file) int {
	line := makeArray(files)
	sum := 0
	for i, num := range line {
		if num != -1 {
			sum += i * num
		}
	}
	return sum
}

func makeArray(files []file) []int {
	line := make([]int, 0, len(files)*2)
	for _, f := range files {
		for j := 0; j < f.length; j++ {
			if f.empty {
				line = append(line, -1)
			} else {
				line = append(line, f.index)
			}
		}
	}
	return line
}

func printFormatted(files []file) {
	for _, f := range files {
		if f.empty {
			fmt.Print(strings.Repeat(".", f.length))
		} else {
			for i := 0; i < f.length; i++ {
				fmt.Print(f.index)
			}
		}
	}
	fmt.Println()
}

func main() {
	// Open the file
	file, err := os.Open("09.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
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
