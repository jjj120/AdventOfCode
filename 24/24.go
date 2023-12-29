package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Hailstone struct {
	x, y, z    int
	vx, vy, vz int
}

const TESTAREA_MIN = 200000000000000
const TESTAREA_MAX = 400000000000000

// const TESTAREA_MIN = 7
// const TESTAREA_MAX = 27

const PRINT = true

func (h *Hailstone) move() {
	h.x += h.vx
	h.y += h.vy
	h.z += h.vz
}

func (h1 *Hailstone) intersects(h2 *Hailstone) bool {
	if math.Abs((float64(h1.vx)*float64(h2.vy) - float64(h1.vy)*float64(h2.vx))) < 0.000000000001 {
		return false
	}

	// t := ((h2.x-h1.x)*h2.vy - (h2.y-h1.y)*h2.vx) / (h1.vx*h2.vy - h1.vy*h2.vx)
	t := (float64(h2.x-h1.x)*float64(h2.vy) - float64(h2.y-h1.y)*float64(h2.vx)) / (float64(h1.vx)*float64(h2.vy) - float64(h1.vy)*float64(h2.vx))
	// s := ((h2.x-h1.x)*h1.vy - (h2.y-h1.y)*h1.vx) / (h1.vx*h2.vy - h1.vy*h2.vx)
	s := (float64(h2.x-h1.x)*float64(h1.vy) - float64(h2.y-h1.y)*float64(h1.vx)) / (float64(h1.vx)*float64(h2.vy) - float64(h1.vy)*float64(h2.vx))

	intersectionX := float64(h1.x) + float64(h1.vx)*t
	intersectionY := float64(h1.y) + float64(h1.vy)*t

	if intersectionX < TESTAREA_MIN || intersectionX > TESTAREA_MAX {
		// out of bounds
		return false
	}
	if intersectionY < TESTAREA_MIN || intersectionY > TESTAREA_MAX {
		// out of bounds
		return false
	}
	if t < 0 {
		// intersection is in the past for h1
		return false
	}
	if s < 0 {
		// intersection is in the past for h2
		return false
	}

	if PRINT {
		fmt.Println("Hailstone 1:", h1.x, h1.y, h1.z, "@", h1.vx, h1.vy, h1.vz)
		fmt.Println("Hailstone 2:", h2.x, h2.y, h2.z, "@", h2.vx, h2.vy, h2.vz)
		fmt.Println("Intersection:", intersectionX, intersectionY, t)
		fmt.Println()
	}
	return true
}

func handleLine(line string) Hailstone {
	var hailstone Hailstone

	_, err := fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &hailstone.x, &hailstone.y, &hailstone.z, &hailstone.vx, &hailstone.vy, &hailstone.vz)
	check(err)

	return hailstone
}

func checkCombinations(hails []Hailstone) int {
	var sum = 0

	for i := 0; i < len(hails); i++ {
		for j := i + 1; j < len(hails); j++ {
			if hails[i].intersects(&hails[j]) {
				sum++
			}
		}
	}

	return sum
}

