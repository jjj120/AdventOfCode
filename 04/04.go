package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hasZeros(hash string) bool {
	return strings.Compare("00000", hash[:5]) == 0
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func handleLine(line string) int {
	currNumber := 1
	for {
		currLine := line + fmt.Sprint(currNumber)
		currHash := GetMD5Hash(currLine)
		if hasZeros(currHash) {
			return currNumber
		}
		currNumber++
	}
}

func main() {
	// Open the file
	file, err := os.Open("04.in")
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
