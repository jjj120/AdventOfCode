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

type FlipFlop struct {
	name       string
	moduleType int
	connectsTo []string
	state      bool
}

type Conjunction struct {
	name       string
	moduleType int
	connectsTo []string
	watches    map[string]bool
}

type Broadcaster struct {
	name       string
	moduleType int
	connectsTo []string
}

const (
	BROADCASTER = iota
	FLIP_FLOP
	CONJUNCTION
)

const DEBUG_PRINT = false

func handleLine(line string, connections map[string]interface{}) {
	if strings.Contains(line, "broadcaster") {
		var module Broadcaster
		module.moduleType = BROADCASTER
		str := strings.Split(line, " -> ")
		module.name = str[0]
		module.connectsTo = append(module.connectsTo, strings.Split(str[1], ", ")...)
		connections[module.name] = module

	} else if strings.Contains(line, "%") {
		var module FlipFlop
		module.moduleType = FLIP_FLOP
		str := strings.Split(line, " -> ")
		module.name = str[0][1:]
		module.connectsTo = append(module.connectsTo, strings.Split(str[1], ", ")...)
		module.state = false
		connections[module.name] = module

	} else {
		var module Conjunction
		module.moduleType = CONJUNCTION
		str := strings.Split(line, " -> ")
		module.name = str[0][1:]
		module.connectsTo = append(module.connectsTo, strings.Split(str[1], ", ")...)
		module.watches = make(map[string]bool)
		connections[module.name] = module
	}
}

func completeWatches(connections map[string]interface{}) {
	for _, module := range connections {
		// loop through all modules
		switch module := module.(type) {
		case Conjunction:
			conj := module

			// loop through all modules again to check for connections
			for _, module2 := range connections {
				switch module2 := module2.(type) {
				case FlipFlop:
					for _, name := range module2.connectsTo {
						if strings.Compare(name, conj.name) == 0 {
							conj.watches[module2.name] = false
						}
					}
				case Conjunction:
					for _, name := range module2.connectsTo {
						if strings.Compare(name, conj.name) == 0 {
							conj.watches[module2.name] = false
						}
					}
				}
			}
			connections[conj.name] = conj
		}
	}
}

type State struct {
	from  string
	name  string
	pulse bool
}

func simulatePress(connections *map[string]interface{}) [2]int {
	queue := make([]State, 0)

	originalConnections := make(map[string]interface{})
	for k, v := range *connections {
		originalConnections[k] = v
	}

	queue = append(queue, State{"button", "broadcaster", false})
	pulses := [2]int{1, 0} // [low, high] (start wit 1 low pulse from button)

	steps := 0
	for len(queue) > 0 {
		currState := queue[0]
		if DEBUG_PRINT {
			fmt.Printf("CurrState: from: %s, name: %s, pulse: %t\n", currState.from, currState.name, currState.pulse)
			fmt.Printf("%d: HighPulses: %d, LowPulses: %d\n", steps, pulses[1], pulses[0])
			fmt.Println()
		}
		queue = queue[1:]

		module := (*connections)[currState.name]

		if currState.name == "out" {
			continue
		}

		pulse := currState.pulse

		switch module := module.(type) {
		case Broadcaster:
			// Broadcaster -- send high pulse to all connected modules
			for _, name := range module.connectsTo {
				queue = append(queue, State{module.name, name, pulse})
				if pulse {
					pulses[1]++
				} else {
					pulses[0]++
				}
			}
			(*connections)[currState.name] = module
		case FlipFlop:
			// Flip flop -- flips state on low pulse, high pulse does nothing
			if !pulse {
				module.state = !module.state
				for _, name := range module.connectsTo {
					(*connections)[currState.name] = module
					queue = append(queue, State{module.name, name, module.state})
					if module.state {
						pulses[1]++
					} else {
						pulses[0]++
					}
				}
			} else if DEBUG_PRINT {
				fmt.Printf("Got High pulse from %s, doing nothing\n", currState.from)
			}
			(*connections)[currState.name] = module
		case Conjunction:
			// Conjunction -- remembers all pulses, sends high pulse when all pulses are high, low pulse otherwise
			module.watches[currState.from] = pulse
			(*connections)[currState.name] = module

			allTrue := true
			for _, state := range module.watches {
				if !state {
					allTrue = false
					break
				}
			}

			for _, name := range module.connectsTo {
				queue = append(queue, State{module.name, name, !allTrue})
				if !allTrue {
					pulses[1]++
				} else {
					pulses[0]++
				}
			}
			(*connections)[currState.name] = module
		}

		steps++
	}
	return pulses
}

func sumHistory(hist [][2]int) int {
	var sum [2]int
	for _, pulses := range hist {
		sum[0] += pulses[0]
		sum[1] += pulses[1]
	}
	return sum[0] * sum[1]
}

func printConnStatus(connections map[string]interface{}) {
	if !DEBUG_PRINT {
		return
	}

	fmt.Println("---------------------------")
	for k, v := range connections {
		switch val := v.(type) {
		case Broadcaster:
			fmt.Printf("%s\n", k)
		case FlipFlop:
			fmt.Printf("%%%s: %t\n", k, val.state)
		case Conjunction:
			fmt.Printf("&%s\n", k)
		}
	}
	fmt.Println("---------------------------\n")
}

func main() {
	// Open the file
	file, err := os.Open("20.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	connections := make(map[string]interface{})
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line, connections)
	}

	completeWatches(connections)

	printConnStatus(connections)

	numButtonPresses := 1000

	hist := make([][2]int, 0)
	for i := 0; i < numButtonPresses; i++ {
		hist = append(hist, simulatePress(&connections))
		printConnStatus(connections)
	}

	if DEBUG_PRINT {
		fmt.Println(hist)
	}

	sum = sumHistory(hist)

	// loopEnd, loopLen, hist := findLoopLen(connections)

	// loopTimes := (1000 - loopEnd) / loopLen
	// remainder := (1000 - loopEnd) % loopLen

	// low := 0
	// high := 0

	// for i, pulses := range hist {
	// 	if i <= remainder {
	// 		low += pulses[0]
	// 		high += pulses[1]
	// 		continue
	// 	}
	// 	low += loopTimes * pulses[0]
	// 	high += loopTimes * pulses[1]
	// }
	// sum = high * low

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
