package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/jjj120/AdventOfCode/lib"
)

const day = 10
const selectExample = false

type Button int
type Light int

type Machine struct {
	LightTarget Light
	LightLength int
	Buttons     []Button
	Jolts       []int

	_CurrentLight Light
}

func (m *Machine) Reset() {
	m._CurrentLight = 0
}

func (m *Machine) Check() bool {
	return m._CurrentLight == m.LightTarget
}

func (m *Machine) Press(buttonIndex int) bool {
	m._CurrentLight ^= Light(m.Buttons[buttonIndex])
	return m.Check()
}

func (m *Machine) PressAll(buttons []int) bool {
	m.Reset()
	for _, b := range buttons {
		m.Press(b)
	}
	return m.Check()
}

func printLight(light Light, length int) {

	for i := range length {
		if (light>>i)&1 == 1 {
			aoc.ColorPrint(aoc.ConstantToAnsiEscapeString(aoc.ANSI_GREEN_FG), aoc.UNICODE_BLACK_CIRCLE)
		} else {
			aoc.ColorPrint(aoc.RGBtoAnsiEscapeString(32, 32, 32, true), ".")
		}
	}
	fmt.Println()
}

func (m *Machine) PrintLight() {
	printLight(m.LightTarget, m.LightLength)
}

func (m *Machine) PrintCurrentLight() {
	printLight(m._CurrentLight, m.LightLength)
}

func (m *Machine) PrintButtons() {
	fmt.Println("Buttons:")
	for _, b := range m.Buttons {
		printLight(Light(b), m.LightLength)
	}
}

func (m *Machine) Print() {
	fmt.Printf("Machine:\n")
	m.PrintLight()
	m.PrintButtons()
	m.PrintCurrentLight()
}

func checkCombinations(m Machine, btnsToUse int) bool {
	fmt.Printf("Checking sequences with length %d\r", btnsToUse)
	current := make([]int, btnsToUse)
	var helper func(int) bool
	helper = func(pos int) bool {
		if pos == btnsToUse {
			res := m.PressAll(current)
			// fmt.Printf("Checking sequence %v --> %t\n", current, res)
			return res
		}
		res := false
		for i := 0; i < len(m.Buttons); i++ {
			current[pos] = i
			res = res || helper(pos+1)
		}
		return res
	}
	return helper(0)
}

func handleLines(lines []string) int {
	machines := make([]Machine, 0, len(lines))
	for _, line := range lines {
		var machine Machine
		machine.Reset()
		splitLine := strings.Split(line, " ")
		lightTargetStr := splitLine[0]
		joltStr := splitLine[len(splitLine)-1]
		buttonsStr := splitLine[1 : len(splitLine)-1]

		machine.LightLength = len(lightTargetStr) - 2
		machine.LightTarget = 0
		for i, v := range lightTargetStr[1 : len(lightTargetStr)-1] {
			if v == '#' {
				machine.LightTarget |= 1 << i
			}
		}

		for _, bs := range buttonsStr {
			var btn Button = 0
			for _, b := range strings.Split(bs[1:len(bs)-1], ",") {
				bi, err := strconv.Atoi(b)
				aoc.Check(err)

				btn |= 1 << bi
			}
			machine.Buttons = append(machine.Buttons, btn)
		}

		for _, j := range strings.Split(joltStr[1:len(joltStr)-1], ",") {
			ji, err := strconv.Atoi(j)
			aoc.Check(err)

			machine.Jolts = append(machine.Jolts, ji)
		}

		machines = append(machines, machine)
	}

	if selectExample {
		aoc.Assert(machines[0].PressAll([]int{0, 1, 2}) == true, "PressAll not correct")
		machines[0].Reset()
	}

	sum := 0
	for _, m := range machines {
		m.Print()
		l := 1
		for !checkCombinations(m, l) {
			l++
		}
		sum += l
		fmt.Println("                                                 ")
	}

	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 7, "Example wrong!")
	}
}
