package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 04
const selectExample = false

type Room struct {
	name     string
	sectorID int
	checksum string
}

func parseRoom(l string) Room {
	r := regexp.MustCompile("([a-z-]+)-([0-9]+)\\[([a-z]+)\\]")
	parsed := r.FindStringSubmatch(l)
	num, err := strconv.Atoi(parsed[2])
	aoc.Check(err)
	return Room{name: parsed[1], sectorID: num, checksum: parsed[3]}
}

func shiftRune(r rune, a int) rune {
	return (rune(int(r-'a')+a)%26 + 'a')
}

func decryptName(r Room) string {
	name := ""

	for _, ru := range strings.ReplaceAll(r.name, "-", " ") {
		if ru == ' ' {
			name += " "
			continue
		}
		name += string(shiftRune(ru, r.sectorID))
	}
	return name
}

func checkRoom(r Room) int {
	type RuneInt struct {
		r rune
		i int
	}

	m := make(map[rune]int)
	for _, c := range r.name {
		if c == '-' {
			continue // ignore - in counting
		}
		if _, ok := m[c]; !ok {
			m[c] = 1
		} else {
			m[c] += 1
		}
	}

	s := make([]RuneInt, 0, len(m))
	for k, v := range m {
		s = append(s, RuneInt{r: k, i: v})
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i].i == s[j].i {
			return s[i].r < s[j].r
		}
		return s[i].i > s[j].i
	})

	for i, ru := range r.checksum {
		if ru != s[i].r {
			return 0
		}
	}

	decName := decryptName(r)

	if strings.Contains(decName, "north") {
		fmt.Printf("Found match: %s\n", decName)
		return r.sectorID
	}

	return 0
}

func handleLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		room := parseRoom(line)
		sum += checkRoom(room)
	}
	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	aoc.Assert(!selectExample || sum == 1514, "Example wrong!")
}
