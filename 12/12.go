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

// returns 0 if line is correct, 1 if line is incorrect, 2 if line has ? in it
func checkLine(line string) int {
	stringParts := strings.Split(line, " ")
	data := stringParts[0]

	for _, char := range data {
		if char == '?' {
			return 2
		}
	}

	checkSums := stringParts[1]

	regDam := regexp.MustCompile("[#]+")
	regNum := regexp.MustCompile("[0-9]+")

	damagedSprings := regDam.FindAllString(data, -1)
	damagedSpringsNumbers := regNum.FindAllString(checkSums, -1)

	if len(damagedSprings) != len(damagedSpringsNumbers) {
		return 1
	}

	for i := 0; i < len(damagedSprings); i++ {
		num, err := strconv.ParseInt(damagedSpringsNumbers[i], 10, 0)
		check(err)

		if len(damagedSprings[i]) != int(num) {
			return 1
		}
	}
	return 0
}

// returns sum if correct and -1 if incorrect
func countLinePoss(line string, startIndex int, channel chan int) {

	check := checkLine(line)
	if check == 0 {
		channel <- 1
		return
	}
	if check == 1 {
		channel <- (-1)
		return
	}

	sum := 0
	for i := startIndex; i < len(line); i++ {
		if line[i] != '?' {
			continue
		}

		poss1 := line[:i] + "." + line[i+1:]
		poss2 := line[:i] + "#" + line[i+1:]

		chan1 := make(chan int)
		chan2 := make(chan int)

		go countLinePoss(poss1, i, chan1)
		go countLinePoss(poss2, i, chan2)

		res1 := <-chan1
		if res1 > 0 {
			// fmt.Printf("Got %d from chan1\n", res1)
			sum += res1
		}

		res2 := <-chan2
		if res2 > 0 {
			// fmt.Printf("Got %d from chan2\n", res2)
			sum += res2
		}
	}

	channel <- sum
}

func handleLine(line string) int {
	channel := make(chan int)
	go countLinePoss(line, 0, channel)
	return <-channel
}

func main() {
	// Open the file
	file, err := os.Open("12.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	i := 0
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\r\r\r\r\r%d", i)
		i++
		sum += handleLine(line)
	}
	fmt.Print("\n")

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
