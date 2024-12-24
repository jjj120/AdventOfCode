package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func bruteForceSwaps(gatesOutput map[string]Gate, inputBits int) string {
	for out11 := range gatesOutput {
		for out12 := range gatesOutput {
			if out11 == out12 {
				continue
			}
			swaps := map[string]string{
				out11: out12,
			}
			alreadySwapped := make(map[string]bool)
			alreadySwapped[out11] = true
			alreadySwapped[out12] = true

			for out21 := range gatesOutput {
				if _, ok := alreadySwapped[out21]; ok {
					continue
				}
				for out22 := range gatesOutput {
					if _, ok := alreadySwapped[out22]; ok {
						continue
					}
					swaps[out21] = out22
					alreadySwapped[out22] = true
					alreadySwapped[out21] = true

					for out31 := range gatesOutput {
						if _, ok := alreadySwapped[out31]; ok {
							continue
						}
						for out32 := range gatesOutput {
							if _, ok := alreadySwapped[out32]; ok {
								continue
							}
							swaps[out31] = out32
							alreadySwapped[out32] = true
							alreadySwapped[out31] = true
							fmt.Printf("Checking swaps: %s-%s, %s-%s, %s-%s                              \r", out11, out12, out21, out22, out31, out32)

							for out41 := range gatesOutput {
								if _, ok := alreadySwapped[out41]; ok {
									continue
								}
								for out42 := range gatesOutput {
									if _, ok := alreadySwapped[out42]; ok {
										continue
									}
									swaps[out41] = out42

									if checkSwaps(swaps, gatesOutput, inputBits) {
										swapSlice := []string{out11, out21, out31, out41, out12, out22, out32, out42}
										sortedSlice := sort.StringSlice(swapSlice)
										sortedSlice.Sort()
										fmt.Print("                                                     \r")
										return strings.Join(sortedSlice, ",")
									}

									delete(swaps, out41)
								}
							}

							delete(swaps, out31)
							delete(alreadySwapped, out32)
							delete(alreadySwapped, out31)
						}
					}

					delete(swaps, out21)
					delete(alreadySwapped, out22)
					delete(alreadySwapped, out21)
				}
			}
		}
	}
	panic("No solution found")
}

func checkSwaps(swaps map[string]string, gatesOutput map[string]Gate, inputBits int) bool {
	gatesCopy := make(map[string]Gate)
	for k, v := range gatesOutput {
		gatesCopy[k] = v
	}

	for out, in := range swaps {
		if strings.HasPrefix(out, "x") || strings.HasPrefix(out, "y") || strings.HasPrefix(in, "x") || strings.HasPrefix(in, "y") {
			// dont swap inputs
			return false
		}
		gatesCopy[out], gatesCopy[in] = gatesCopy[in], gatesCopy[out]

		gatesCopy[in].invalidateCache()
		gatesCopy[out].invalidateCache()
	}

	output := checkOutputs(gatesCopy, inputBits)

	for out, in := range swaps {
		// once again invalidate the cache for the swapped gates, because they will be swapped back
		gatesCopy[in].invalidateCache()
		gatesCopy[out].invalidateCache()
	}

	return output
}

func checkOutputs(gatesOutput map[string]Gate, inputBits int) bool {
	fmt.Printf("Start checking outputs, should be %d x %d\n", (1<<inputBits)-1, (1<<inputBits)-1)
	for inX := 0; inX < 1<<inputBits; inX++ {
		for inY := 0; inY < 1<<inputBits; inY++ {
			fmt.Printf("Checking %d-%d                              \r", inX, inY)
			if getOutputFromInputs(inX, inY, gatesOutput, inputBits) != (inX + inY) {
				fmt.Printf("Failed for %d+%d\n", inX, inY)
				fmt.Printf("Exp: %044b\n", inX+inY)
				fmt.Printf("Got: %044b\n", getOutputFromInputs(inX, inY, gatesOutput, inputBits))
				return false
			}
		}
	}
	return true
}

func getOutputFromInputs(inX, inY int, gatesOutput map[string]Gate, inputBits int) int {
	for gates := range gatesOutput {
		gatesOutput[gates].invalidateCache()
	}

	for i := 0; i < inputBits; i++ {
		gatesOutput[fmt.Sprintf("x%02d", i)] = ConstGate{(inX >> i) & 1}
		gatesOutput[fmt.Sprintf("y%02d", i)] = ConstGate{(inY >> i) & 1}
	}

	return getOutputNumber(gatesOutput)
}

func getOutputNumber(gatesOutput map[string]Gate) (sum int) {
	for output := range gatesOutput {
		if strings.HasPrefix(output, "z") {
			outVal := gatesOutput[output].compute(gatesOutput)
			outNum, err := strconv.Atoi(output[1:])
			check(err)
			sum |= outVal << outNum
		}
	}
	return
}
