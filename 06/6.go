package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func simulate(holddownTime int, totalTime int) int {
	return (totalTime - holddownTime) * holddownTime
}

func countdSimPoss(totalTime int, dist int) int {
	poss := 0
	for i := 0; i < totalTime; i++ {
		if simulate(i, totalTime) > dist {
			poss++
		}
	}
	return poss
}

func calcPossibilities(data []string) int {
	data[0] = strings.ReplaceAll(data[0], " ", "")
	data[1] = strings.ReplaceAll(data[1], " ", "")

	reg := regexp.MustCompile("[0-9]+")
	times := reg.FindAllString(data[0], -1)
	dists := reg.FindAllString(data[1], -1)

	acc := 1

	for i := 0; i < len(times); i++ {
		time, err := strconv.ParseInt(times[i], 10, 0)
		check(err)
		dist, err := strconv.ParseInt(dists[i], 10, 0)
		check(err)

		acc *= countdSimPoss(int(time), int(dist))
	}

	return acc
}

func main() {
	// Open the file
	file, err := os.Open("6.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	data := make([]string, 2)
	// Iterate through each line
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	sum += handleLine(line)
	// }
	scanner.Scan()
	data[0] = scanner.Text()
	scanner.Scan()
	data[1] = scanner.Text()

	sum := calcPossibilities(data)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
