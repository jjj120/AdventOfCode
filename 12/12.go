package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sumDataArr(data []interface{}) int {
	sum := 0
	for _, v := range data {
		if obj, ok := v.(map[string]interface{}); ok {
			sum += sumDataObj(obj)
		} else if arr, ok := v.([]interface{}); ok {
			sum += sumDataArr(arr)
		} else if num, ok := v.(float64); ok {
			sum += int(num)
		}
	}
	return sum
}

func sumDataObj(data map[string]interface{}) int {
	sum := 0
	for k, v := range data {
		switch v := v.(type) {
		case string:
			if strings.Contains(v, "red") {
				return 0
			}
		case float64:
			sum += int(v)
		case []interface{}:
			sum += sumDataArr(v)
		case map[string]interface{}:
			sum += sumDataObj(v)
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
	return sum
}

func handleLine(line string) int {
	var v []interface{}

	err := json.Unmarshal([]byte(line), &v)
	check(err)

	sum := sumDataArr(v)

	return sum
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
