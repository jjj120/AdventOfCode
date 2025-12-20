package main

import (
	"fmt"
	"os"
	"os/exec"
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

func run_z3(filename string) int {
	cmd := exec.Command("z3", filename)

	out, err := cmd.Output()
	aoc.Check(err)

	output := string(out)
	outputSplit := strings.Split(output, "\n")

	if outputSplit[0] != "sat" {
		panic("File not sat!")
	}

	outputSplit = outputSplit[2 : len(outputSplit)-2]

	parsedSols := make(map[string]int)

	for i := range len(outputSplit) / 2 {
		line1 := strings.TrimSpace(outputSplit[2*i])
		line2 := strings.TrimSpace(outputSplit[2*i+1])

		varName := ""
		fmt.Sscanf(line1, "(define-fun %s () Int", &varName)

		// remove all chars that are not part of the number
		line2 = strings.ReplaceAll(line2, "(", "")
		line2 = strings.ReplaceAll(line2, ")", "")
		line2 = strings.ReplaceAll(line2, " ", "")

		num, err := strconv.Atoi(line2)
		aoc.Check(err)

		parsedSols[varName] = num
	}

	sum := 0
	for varName, num := range parsedSols {
		if selectExample {
			fmt.Printf("%s: %d\n", varName, num)
		}
		sum += num
	}

	return sum
}

func solveJolts(m Machine) int {
	const z3FileName = ".z3.tmp"
	z3File, err := os.Create(z3FileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer z3File.Close()

	for i, _ := range m.Buttons {
		// (declare-const x Int)
		z3File.WriteString(fmt.Sprintf("(declare-const b%d Int)\n", i))
	}

	for eqi, jolt := range m.Jolts {
		// eq := fmt.Sprintf("(assert (= (%d) ", jolt)
		eq := ""

		for btni, btn := range m.Buttons {
			if (int(btn)>>eqi)&1 == 1 {
				// button is pressed here
				eq = fmt.Sprintf("(+ b%d %s)", btni, eq)
			}
		}
		// _, err := z3File.WriteString(fmt.Sprintf("%d = ", jolt) + eq[2:] + "\n")
		_, err := z3File.WriteString(fmt.Sprintf("(assert (= %d ", jolt) + eq + "))\n")
		aoc.Check(err)
	}

	min_eq := ""
	for btni := range m.Buttons {
		min_eq = fmt.Sprintf("(+ b%d %s)", btni, min_eq)
		_, err := z3File.WriteString(fmt.Sprintf("(assert (>= b%d 0))\n", btni))
		aoc.Check(err)
	}

	z3File.WriteString("(minimize " + min_eq + ")\n")
	z3File.WriteString("(check-sat)\n")
	z3File.WriteString("(get-model)\n")

	z3File.Sync()

	sol := run_z3(z3FileName)

	if selectExample {
		fmt.Println(sol)
		fmt.Println()
	}

	return sol
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
		fmt.Println("Start solving Z3")
	}

	sum := 0
	for i, m := range machines {
		fmt.Printf("Start checking machine %d/%d                  \r", i, len(machines))
		sum += solveJolts(m)
	}

	fmt.Printf("                                                                                                        \r")

	return sum
}

func main() {
	lines, err := aoc.ParseInput(day, selectExample)
	aoc.Check(err)

	fmt.Println("Starting")

	sum := handleLines(lines)

	fmt.Printf("Sum: %d\n", sum)
	if selectExample {
		aoc.Assert(sum == 33, "Example wrong!")
	}
}
