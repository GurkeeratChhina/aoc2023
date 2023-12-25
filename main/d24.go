package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

func intersect_planes(n1 threepoint, d1 int, n2 threepoint, d2 int) (basex float64, basey float64, basez float64, direction threepoint) {
	// n1 <x,y,z> = d1 and n2<x,y,z> = d2
	direction = cross_product(n1, n2)
	fmt.Println(direction)
	basex = 1
	basez = float64(n2.y*d1-n1.y*d2+n1.y*n2.x-n2.y*n1.x) / float64(n2.y*n1.z-n1.y*n2.z)
	basey = (float64(d1-n1.x) - basez*float64(n1.z)) / float64(n1.y)
	return
}

func cross_product(a, b threepoint) (c threepoint) {
	c.x = a.y*b.z - a.z*b.y
	c.y = -a.x*b.z + a.z*b.x
	c.z = a.x*b.y - a.y*b.x
	return
}

func find_intersecting_line(hailstones [][]int) (float64, float64, float64, threepoint) {
	var planes_normals []threepoint
	var planes_constants []int
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := threepoint{hailstones[i][0], hailstones[i][1], hailstones[i][2]}
			h2 := threepoint{hailstones[j][0], hailstones[j][1], hailstones[j][2]}
			dh1 := threepoint{hailstones[i][3], hailstones[i][4], hailstones[i][5]}
			dh2 := threepoint{hailstones[j][3], hailstones[j][4], hailstones[j][5]}
			if dh1.x/dh2.x == dh1.y/dh2.y && dh1.x/dh2.x == dh1.z/dh2.z {
				// parallel lines
				plane_normal := cross_product(dh1, threepoint{h1.x - h2.x, h1.y - h2.y, h1.z - h2.z})
				plane_const := plane_normal.x*h1.x + plane_normal.y*h1.y + plane_normal.z*h1.z
				if !slices.Contains(planes_normals, plane_normal) {
					planes_normals = append(planes_normals, plane_normal)
					planes_constants = append(planes_constants, plane_const)
				}
				if len(planes_normals) == 2 {
					fmt.Println(planes_normals, planes_constants)
					return intersect_planes(planes_normals[0], planes_constants[0], planes_normals[1], planes_constants[1])
				}
			}
		}
	}
	panic("no intersection found??")
}

func find_basepoint(x float64, y float64, z float64, dir threepoint, hailstones [][]int) threepoint {
	fmt.Println(x, y, z)
	t1 := (float64(dir.y)*x - float64(dir.x)*y + float64(dir.x*hailstones[0][1]-dir.y*hailstones[0][0])) / float64(dir.y*hailstones[0][3]-dir.x*hailstones[0][4])
	a1 := (float64(hailstones[0][0]) + t1*float64(hailstones[0][3]) - x) / float64(dir.x)

	t2 := (float64(dir.y)*x - float64(dir.x)*y + float64(dir.x*hailstones[1][1]-dir.y*hailstones[1][0])) / float64(dir.y*hailstones[1][3]-dir.x*hailstones[1][4])
	a2 := (float64(hailstones[1][0]) + t2*float64(hailstones[1][3]) - x) / float64(dir.x)

	fmt.Println(t1, t2)
	c := (t2*a1 - t1*a2) / (t2 - t1)
	x += c * float64(dir.x)
	y += c * float64(dir.y)
	z += c * float64(dir.z)
	fmt.Println(x, y, z)
	return threepoint{int(math.Round(x)), int(math.Round(y)), int(math.Round(z))}
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

	var hailstones [][]int
	for scanner.Scan() {
		hailstones = append(hailstones, extract_nums(scanner.Text()))
	}

	x, y, z, line_dir := find_intersecting_line(hailstones)

	bp := find_basepoint(x, y, z, line_dir, hailstones)
	fmt.Println(bp)

	return bp.x + bp.y + bp.z
}
