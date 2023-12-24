package main

import (
	"bufio"
	"os"
	"time"
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
