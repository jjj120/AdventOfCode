package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Relation struct {
	from      string
	to        string
	happiness int
}

func parseRelation(line string) Relation {
	var rel Relation

	rel.from = strings.Split(line, " ")[0]
	rel.to = strings.Split(line, " ")[len(strings.Split(line, " "))-1]
	rel.to = strings.TrimRight(rel.to, ".")

	if strings.Contains(line, "gain") {
		rel.happiness, _ = strconv.Atoi(strings.Split(line, " ")[3])
	} else {
		rel.happiness, _ = strconv.Atoi("-" + strings.Split(line, " ")[3])
	}

	return rel
}

func handleLine(relations map[string]Relation, line string) {
	rel := parseRelation(line)

	relations[rel.from+rel.to] = rel
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}

	return result
}

func checkHappiens(relations map[string]Relation, seating []string) int {
	var sum = 0

	for i := 0; i < len(seating); i++ {
		var from = seating[i]
		var to = seating[(i+1)%len(seating)]

		sum += relations[from+to].happiness
		sum += relations[to+from].happiness
	}

	return sum
}

func optimizePlacement(relations map[string]Relation, seating []string) int {
	var sum = -1000000000

	var people []string
	for k := range relations {
		people = append(people, relations[k].from)
	}

	people = removeDuplicates(people)

	if len(people) == len(seating) {
		return checkHappiens(relations, seating)
	}

	for i := 0; i < len(people); i++ {
		if !slices.Contains(seating, people[i]) {
			sum = max(sum, optimizePlacement(relations, append(seating, people[i])))
		}
	}

	return sum
}

func main() {
	// Open the file
	file, err := os.Open("13.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	relations := make(map[string]Relation)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(relations, line)
	}

	sum = optimizePlacement(relations, []string{})

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
