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

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseIngredient(line string) Ingredient {
	var ing Ingredient
	// n, err := fmt.Sscanf(line, "%s: capacity %d, durability %d, flavor %d, texture %d, calories %d", &ing.name, &ing.capacity, &ing.durability, &ing.flavor, &ing.texture, &ing.calories)

	str1 := strings.Split(line, ": ")
	ing.name = str1[0]

	str2 := strings.Split(str1[1], ", ")
	for _, str := range str2 {
		str3 := strings.Split(str, " ")

		var err error
		switch str3[0] {
		case "capacity":
			ing.capacity, err = strconv.Atoi(str3[1])
			check(err)
		case "durability":
			ing.durability, err = strconv.Atoi(str3[1])
			check(err)
		case "flavor":
			ing.flavor, err = strconv.Atoi(str3[1])
			check(err)
		case "texture":
			ing.texture, err = strconv.Atoi(str3[1])
			check(err)
		case "calories":
			ing.calories, err = strconv.Atoi(str3[1])
			check(err)
		}
	}

	return ing
}

func calcScore(ingredients []Ingredient, amount []int) int {
	sum := 0
	for _, am := range amount {
		sum += am
	}

	if sum != 100 {
		return 0
	}

	cap := 0
	dur := 0
	fla := 0
	tex := 0

	for i, ing := range ingredients {
		cap += ing.capacity * amount[i]
		dur += ing.durability * amount[i]
		fla += ing.flavor * amount[i]
		tex += ing.texture * amount[i]
	}

	if cap < 0 || dur < 0 || fla < 0 || tex < 0 {
		return 0
	}

	return cap * dur * fla * tex
}

func calcCalories(ingredients []Ingredient, amount []int) int {
	cal := 0
	for i, ing := range ingredients {
		cal += ing.calories * amount[i]
	}

	return cal
}

func findMaxScore(ingredients []Ingredient, amounts []int) int {
	if len(amounts) == len(ingredients) {
		if calcCalories(ingredients, amounts) != 500 {
			return 0
		}
		return calcScore(ingredients, amounts)
	}

	maxAmount := 100
	if len(amounts) > 0 {
		for _, am := range amounts {
			maxAmount -= am
		}
	}

	maxScore := 0
	for i := 0; i <= maxAmount; i++ {
		newAmounts := make([]int, len(amounts))
		copy(newAmounts, amounts)
		newAmounts = append(newAmounts, i)

		score := findMaxScore(ingredients, newAmounts)
		maxScore = max(maxScore, score)
	}

	return maxScore
}

func handleLine(line string) Ingredient {
	return parseIngredient(line)
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
	ingredients := make([]Ingredient, 0)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		ingredients = append(ingredients, handleLine(line))
	}

	sum = findMaxScore(ingredients, make([]int, 0))

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
