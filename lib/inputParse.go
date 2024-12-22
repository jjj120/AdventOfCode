package lib

import (
	"bufio"
	"fmt"
	"os"
)

func ParseInput(day int, example bool) (input []string, err error) {
	if example {
		inputFilename := fmt.Sprintf("%02d.ex", day)
		input, err = ParseInputFromFile(inputFilename)
	} else {
		inputFilename := fmt.Sprintf("%02d.in", day)
		input, err = ParseInputFromFile(inputFilename)
	}
	if err != nil {
		return
	}
	return
}

func ParseInputFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}
