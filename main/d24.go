package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/davidkleiven/gononlin/nonlin"
	"gonum.org/v1/exp/linsolve"
)

const input24 = "../inputs/d24.txt"

func intersect_lines(p1, p2, d1, d2, valid_range point) bool {
	x1, y1, dx1, dy1, x2, y2, dx2, dy2 := float64(p1.x), float64(p1.y), float64(d1.x), float64(d1.y), float64(p2.x), float64(p2.y), float64(d2.x), float64(d2.y)
	minimum, maximum := float64(valid_range.x), float64(valid_range.y)

	if dx1 == 0 || dx2 == 0 {
		panic("haven't implemented vertical slopes")
	}
	m1 := dy1 / dx1
	m2 := dy2 / dx2

	if m1 == m2 {
		//assume parallel lines dont cross
		return false
	}
	intx := (m1*x1 - m2*x2 + y2 - y1) / (m1 - m2)
	inty := m1*(intx-x1) + y1

	if intx < minimum || intx > maximum || inty < minimum || inty > maximum {
		return false
	}

	t1 := (intx - x1) / dx1
	t2 := (intx - x2) / dx2

	if t1 > 0 && t2 > 0 {
		// fmt.Println(intx, inty)
		return true
	} else {
		return false
	}
}

func d24p1() int {
	defer timeTrack(time.Now(), "d24p1")
	f, err := os.Open(input24)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var hailstones [][]int
	for scanner.Scan() {
		hailstones = append(hailstones, extract_nums(scanner.Text()))
	}
	// fmt.Println(hailstones)

	test_area := point{200000000000000, 400000000000000}
	count := 0

	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if intersect_lines(point{hailstones[i][0], hailstones[i][1]}, point{hailstones[j][0], hailstones[j][1]}, point{hailstones[i][3], hailstones[i][4]}, point{hailstones[j][3], hailstones[j][4]}, test_area) {
				count++
			}
		}
	}
	return count
}

func d24p2() int {
	defer timeTrack(time.Now(), "d24p2")
	f, err := os.Open(input24)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var hailstones [][]float64
	for scanner.Scan() {
		small_nums := extract_nums(scanner.Text())
		var big_nums []float64
		for _, num := range small_nums {
			big_nums = append(big_nums, float64(num))
		}

		hailstones = append(hailstones, big_nums)
	}

	problem := nonlin.Problem{F: func(out, x []float64) {
		out[0] = x[0] + x[3]*x[6] - hailstones[0][0] - hailstones[0][3]*x[6]
		out[1] = x[1] + x[4]*x[6] - hailstones[0][1] - hailstones[0][4]*x[6]
		out[2] = x[2] + x[5]*x[6] - hailstones[0][2] - hailstones[0][5]*x[6]
		out[3] = x[0] + x[3]*x[7] - hailstones[1][0] - hailstones[1][3]*x[7]
		out[4] = x[1] + x[4]*x[7] - hailstones[1][1] - hailstones[1][4]*x[7]
		out[5] = x[2] + x[5]*x[7] - hailstones[1][2] - hailstones[1][5]*x[7]
		out[6] = x[0] + x[3]*x[8] - hailstones[2][0] - hailstones[2][3]*x[8]
		out[7] = x[1] + x[4]*x[8] - hailstones[2][1] - hailstones[2][4]*x[8]
		out[8] = x[2] + x[5]*x[8] - hailstones[2][2] - hailstones[2][5]*x[8]
	},
	}

	x0 := []float64{181152535714637, 181152535714637, 181152535714637, 15, 10, 5, 50, 52, 54}
	solver := nonlin.NewtonKrylov{Maxiter: 1e18, StepSize: 1e-1, Tol: 1e-7, InnerSettings: &linsolve.Settings{Tolerance: 1e-7, MaxIterations: 1e9}}

	res := solver.Solve(problem, x0)

	fmt.Println(res.X)

	return 0
}
