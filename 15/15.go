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

func hashString(part string) int {
	currValue := 0
	for _, char := range part {
		currValue += int(char)
		currValue *= 17
		currValue %= 256
	}
	return currValue
}

type Lens struct {
	label  string
	number int
}

func remove(slice []Lens, s int) []Lens {
	return append(slice[:s], slice[s+1:]...)
}

func handlePart(hashmap map[int][]Lens, part string) {
	if strings.Contains(part, "=") {
		splitStr := strings.Split(part, "=")
		var lens Lens
		lens.label = splitStr[0]
		num, err := strconv.ParseInt(splitStr[1], 10, 0)
		check(err)
		lens.number = int(num)

		found := false
		for i, oldLens := range hashmap[hashString(lens.label)] {
			if strings.Compare(oldLens.label, lens.label) == 0 {
				hashmap[hashString(lens.label)][i].number = lens.number
				found = true
				break
			}
		}
		if !found {
			hashmap[hashString(lens.label)] = append(hashmap[hashString(lens.label)], lens)
		}

	} else {
		splitStr := strings.Split(part, "-")
		label := splitStr[0]

		for i, oldLens := range hashmap[hashString(label)] {
			if strings.Compare(oldLens.label, label) == 0 {
				hashmap[hashString(label)] = remove(hashmap[hashString(label)], i)
				break
			}
		}
	}
}

func calcPower(box int, lenses []Lens) int {
	power := 0
	for i, lens := range lenses {
		power += (box + 1) * (i + 1) * lens.number
	}
	return power
}

func handleLine(line string) int {
	lineParts := strings.Split(line, ",")
	hashmap := make(map[int][]Lens)
	for _, part := range lineParts {
		handlePart(hashmap, part)
	}
	sum := 0
	for box, lenses := range hashmap {
		sum += calcPower(box, lenses)
	}

	return sum
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
