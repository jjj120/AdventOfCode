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

func checkHash(s string) (int, string) {
	if strings.HasPrefix(s, "00000") {
		if (s[5]-'0') >= 0 && (s[5]-'0') < 8 {
			return int(s[5] - '0'), string(s[6])
		}
	}
	return -1, ""
}

func handleLines(lines []string) string {
	line := lines[0]
	h := md5.New()
	pw := make([]string, 8)
	for i := range pw {
		pw[i] = ""
	}
	i := 0
	found := 0

	for found < 8 {
		h.Reset()
		lineWithIndex := fmt.Sprintf("%s%d", line, i)
		io.WriteString(h, lineWithIndex)
		md5Hash := fmt.Sprintf("%x", h.Sum(nil))

		ind, newChar := checkHash(md5Hash)
		if ind >= 0 {
			// there is a new char
			if len(pw[ind]) == 0 {
				// the index was not taken before
				pw[ind] = newChar
				found += 1

				// print out the new password
				for _, c := range pw {
					if len(c) > 0 {
						fmt.Printf("%s", c)
					} else {
						fmt.Printf("_")
					}
				}
				fmt.Printf("\r")
			}
		}

		i += 1
	}

	sol := ""
	for _, c := range pw {
		sol += c
	}
	return sol
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %s\n", sum)
}
