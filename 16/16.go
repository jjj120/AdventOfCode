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

type Aunt struct {
	Number int
	Things map[string]int
}

var mainAunt Aunt = Aunt{
	Number: -1,
	Things: map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	},
}

func parseAunt(line string) Aunt {
	var aunt Aunt
	var err error

	strings1 := strings.Split(line, ": ")
	strings2 := strings.Split(strings1[0], " ")
	aunt.Number, err = strconv.Atoi(strings2[1])
	check(err)

	things := strings.Split(line, ": ")
	things = things[1:]
	thingsLine := strings.Join(things, ": ")
	things = strings.Split(thingsLine, ", ")
	aunt.Things = make(map[string]int)

	for _, thing := range things {
		thing2 := strings.Split(thing, ": ")
		aunt.Things[thing2[0]], err = strconv.Atoi(thing2[1])
		check(err)
	}

	return aunt
}

func findAunt(aunts []Aunt) int {
	for _, aunt := range aunts {
		if aunt.Number == -1 {
			continue
		}

		found := true
		for thing, value := range aunt.Things {
			if thing == "cats" || thing == "trees" {
				if mainAunt.Things[thing] >= value {
					found = false
					break
				}
			} else if thing == "pomeranians" || thing == "goldfish" {
				if mainAunt.Things[thing] <= value {
					found = false
					break
				}
			} else if mainAunt.Things[thing] != value {
				found = false
				break
			}
		}

		if found {
			return aunt.Number
		}
	}
	return -1
}

func main() {
	// Open the file
	file, err := os.Open("16.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	aunts := make([]Aunt, 0)
	aunts = append(aunts, mainAunt)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		aunt := parseAunt(line)
		aunts = append(aunts, aunt)
	}

	sum = findAunt(aunts)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