/*
	// the following Z3 code should work, but i cannot get the z3 library to work :(

	func getRockPosSpeed(hailstones []Hailstone) (xRes, yRes, zRes, vxRes, vyRes, vzRes int) {
		config := z3.NewConfig()
		ctx := z3.NewContext(config)
		defer ctx.Close()
		defer config.Close()

		// create Solver
		s := ctx.NewSolver()
		defer s.Close()

		// Create the variables
		x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
		y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())
		z := ctx.Const(ctx.Symbol("z"), ctx.IntSort())
		vx := ctx.Const(ctx.Symbol("vx"), ctx.IntSort())
		vy := ctx.Const(ctx.Symbol("vy"), ctx.IntSort())
		vz := ctx.Const(ctx.Symbol("vz"), ctx.IntSort())

		zero := ctx.Int(0, ctx.IntSort())

		for i := 0; i < 3; i++ {
			currentHailstone := hailstones[i]
			hx := ctx.Int(currentHailstone.x, ctx.IntSort())
			hy := ctx.Int(currentHailstone.y, ctx.IntSort())
			hz := ctx.Int(currentHailstone.z, ctx.IntSort())
			hvx := ctx.Int(currentHailstone.vx, ctx.IntSort())
			hvy := ctx.Int(currentHailstone.vy, ctx.IntSort())
			hvz := ctx.Int(currentHailstone.vz, ctx.IntSort())

			// Create time variable
			timeName := fmt.Sprintf("t%d", i)
			t := ctx.Const(ctx.Symbol(timeName), ctx.IntSort())

			// Create Constraint
			s.Assert(t.Gt(zero))

			// Create the equations
			// hx + t*hvx = x + vx*t
			s.Assert(hx.Add(t.Mul(hvx)).Eq(x.Add(t.Mul(vx))))
			// hy + t*hvy = y + vy*t
			s.Assert(hy.Add(t.Mul(hvy)).Eq(y.Add(t.Mul(vy))))
			// hz + t*hvz = z + vz*t
			s.Assert(hz.Add(t.Mul(hvz)).Eq(z.Add(t.Mul(vz))))

			// print the equations
			if PRINT {
				fmt.Println("Hailstone:", currentHailstone.x, currentHailstone.y, currentHailstone.z, "@", currentHailstone.vx, currentHailstone.vy, currentHailstone.vz)
				fmt.Println(currentHailstone.x, "+", timeName, "*", currentHailstone.vx, "= x +", timeName, "* vx")
				fmt.Println(currentHailstone.y, "+", timeName, "*", currentHailstone.vy, "= y +", timeName, "* vy")
				fmt.Println(currentHailstone.z, "+", timeName, "*", currentHailstone.vz, "= z +", timeName, "* vz")
				fmt.Println()
			}
		}

		if v := s.Check(); v != z3.LBool(ctx.True().Int()) {
			fmt.Println("Unsolveable")
			return
		}

		m := s.Model()
		defer m.Close()

		xRes = int(m.Eval(x).Int())
		yRes = int(m.Eval(y).Int())
		zRes = int(m.Eval(z).Int())
		vxRes = int(m.Eval(vx).Int())
		vyRes = int(m.Eval(vy).Int())
		vzRes = int(m.Eval(vz).Int())

		return
	}
*/

func getRockPosSpeedEquations(hailstones []Hailstone) {
	timeNames := []string{"i", "j", "k"}
	for i := 0; i < 3; i++ {
		currentHailstone := hailstones[i]
		timeName := timeNames[i]

		// hx + t*hvx = x + vx*t
		fmt.Println(currentHailstone.x, "+", timeName, "*", currentHailstone.vx, "= x +", timeName, "* a")
		// hy + t*hvy = y + vy*t
		fmt.Println(currentHailstone.y, "+", timeName, "*", currentHailstone.vy, "= y +", timeName, "* b")
		// hz + t*hvz = z + vz*t
		fmt.Println(currentHailstone.z, "+", timeName, "*", currentHailstone.vz, "= z +", timeName, "* c")
	}
	fmt.Println("x, y, z: position of the rock")
	fmt.Println("a, b, c: speed of the rock")
	fmt.Println("i, j, k: time of collisions between the rock and the hailstones")
	// put the equations into the solver:
	// https://quickmath.com/webMathematica3/quickmath/equations/solve/advanced.jsp
}

func main() {
	// Open the file
	file, err := os.Open("24.in")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sum = 0
	hails := make([]Hailstone, 0)
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		hails = append(hails, handleLine(line))
	}

	// sum = checkCombinations(hails)
	// x, y, z, vx, vy, vz := getRockPosSpeed(hails)
	// fmt.Println(x, y, z, vx, vy, vz)

	getRockPosSpeedEquations(hails)
	x := 180391926345105
	y := 241509806572899
	z := 127971479302113
	sum = x + y + z

	// sum = x + y + z

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
