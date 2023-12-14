package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reduceLine(line string) int {
	line2 := regexp.MustCompile(`\\\\`).ReplaceAll([]byte(line), []byte(`.`))
	line2 = regexp.MustCompile(`\\"`).ReplaceAll([]byte(line2), []byte(`.`))
	line2 = regexp.MustCompile(`"`).ReplaceAll([]byte(line2), []byte(``))
	line2 = regexp.MustCompile(`\\x[0-9a-fA-F][0-9a-fA-F]`).ReplaceAll([]byte(line2), []byte(`.`))
	fmt.Printf("%s -> %s, %d -> %d\n", line, line2, len(line), len(line2))
	return len(line) - len(line2)
}

func enlargeLine(line string) int {
	line2 := regexp.MustCompile(`\\`).ReplaceAll([]byte(line), []byte(`\\`))
	line2 = regexp.MustCompile(`"`).ReplaceAll([]byte(line2), []byte(`\"`))
	fmt.Printf("%s -> %s, %d -> %d\n", line, line2, len(line), len(line2))
	return 2 + len(line2) - len(line)
}

func handleLine(line string) int {
	return enlargeLine(line)
}

func main() {
	// Open the file
	file, err := os.Open("08.in")
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
