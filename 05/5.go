package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type lineStruct struct {
	from   int64
	to     int64
	length int64
}

type dataMap struct {
	title string
	from  string
	to    string
	data  []lineStruct
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLine(line string) int {
	fmt.Print(line)
	return 0
}

func parseData(scanner bufio.Scanner) ([]int, []dataMap) {
	regexNumbers := regexp.MustCompile("[0-9]+")

	scanner.Scan()
	seeds := regexNumbers.FindAllString(scanner.Text(), -1)
	seedsParsed := make([]int, 0)
	for _, seed := range seeds {
		num, err := strconv.ParseInt(seed, 10, 0)
		check(err)
		seedsParsed = append(seedsParsed, int(num))
	}

	scanner.Scan() // to empty line

	var maps []dataMap
	maps = make([]dataMap, 0)

	for scanner.Scan() {
		title := scanner.Text()
		scanner.Scan()
		var data dataMap
		data.title = title
		data.from = strings.Split(strings.Split(title, " ")[0], "-")[0]
		data.to = strings.Split(strings.Split(title, " ")[0], "-")[2]

		data.data = make([]lineStruct, 0)

		for scanner.Text() != "" {
			numbers := regexNumbers.FindAllString(scanner.Text(), -1)
			if len(numbers) < 3 {
				fmt.Printf("Error on finding numbers in %s\n", scanner.Text())
				scanner.Scan()
				continue
			}

			var line lineStruct
			var err error
			line.to, err = strconv.ParseInt(numbers[0], 10, 0)
			check(err)

			line.from, err = strconv.ParseInt(numbers[1], 10, 0)
			check(err)

			line.length, err = strconv.ParseInt(numbers[2], 10, 0)
			check(err)

			data.data = append(data.data, line)

			if !scanner.Scan() {
				break
			}
		}

		maps = append(maps, data)
	}

	return seedsParsed, maps
}

func getRightMap(from string, maps []dataMap) dataMap {
	for _, singleMap := range maps {
		if strings.Compare(from, singleMap.from) == 0 {
			// fmt.Printf("Found map for %s: %s\n", from, singleMap.title)
			return singleMap
		}
	}
	return maps[0]
}

func getToNumber(number int, singleMap dataMap) int {
	for _, line := range singleMap.data {
		if number >= int(line.from) && number <= int(line.from)+int(line.length) {
			// fmt.Printf("Found number %d -> %d in map %s between %d and %d\n", number, number-int(line.from)+int(line.to), singleMap.title, line.from, line.from+line.length)
			// fmt.Printf("Map %s: \n", singleMap.title)
			// fmt.Println(singleMap)
			return number - int(line.from) + int(line.to)
		}
	}
	// fmt.Printf("Found number %d -> %d in map %s in no line\n", number, number, singleMap.title)
	// fmt.Printf("Map %s\n", singleMap.title)
	// fmt.Println(singleMap)

	return number
}

func findLocation(seed int, maps []dataMap) int {
	currFrom := "seed"
	currNumber := seed

	for currFrom != "location" {
		currMap := getRightMap(currFrom, maps)

		currNumber = getToNumber(currNumber, currMap)

		currFrom = currMap.to
	}

	// fmt.Printf("Found location %d -> %d\n\n", seed, currNumber)

	return currNumber
}

func findClosestLocation(scanner bufio.Scanner) int {
	seeds, maps := parseData(scanner)

	seedLocations := make([]int, len(seeds))

	smallestIndex := 0

	for i, seed := range seeds {
		seedLocations[i] = findLocation(seed, maps)
		if seedLocations[i] < seedLocations[smallestIndex] {
			smallestIndex = i
		}
	}

	return seedLocations[smallestIndex]
}

func main() {
	// Open the file
	file, err := os.Open("5.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// var sum = 0
	// // Iterate through each line
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	sum += handleLine(line)
	// }

	sum := findClosestLocation(*scanner)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
