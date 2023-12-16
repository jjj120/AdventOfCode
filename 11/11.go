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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func strToInt(str string) int {
	num := 0
	for _, char := range []byte((str)) {
		num *= 26
		num += int((char) - ('a'))
	}
	return num
}

func intToStr(num int) string {
	var str []byte

	for num != 0 {
		str = append(str, byte((num%26)+'a'))
		num /= 26
	}

	return Reverse(string(str))
}

func checkStr(str string) bool {
	found := false
	for i := 0; i < len(str)-2; i++ {
		if (str[i] == (str[i+1])-1) && (str[i+1] == (str[i+2])-1) {
			found = true
			break
		}
	}

	if !found {
		return false
	}

	if strings.ContainsAny(str, "iol") {
		return false
	}

	counter := 0
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			counter++
			i++
		}
	}

	return counter >= 2
}

func handleLine(line string) int {
	num := strToInt(line)

	for !checkStr(intToStr(num)) {
		num++
	}

	num++

	for !checkStr(intToStr(num)) {
		num++
	}

	fmt.Println(intToStr(num))
	return 0
}

func main() {
	// Open the file
	file, err := os.Open("11.in")
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
