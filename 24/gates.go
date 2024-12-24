package main

func checkVal(val string, gatesOutput map[string]Gate) bool {
	_, ok := gatesOutput[val]
	return ok
}

type Gate interface {
	compute(gatesOutput map[string]Gate) int
}

type AndGate struct {
	input1       string
	input2       string
	cachedOutput int
	outputValid  bool
}

func (g *AndGate) compute(gatesOutput map[string]Gate) int {
	if g.outputValid {
		return g.cachedOutput
	}
	g.outputValid = true
	if !checkVal(g.input1, gatesOutput) {
		assertf(false, "Invalid input1 of AND gate %s AND %s", g.input1, g.input2)
	}
	if !checkVal(g.input2, gatesOutput) {
		assertf(false, "Invalid input2 of AND gate %s AND %s", g.input1, g.input2)
	}
	g.cachedOutput = gatesOutput[g.input1].compute(gatesOutput) & gatesOutput[g.input2].compute(gatesOutput)
	return g.cachedOutput
}

type OrGate struct {
	input1       string
	input2       string
	cachedOutput int
	outputValid  bool
}

func (g *OrGate) compute(gatesOutput map[string]Gate) int {
	if g.outputValid {
		return g.cachedOutput
	}
	g.outputValid = true
	if !checkVal(g.input1, gatesOutput) {
		assertf(false, "Invalid input1 of AND gate %s AND %s", g.input1, g.input2)
	}
	if !checkVal(g.input2, gatesOutput) {
		assertf(false, "Invalid input2 of AND gate %s AND %s", g.input1, g.input2)
	}
	g.cachedOutput = gatesOutput[g.input1].compute(gatesOutput) | gatesOutput[g.input2].compute(gatesOutput)
	return g.cachedOutput
}

type XOrGate struct {
	input1       string
	input2       string
	cachedOutput int
	outputValid  bool
}

func (g *XOrGate) compute(gatesOutput map[string]Gate) int {
	if g.outputValid {
		return g.cachedOutput
	}
	g.outputValid = true
	if !checkVal(g.input1, gatesOutput) {
		assertf(false, "Invalid input1 of AND gate %s AND %s", g.input1, g.input2)
	}
	if !checkVal(g.input2, gatesOutput) {
		assertf(false, "Invalid input2 of AND gate %s AND %s", g.input1, g.input2)
	}
	g.cachedOutput = gatesOutput[g.input1].compute(gatesOutput) ^ gatesOutput[g.input2].compute(gatesOutput)
	return g.cachedOutput
}

type ConstGate struct {
	value int
}

func (g ConstGate) compute(gatesOutput map[string]Gate) int {
	return g.value
}
