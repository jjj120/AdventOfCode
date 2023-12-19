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

type Part struct {
	x, m, a, s int
}

type Rule struct {
	a       string
	smaller bool
	num     int
	next    string
	end     bool
}

func parseRule(line string, ruleMap map[string][]Rule) {
	line = strings.ReplaceAll(line, "}", "")
	split1 := strings.Split(line, "{")
	name := split1[0]

	rules := strings.Split(split1[1], ",")

	for _, rule := range rules {
		var r Rule

		if strings.Contains(rule, "<") {
			r.next = strings.Split(rule, ":")[1]
			rule = strings.Split(rule, ":")[0]

			r.smaller = true
			r.end = false
			num, err := strconv.Atoi(strings.Split(rule, "<")[1])
			check(err)
			r.num = num
			r.a = strings.Split(rule, "<")[0]
		} else if strings.Contains(rule, ">") {
			r.next = strings.Split(rule, ":")[1]
			rule = strings.Split(rule, ":")[0]

			r.smaller = false
			r.end = false
			num, err := strconv.Atoi(strings.Split(rule, ">")[1])
			check(err)
			r.num = num
			r.a = strings.Split(rule, ">")[0]
		} else {
			r.end = true
			r.a = rule
		}

		ruleMap[name] = append(ruleMap[name], r)
	}
}

func parsePart(line string) Part {
	var p Part
	line = strings.ReplaceAll(line, "{", "")
	line = strings.ReplaceAll(line, "}", "")

	split := strings.Split(line, ",")
	for _, s := range split {
		if strings.Contains(s, "x") {
			x, err := strconv.Atoi(strings.Split(s, "=")[1])
			check(err)
			p.x = x
		} else if strings.Contains(s, "m") {
			m, err := strconv.Atoi(strings.Split(s, "=")[1])
			check(err)
			p.m = m
		} else if strings.Contains(s, "a") {
			a, err := strconv.Atoi(strings.Split(s, "=")[1])
			check(err)
			p.a = a
		} else if strings.Contains(s, "s") {
			s, err := strconv.Atoi(strings.Split(s, "=")[1])
			check(err)
			p.s = s
		}
	}
	return p
}

const (
	ACCEPT = iota
	REJECT
	CONTINUE
)

func nextRule(currRule []Rule, part Part, ruleMap map[string][]Rule) ([]Rule, int) {
	if len(currRule) == 0 {
		return currRule, REJECT
	}

	for _, rule := range currRule {
		if rule.end {
			if rule.a == "A" {
				return currRule, ACCEPT
			} else if rule.a == "R" {
				return currRule, REJECT
			} else {
				return ruleMap[rule.a], CONTINUE
			}
		}

		if rule.smaller {
			if rule.a == "x" {
				if part.x < rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			} else if rule.a == "m" {
				if part.m < rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			} else if rule.a == "a" {
				if part.a < rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			} else if rule.a == "s" {
				if part.s < rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			}
		} else {
			if rule.a == "x" {
				if part.x > rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			} else if rule.a == "m" {
				if part.m > rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			} else if rule.a == "a" {
				if part.a > rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			} else if rule.a == "s" {
				if part.s > rule.num {
					return ruleMap[rule.next], CONTINUE
				}
			}
		}
	}

	return currRule, REJECT
}

func accept(part Part, ruleMap map[string][]Rule) bool {
	currRule := ruleMap["in"]
	var ok int
	for {
		// fmt.Printf("CurrPart: %v\tat rule %v\n", part, currRule)

		currRule, ok = nextRule(currRule, part, ruleMap)
		if ok == ACCEPT {
			fmt.Printf("ACCEPT: %v\n", part)
			return true
		} else if ok == REJECT {
			fmt.Printf("REJECT: %v\n", part)
			return false
		}
	}
}

func calcWorth(part Part) int {
	return part.x + part.m + part.a + part.s
}

func calcAccepted(parts []Part, ruleMap map[string][]Rule) int {
	sum := 0
	for _, part := range parts {
		if accept(part, ruleMap) {
			sum += calcWorth(part)
		}
	}
	return sum
}

func main() {
	// Open the file
	file, err := os.Open("19.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	ruleMap := make(map[string][]Rule)

	var acc Rule
	acc.a = "A"
	acc.end = true
	ruleMap["A"] = append(ruleMap["A"], acc)

	var rej Rule
	rej.a = "R"
	rej.end = true
	ruleMap["R"] = append(ruleMap["R"], rej)

	var parts []Part
	rules := true

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			rules = false
			continue
		}

		if rules {
			parseRule(line, ruleMap)
		} else {
			parts = append(parts, parsePart(line))
		}
	}

	fmt.Println(ruleMap)
	fmt.Println(parts)

	sum = calcAccepted(parts, ruleMap)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
