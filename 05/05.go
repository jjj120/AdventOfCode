package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 05
const selectExample = false

func checkHash(s string) string {
	if strings.HasPrefix(s, "00000") {
		return string(s[5])
	}
	return ""
}

func handleLines(lines []string) string {
	line := lines[0]
	h := md5.New()
	pw := ""
	i := 0

	for len(pw) < 8 {
		h.Reset()
		lineWithIndex := fmt.Sprintf("%s%d", line, i)
		io.WriteString(h, lineWithIndex)
		md5Hash := fmt.Sprintf("%x", h.Sum(nil))

		newChar := checkHash(md5Hash)
		pw += newChar
		if len(newChar) > 0 {
			fmt.Printf("%s\r", pw)
		}
		i += 1
	}

	return pw
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %s\n", sum)
}
